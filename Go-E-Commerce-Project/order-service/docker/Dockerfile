FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Copy go modules first
COPY go.mod go.sum ./
RUN go mod download

# Copy ALL source code
COPY . .

# Debug: Show what was actually copied
RUN echo "=== Files in /app ===" && ls -la
RUN echo "=== cmd directory ===" && ls -la cmd/ || echo "cmd/ not found"
RUN echo "=== server directory ===" && ls -la cmd/server/ || echo "cmd/server not found"

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /order-service ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /order-service .

EXPOSE 8083

CMD ["./order-service"]