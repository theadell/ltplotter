package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"ltplotter/pkg/mapper"
	"ltplotter/pkg/parser"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port int
var host string

func main() {

	flag.IntVar(&port, "port", 8080, "TCP Port to bind server to")
	flag.StringVar(&host, "host", "localhost", "Network to bind to")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/parse", parseSimulationHandler)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", host, port),
		Handler:           mux,
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 2,
		WriteTimeout:      time.Second * 5,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info(fmt.Sprintf("Server is running on port %d", port))
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
