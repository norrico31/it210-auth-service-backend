FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/migrate ./cmd/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/seed ./cmd/seed/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/auth-service ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache netcat-openbsd
RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/auth-service /app/
COPY --from=builder /app/migrate /app/
COPY --from=builder /app/seed /app/
COPY entrypoint.sh /app/

# Copy the migration SQL files
COPY --from=builder /app/cmd/migrate/migrations /app/cmd/migrate/migrations

RUN chmod +x /app/migrate /app/seed /app/auth-service

EXPOSE 8081

ENTRYPOINT ["/app/entrypoint.sh"]
