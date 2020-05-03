build_wasm:
	GOOS=js GOARCH=wasm go build -o ./web/main.wasm

exec_wasm:
	GOOS=js GOARCH=wasm go run -exec="$$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .	

test_wasm:
	GOOS=js GOARCH=wasm go test -exec="$$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .	

serve_wasm:
	go run tools/serve.go
