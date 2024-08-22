package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"ltplotter/gateway/internal/rpc"
	"ltplotter/gateway/pkg/expr/validation"
	"ltplotter/gateway/pkg/jobmanager"
	"ltplotter/gen/pb"
	"net/http"
	"time"
)

func CreateExprPlotHandler(clientManager *rpc.ExpressionPlotServiceClientManager, jm *jobmanager.JobManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := clientManager.GetClient()
		if err != nil {
			http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
			return
		}

		var req pb.ExprPlotRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			slog.Error("failed to decode request body", "error", err.Error())
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if err := validation.ValidateRequest(&req); err != nil {
			slog.Warn("recieved invalid or potentially malicious request", "validation error", err.Error(), "request", &req)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		jobID := jm.GenerateJobID()
		jm.SetJobResult(jobID, jobmanager.StatusPending, nil)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("X-Request-ID", jobID)
		w.Header().Set("Location", fmt.Sprintf("/api/v2/plot/expr/%s", jobID))

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"jobID": jobID,
		})
		slog.Debug("Dispatched job for expression plot generation",
			"jobID", jobID,
			"expressions", req.Plots,
		)

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()
			res, err := client.GeneratePlot(ctx, &req)
			if err != nil {
				slog.Error("Failed to generate plot from expression", "error", err.Error())
				jm.SetJobResult(jobID, jobmanager.StatusFailed, err)
				return
			}
			slog.Info("Plot generation from expression completed successfully", "jobID", jobID)
			jm.SetJobResult(jobID, jobmanager.StatusCompleted, res)
		}()
	}
}

func CreateGetExprPlotHandler(jm *jobmanager.JobManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobID := r.PathValue("id")
		if jobID == "" {
			http.Error(w, "Bad Request: missing job ID", http.StatusBadRequest)
			return
		}

		format := r.URL.Query().Get("format")

		result, ok := jm.GetJobResult(jobID)
		if !ok {
			http.Error(w, "Job Not Found", http.StatusNotFound)
			return
		}

		switch result.Status {
		case jobmanager.StatusPending:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Retry-After", "2")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(map[string]string{
				"jobID":   jobID,
				"status":  "pending",
				"message": "Job is still in progress. Please try again later.",
			})
			return

		case jobmanager.StatusFailed:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"jobID":  jobID,
				"status": "failed",
				"error":  result.Result.(string),
			})
			return

		case jobmanager.StatusCompleted:
			res := result.Result.(*pb.ExprPlotResponse)
			switch format {
			case "latex":
				w.Header().Set("Content-Type", "application/x-latex")
				w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=plot_%s.tex", jobID))
				if _, err := w.Write([]byte(res.LatexSrcCode)); err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			case "pdf", "":
				w.Header().Set("Content-Type", "application/pdf")
				w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=plot_%s.pdf", jobID))
				if _, err := w.Write(res.PdfPlot); err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			default:
				http.Error(w, "Unsupported format", http.StatusBadRequest)
			}
			return

		default:
			http.Error(w, "Unexpected result type", http.StatusInternalServerError)
			return
		}
	}
}
