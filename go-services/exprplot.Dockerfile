# Stage 1: Build the Go application
FROM golang:1.22.5-alpine AS builder

WORKDIR /grpc-service

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o app ./exprplot/cmd/server/main.go

FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    texlive-latex-base \
    texlive-latex-extra 

RUN useradd -m latexuser

WORKDIR /app

COPY --from=builder /grpc-service/app .
COPY --from=builder /grpc-service/exprplot/static/expr_plot.go.tex ./expr_plot_template.go.tex

RUN chown -R latexuser:latexuser /app
USER latexuser

EXPOSE 50051

ENTRYPOINT ["./app"]
CMD ["-port", "50051", "-host", "0.0.0.0"]
