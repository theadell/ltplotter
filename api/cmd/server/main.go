package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"ltplotter/gen/pb"
	"ltplotter/internal/rpc"
	"ltplotter/pkg/adapter"
	"ltplotter/pkg/config"
	"ltplotter/pkg/mapper"
	"ltplotter/pkg/parser"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/urfave/negroni"
)

func main() {
	config := config.LoadConfig()

	proxy, err := createReverseProxy(config.PGFPlotsServiceURL)
	if err != nil {
		slog.Warn("Failed to parse PGFPlots service URL", "error", err.Error())
	}
	chartRPCManager := rpc.NewChartServiceClientManager(config.PGFPlotsServiceURL)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/parse", parseSimulationHandler)
	mux.HandleFunc("POST /api/v1/pgfplot", createPGFPlotHandler(proxy))
	mux.HandleFunc("POST /api/v2/pgfplot", createPGFPlotHandlerRPC(chartRPCManager))

	n := negroni.New(adapter.ToNegroni(middleware.Recoverer), adapter.ToNegroni(middleware.Logger))
	n.UseHandler(mux)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:           n,
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 2,
		WriteTimeout:      time.Second * 5,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info(fmt.Sprintf("Server is running on port %d", config.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "err", err.Error())
			os.Exit(1)
		}
	}()

	<-stop
	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "err", err.Error())
	}

	slog.Info("Server has shut down gracefully")
}

func parseSimulationHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		slog.Debug("failed to obtain file from request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	simData, err := parser.Parse(file)
	if err != nil {
		slog.Error("failed to parse file", "file", file, "error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := mapper.ToDataSet(simData)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to encode parsed data into json", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createPGFPlotHandler(proxy *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if proxy == nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		proxy.ServeHTTP(w, r)
	}
}

func createReverseProxy(targetURL string) (*httputil.ReverseProxy, error) {
	if targetURL == "" {
		return nil, fmt.Errorf("target URL is empty")
	}

	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)
	proxy.Director = func(r *http.Request) {
		r.URL.Scheme = parsedURL.Scheme
		r.URL.Host = parsedURL.Host
		r.URL.Path = "/generate-plot"
		r.Host = parsedURL.Host
		slog.Info("proxy passing request", "url", r.URL.String())
	}
	return proxy, nil
}

func createPGFPlotHandlerRPC(clientManager *rpc.ChartServiceClientManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := clientManager.GetClient()
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}

		var req pb.PlotRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		res, err := client.GeneratePlot(ctx, &req)
		if err != nil {
			log.Printf("Error generating chart: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=chart.pdf")

		w.Write(res.Pdf)
	}
}
