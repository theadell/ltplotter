FROM golang:1.22.5-alpine AS builder

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o app ./gateway/cmd/server/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /api/app .

EXPOSE 8080

ENTRYPOINT ["./app"]
CMD ["-port", "8080", "-host", "0.0.0.0"]
