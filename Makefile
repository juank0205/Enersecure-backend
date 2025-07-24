APP_NAME := app
CMD_DIR := ./cmd/$(APP_NAME)
BIN_DIR := ./bin
BIN := $(BIN_DIR)/$(APP_NAME)

.PHONY: all build run test clean

all: build

build:
	@echo "Building..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN) $(CMD_DIR)

run: build
	@echo "Running"
	$(BIN)

test:
	@echo "Running tests..."
	go test ./...
	

clean:
	@echo "Cleaning..."
	rm -rf $(BIN_DIR)
