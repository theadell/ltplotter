package handlers

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"ltplotter/gateway/internal/rpc"
	"ltplotter/gen/pb"
	"net/http"
	"time"
)

func CreateManimHandler(clientManager *rpc.ManimServiceClientManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := clientManager.GetClient()
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}

		var req pb.ManimRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			slog.Error("failed to decode request body", "error", err.Error())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		res, err := client.GenerateVideo(ctx, &req)
		if err != nil {
			log.Printf("Error generating video: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		log.Printf(res.ManimVideoUrl)
		w.WriteHeader(http.StatusAccepted)
		err = json.NewEncoder(w).Encode(map[string]string{
			"manimVideoUrl": res.ManimVideoUrl,
		})
		if err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	}
}
