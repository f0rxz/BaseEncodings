GO = go
GOTEST = go test
RUN = go run
SRC_DIRS = ./base16util ./base32util ./base64util ./encoder ./internal

ENCODING ?= base64

.PHONY: all
all: run

test:
	$(GOTEST) ./...

run: test
	$(RUN) main.go -encoding=$(ENCODING)

build:
	$(GO) build -o myapp main.go
