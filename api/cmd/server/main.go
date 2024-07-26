package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"ltplotter/pkg/mapper"
	"ltplotter/pkg/parser"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Port               int
	Host               string
	PGFPlotsServiceURL string
}

func main() {
	config := &Config{}

	flag.IntVar(&config.Port, "port", 8080, "TCP Port to bind server to")
	flag.StringVar(&config.Host, "host", "localhost", "Network to bind to")

	flag.Parse()

	config.PGFPlotsServiceURL = os.Getenv("PGFPLOT_SVC_URL")
	proxy, err := createReverseProxy(config.PGFPlotsServiceURL)
	if err != nil {
		log.Printf("Warning: Failed to parse PGFPlots service URL: %v\n", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/parse", parseSimulationHandler)
	mux.HandleFunc("POST /api/v1/pgfplot", createPGFPlotHandler(proxy))

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:           mux,
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
		slog.Info("RECEIVED A REQUEST", "host", r.Host, "path", r.URL.Path, "url", r.URL)
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
