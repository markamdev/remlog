GO=go
BUILD_DIR=./build

all: server tester

outdir:
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

server: outdir
	$(GO) build -o $(BUILD_DIR)/rlserver ./cmd/server

tester: outdir
	$(GO) build -o $(BUILD_DIR)/rltester ./cmd/tester
