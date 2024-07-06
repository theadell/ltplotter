package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"ltplotter/pkg/mapper"
	"ltplotter/pkg/parser"
	"net/http"
)

func main() {

	http.HandleFunc("POST /api/parse", parseSimulation)
	slog.Info("Server is running on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func parseSimulation(w http.ResponseWriter, r *http.Request) {
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
