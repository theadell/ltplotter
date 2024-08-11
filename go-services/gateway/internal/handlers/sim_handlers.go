package handlers

import (
	"encoding/json"
	"log/slog"
	"ltplotter/gateway/pkg/mapper"
	"ltplotter/gateway/pkg/parser"
	"net/http"
)

func ParseSimulationHandler(w http.ResponseWriter, r *http.Request) {
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
