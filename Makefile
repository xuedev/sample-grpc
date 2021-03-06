BINARY := grpc

all: build run

build:
	@echo "==> Go build"
	@go build -o $(BINARY)

run:
	@echo "==> Running"
	@./$(BINARY) print

.PHONY: build run