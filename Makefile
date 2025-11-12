.PHONY: build build-alpine run docker-build docker-run docker-stop docker-clean help

# Variables
IMAGE_NAME ?= vkr-backend
IMAGE_TAG ?= latest
CONTAINER_NAME ?= vkr-backend

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application locally
	@echo "Building application..."
	CGO_ENABLED=0 go build -ldflags="-w -s" -trimpath -o bin/vkr ./cmd/vkr
	@echo "Build complete: bin/vkr"

build-alpine: ## Build the application for Alpine Linux
	@echo "Building application for Alpine..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -trimpath -o bin/vkr-linux ./cmd/vkr
	@echo "Build complete: bin/vkr-linux"

run: ## Run the application locally
	@echo "Running application..."
	go run ./cmd/vkr

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	@echo "Docker image built: $(IMAGE_NAME):$(IMAGE_TAG)"

docker-build-alpine: ## Build Docker image using Alpine
	@echo "Building Docker image with Alpine..."
	docker build -f Dockerfile.alpine -t $(IMAGE_NAME):$(IMAGE_TAG)-alpine .
	@echo "Docker image built: $(IMAGE_NAME):$(IMAGE_TAG)-alpine"

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -d \
		--name $(CONTAINER_NAME) \
		-p 8081:8081 \
		--env-file .env \
		$(IMAGE_NAME):$(IMAGE_TAG)
	@echo "Container started: $(CONTAINER_NAME)"

docker-compose-up: ## Start services with docker-compose
	@echo "Starting services with docker-compose..."
	docker-compose up -d
	@echo "Services started"

docker-compose-down: ## Stop services with docker-compose
	@echo "Stopping services..."
	docker-compose down
	@echo "Services stopped"

docker-stop: ## Stop Docker container
	@echo "Stopping container..."
	docker stop $(CONTAINER_NAME) || true
	@echo "Container stopped"

docker-rm: ## Remove Docker container
	@echo "Removing container..."
	docker rm $(CONTAINER_NAME) || true
	@echo "Container removed"

docker-clean: docker-stop docker-rm ## Stop and remove Docker container
	@echo "Cleaned up container"

docker-logs: ## Show Docker container logs
	docker logs -f $(CONTAINER_NAME)

docker-shell: ## Open shell in Docker container (Alpine only)
	docker exec -it $(CONTAINER_NAME) /bin/sh

test: ## Run tests
	@echo "Running tests..."
	go test ./...

lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run ./...

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

mod-tidy: ## Tidy go modules
	@echo "Tidying modules..."
	go mod tidy

mod-download: ## Download go modules
	@echo "Downloading modules..."
	go mod download

