.PHONY: all proto grpc backend frontend install clean help

## Default: show help
help:
	@echo ""
	@echo "  LoreKeeper — Makefile Commands"
	@echo ""
	@echo "  make install      Install all dependencies"
	@echo "  make proto        Generate gRPC code from .proto files"
	@echo "  make grpc         Start the Go gRPC RAG Engine (port 50051)"
	@echo "  make backend      Start the Node.js Backend (port 3000)"
	@echo "  make frontend     Start the Angular Frontend (port 4200)"
	@echo "  make dev          Print instructions to start all services"
	@echo "  make build        Build backend and frontend for production"
	@echo "  make clean        Remove build artifacts"
	@echo ""

## Install all dependencies
install: install-backend install-frontend
	@echo "All dependencies installed."

install-backend:
	@echo "Installing Backend dependencies..."
	cd backend && npm install

install-frontend:
	@echo "Installing Frontend dependencies..."
	cd frontend && npm install

## Download Go module dependencies
install-go:
	@echo "Downloading Go module dependencies..."
	cd grpc && go mod tidy

## Generate protobuf code
proto: proto-go proto-node

proto-go:
	@echo "Generating Go protobuf stubs..."
	mkdir -p grpc/pb
	protoc --go_out=grpc/pb --go_opt=paths=source_relative \
	       --go-grpc_out=grpc/pb --go-grpc_opt=paths=source_relative \
	       --proto_path=grpc/proto grpc/proto/rag_service.proto
	@echo "Go stubs generated in grpc/pb/"

proto-node:
	@echo "Generating Node.js protobuf stubs..."
	cd backend && npm run proto:gen
	@echo "Node.js stubs generated."

## Start services
grpc:
	@echo "Killing any existing process on port 50051..."
	-fuser -k 50051/tcp || true
	@echo "Starting RAG Engine (Go gRPC) on :50051..."
	cd grpc && go run main.go

backend:
	@echo "Killing any existing process on port 3000..."
	-fuser -k 3000/tcp || true
	@echo "Starting Backend (Node.js) on :3000..."
	cd backend && npm run dev

frontend:
	@echo "Killing any existing process on port 4200..."
	-fuser -k 4200/tcp || true
	@echo "Starting Frontend (Angular) on :4200..."
	cd frontend && npm start

## Dev mode - print instructions
dev:
	@echo ""
	@echo "  To start all LoreKeeper services, open 3 terminals:"
	@echo ""
	@echo "  Terminal 1 (gRPC RAG Engine):  make grpc"
	@echo "  Terminal 2 (Backend):          make backend"
	@echo "  Terminal 3 (Frontend):         make frontend"
	@echo ""
	@echo "  Then open: http://localhost:4200"
	@echo ""

## Build for production
build: build-backend build-frontend

build-backend:
	@echo "Building Backend..."
	cd backend && npm run build

build-frontend:
	@echo "Building Frontend..."
	cd frontend && npm run build

## Clean artifacts
clean:
	rm -rf backend/dist backend/node_modules
	rm -rf frontend/dist frontend/node_modules
	rm -rf grpc/pb/*.go
	@echo "Cleaned."
