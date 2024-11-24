package routes

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/urfave/negroni"
	"ltplotter/gateway/internal/handlers"
	"ltplotter/gateway/internal/rpc"
	"ltplotter/gateway/pkg/adapter"
	"ltplotter/gateway/pkg/jobmanager"
	"net/http"
)

func SetupRoutes(
	jm *jobmanager.JobManager,
	chartRPCManager *rpc.ChartServiceClientManager,
	exprPlotRPCManager *rpc.ExpressionPlotServiceClientManager,
	manimRPCManager *rpc.ManimServiceClientManager,
) http.Handler {
	mux := http.NewServeMux()
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	// Add your routes here
	mux.HandleFunc("/api/v2/parse", handlers.ParseSimulationHandler)
	mux.HandleFunc("/api/v2/plot/data", handlers.CreateDataPlotHandlerRPC(chartRPCManager))
	mux.HandleFunc("/api/v2/plot/expr", handlers.CreateExprPlotHandler(exprPlotRPCManager, jm, v))
	mux.HandleFunc("/api/v2/plot/manim", handlers.CreateManimHandler(manimRPCManager))
	mux.HandleFunc("/api/v2/plot/expr/{id}", handlers.CreateGetExprPlotHandler(jm))

	n := negroni.New(
		adapter.ToNegroni(middleware.Recoverer),
		adapter.ToNegroni(middleware.Logger),
	)
	n.UseHandler(mux)

	return n
}
