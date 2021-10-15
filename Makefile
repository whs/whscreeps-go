TOKEN ?= dd7de6b8-17dd-4dac-8f8a-5d99f2c59237
SOURCES = $(shell find . -name "*.go")

.PHONY: deploy

deploy: cmd/deploy/deploy deploy/runtime/main.js main.wasm deploy/runtime/wasm_exec.js
	@cmd/deploy/deploy -token $(TOKEN) deploy/runtime/main.js main.wasm deploy/runtime/wasm_exec.js

cmd/deploy/deploy: cmd/deploy/main.go
	cd cmd/deploy && go build -o deploy ./...

main.wasm: $(SOURCES) go.mod go.sum
	GOOS=js GOARCH=wasm go build -o main.wasm
