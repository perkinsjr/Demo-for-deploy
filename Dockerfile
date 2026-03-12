# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /app/demo_api_go .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/demo_api_go .

EXPOSE 8080

CMD ["./demo_api_go"]
