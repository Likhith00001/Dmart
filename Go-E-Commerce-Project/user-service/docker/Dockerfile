FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy workspace files
COPY go.work go.work.sum* ./
COPY shared/ shared/
COPY user-service/ user-service/

# Build inside user-service
WORKDIR /app/user-service

RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Final stage - lightweight image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary and .env from builder stage
COPY --from=builder /app/user-service/main .
COPY user-service/.env .env

EXPOSE 8081

CMD ["./main"]