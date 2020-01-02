GO=go
BUILD_DIR=./build

# By default binaries are build for current host architecture
all: server tester

outdir:
	@echo "-- OUTDIR --"
	@mkdir -p $(BUILD_DIR)

clean:
	@echo "-- CLEAN --"
	@rm -rf $(BUILD_DIR)

server: outdir
	@echo "-- SERVER --"
	@$(ARCHSET) $(GO) build -o $(BUILD_DIR)/rlserver$(SUFFIX) ./cmd/server

tester: outdir
	@echo "-- TESTER --"
	@$(ARCHSET) $(GO) build -o $(BUILD_DIR)/rltester$(SUFFIX) ./cmd/tester

# Targets for cross compilation
intel32: ARCHSET=GOARCH=386
intel32: SUFFIX=-x86
intel32: .crossinfo
intel32: all

amd64: ARCHSET=GOARCH=amd64
amd64: SUFFIX=-x64
amd64: .crossinfo
amd64: all

pizero: ARCHSET=GOARCH=arm GOARM=6 GOOS=linux
pizero: SUFFIX=-armv6
pizero: .crossinfo
pizero: all

pi3: ARCHSET=GOARCH=arm GOARM=7 GOOS=linux
pi3: SUFFIX=-armv7
pi3: .crossinfo
pi3: all

.crossinfo:
	@echo 'Forced cross-compilation for $(SUFFIX)'