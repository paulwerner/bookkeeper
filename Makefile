.PHONY: all
all: build
FORCE: ;

SHELL := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf $(BIN_DIR)/*

depdendencies:
	go mod download

build: depdendencies
	go build -tags $(LIBRARY_ENV) -o $(BIN_DIR)/bookkeeper-api cmd/bookkeeper-api/main.go

ci: depdendencies test

test:
	SERVER_HOST=localhost DB_HOST=localhost go test -p 1 -tags testing ./...

run: 
	SERVER_HOST=localhost DB_HOST=localhost go run $(PWD)/cmd/bookkeeper-api/main.go

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
