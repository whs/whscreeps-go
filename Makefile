BRANCH ?= sim
TOKEN ?= dd7de6b8-17dd-4dac-8f8a-5d99f2c59237
SOURCES = $(shell find . -name "*.go")

.PHONY: deploy

deploy: cmd/deploy/deploy deploy/runtime/main.js main.wasm
	@cmd/deploy/deploy -token $(TOKEN) -branch $(BRANCH) deploy/runtime/main.js main.wasm

cmd/deploy/deploy: cmd/deploy/main.go
	cd cmd/deploy && go build -o deploy ./...

main.wasm: $(SOURCES) go.mod go.sum
	#GOOS=js GOARCH=wasm go build -o $@
	tinygo build --no-debug -o $@ -target wasm main.go

deploy/runtime/main.js: deploy/runtime/wasm_exec_tinygo.js deploy/runtime/bootloader.js
	npx esbuild deploy/runtime/bootloader.js --bundle --external:main.wasm \
		--outfile=$@ --target=node12 --format=cjs --define:TINYGO=true
