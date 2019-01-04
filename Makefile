APPLICATION := $(lastword $(subst /, ,$(dir $(CURDIR))))
BIN := $(value GOPATH)\bin\$(APPLICATION).exe

# build when changed
$(BIN): *.go
	go build -o $(BIN)

# build and install application
install: $(BIN)

# remove application
clean: 
	@go clean -i

.PHONY: clean install
