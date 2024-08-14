package routes

import (
	"fmt"
	"log/slog"
	"ltplotter/gateway/internal/config"
	"ltplotter/gateway/internal/handlers"
	"ltplotter/gateway/internal/rpc"
	"ltplotter/gateway/pkg/adapter"
	"ltplotter/gateway/pkg/jobmanager"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/urfave/negroni"
)

func SetupRoutes(config config.Config, jm *jobmanager.JobManager, chartRPCManager *rpc.ChartServiceClientManager, exprPlotRPCManager *rpc.ExpressionPlotServiceClientManager) http.Handler {
	mux := http.NewServeMux()

	// Add your routes here
	mux.HandleFunc("/api/v2/parse", handlers.ParseSimulationHandler)

	proxy, err := createReverseProxy(config.PGFPlotsServiceURL)
	if err != nil {
		slog.Warn("Failed to parse PGFPlots service URL", "error", err.Error())
	}
	mux.HandleFunc("/api/v1/plot/data", handlers.CreateDatalotHandlerREST(proxy))
	mux.HandleFunc("/api/v2/plot/data", handlers.CreateDataPlotHandlerRPC(chartRPCManager))

	mux.HandleFunc("/api/v2/plot/expr", handlers.CreateExprPlotHandler(exprPlotRPCManager, jm))
	mux.HandleFunc("/api/v2/plot/expr/{id}", handlers.CreateGetExprPlotHandler(jm))

	n := negroni.New(adapter.ToNegroni(middleware.Recoverer), adapter.ToNegroni(middleware.Logger))
	n.UseHandler(mux)

	return n
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
