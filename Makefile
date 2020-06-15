GO=go
BUILD_DIR=./build

.PHONY: server client

# By default binaries are build for current host architecture
all: server client

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

server: $(BUILD_DIR)
	$(ARCHSET) $(GO) build -o $(BUILD_DIR)/rlserver$(SUFFIX) ./cmd/server

client: $(BUILD_DIR)
	$(ARCHSET) $(GO) build -o $(BUILD_DIR)/rlclient$(SUFFIX) ./cmd/client

# Targets for cross compilation
intel32: ARCHSET=GOARCH=386
intel32: SUFFIX=-x86
intel32: all

amd64: ARCHSET=GOARCH=amd64
amd64: SUFFIX=-x64
amd64: all

pizero: ARCHSET=GOARCH=arm GOARM=6 GOOS=linux
pizero: SUFFIX=-armv6
pizero: all

pi3: ARCHSET=GOARCH=arm GOARM=7 GOOS=linux
pi3: SUFFIX=-armv7
pi3: all
