# Use the official Go image for building the application
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /app

# Install git (if needed)
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application
COPY . .

# Clean up go.mod and go.sum
RUN go mod tidy

# Build the Go applications
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/migrate ./cmd/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/seed ./cmd/seed/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/auth-service ./cmd/main.go

# Use a lightweight image for the final stage
FROM alpine:latest
RUN apk add --no-cache netcat-openbsd
RUN apk add --no-cache libc6-compat

# Set the working directory for the final image
WORKDIR /app/auth-service

# Copy the binaries from the builder stage
COPY --from=builder /app/auth-service /app/auth-service/
COPY --from=builder /app/migrate /app/auth-service/migrate
COPY --from=builder /app/seed /app/auth-service/seed
COPY entrypoint.sh /app/auth-service/

# Copy the migration SQL files
COPY --from=builder /app/cmd/migrate/migrations /app/auth-service/cmd/migrate/migrations

RUN chmod +x /app/auth-service/migrate /app/auth-service/seed /app/auth-service/auth-service

EXPOSE 8081

ENTRYPOINT ["/app/auth-service/entrypoint.sh"]
