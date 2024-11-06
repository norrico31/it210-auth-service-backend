GO_CMD=go
SEED_CMD=cmd/seed

build:
	@go build -o bin/it210 cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/it210

run-dev:
	@go run ./cmd
	# @echo "Running in development mode with live reload"
	# Assuming air or a similar live reload tool is installed (you can modify this as needed)
	# @air

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-up-docker:
	@docker-compose run --rm auth-service make migrate-up

migrate-down:
	@go run cmd/migrate/main.go down


build-seeder:
	$(GO_CMD) build -o bin/seeder $(SEED_CMD)/main.go

seed-all: build-seeder
	./bin/seeder

run-docker-up:
	@docker-compose up -d --build

clean-image:
	@echo "Cleaning up Docker containers and volumes"
	@docker-compose down -v

clean:
	rm -rf bin