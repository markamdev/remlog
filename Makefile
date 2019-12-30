GO=go
BUILD_DIR=./build

all: server tester

outdir:
	@echo "-- OUTDIR --"
	@mkdir -p $(BUILD_DIR)

clean:
	@echo "-- CLEAN --"
	@rm -rf $(BUILD_DIR)

server: outdir
	@echo "-- SERVER --"
	@$(GO) build -o $(BUILD_DIR)/rlserver ./cmd/server

tester: outdir
	@echo "-- TESTER --"
	@$(GO) build -o $(BUILD_DIR)/rltester ./cmd/tester
