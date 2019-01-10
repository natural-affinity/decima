APPLICATION := $(lastword $(subst /, ,$(dir $(CURDIR))))
PACKAGE := $(shell go list)/...
SOURCE := $(patsubst %_test.go, %.go, $(wildcard *_test.go **/*_test.go))
BIN := $(value GOPATH)\bin\$(APPLICATION).exe

# build when changed
$(BIN): $(SRC)
	go build -o $(BIN)

# run all tests and rebuild when changed
test: $(BIN)
	@go test $(PACKAGE)

# build and install application
install: $(BIN)

# remove application
clean: 
	@go clean -i

.PHONY: clean test install
