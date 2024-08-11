package main

import (
	"context"
	"fmt"
	"log/slog"
	"ltplotter/gateway/internal/config"
	"ltplotter/gateway/internal/routes"
	"ltplotter/gateway/internal/rpc"
	"ltplotter/gateway/pkg/jobmanager"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := config.LoadConfig()
	jm := jobmanager.NewJobManager(5*time.Minute, 10*time.Minute)

	dataPlotRPCManager := rpc.NewChartServiceClientManager(config.PGFPlotServiceURLRPC, config.CertsPath)
	exprPlotRPCManager := rpc.NewExpressionPlotServiceClientManager(config.ExprPlotServiceUrl, config.CertsPath)

	handler := routes.SetupRoutes(*config, jm, dataPlotRPCManager, exprPlotRPCManager)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:           handler,
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
