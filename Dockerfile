FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

RUN go mod tidy

# Build first stage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/auth-service ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/auth-service /app/auth-service

RUN chmod +x /app/auth-service

# Expose the ports used by the application
EXPOSE 81 8081

# Specify the default command to run the main service
CMD ["/app/auth-service"]
