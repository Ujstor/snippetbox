# Simple Makefile for a Go project

# Build the application
all: build test
tailwind:
	@if [ ! -f tailwindcss ]; then curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss; fi
	@chmod +x tailwindcss

build: tailwind
	@echo "Building..."
	@./tailwindcss -i ui/static/css/input.css -o ui/static/css/output.css
	@go build -o main cmd/web/main.go

# Run the application
run:
	@go run cmd/web/main.go
# Create DB container
p docker-run-prod:
	@if docker compose up app-prod --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up app-prod --build; \
	fi

d docker-run-dev:
	@if docker compose up app-dev --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up app-dev --build; \
	fi

# Shutdown DB container
c docker-down:
	@if docker compose down --volumes 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down --volumes; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch tailwind docker-run docker-down itest

