package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"log/slog"
	"ltplotter/gateway/internal/rpc"
	"ltplotter/gen/pb"
	"net/http"
	"net/http/httputil"
	"time"
)

func CreateDataPlotHandlerRPC(clientManager *rpc.ChartServiceClientManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := clientManager.GetClient()
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}

		var req pb.PlotRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			slog.Error("failed to decode request body", "error", err.Error())
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

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"latex": res.Latex,
			"pdf":   base64.StdEncoding.EncodeToString(res.Pdf),
		})
	}
}

func CreateDatalotHandlerREST(proxy *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if proxy == nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}
		proxy.ServeHTTP(w, r)
	}
}
