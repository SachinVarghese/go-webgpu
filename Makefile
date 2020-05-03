build_wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

exec_wasm:
	GOOS=js GOARCH=wasm go run -exec="$$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .	

test_wasm:
	GOOS=js GOARCH=wasm go test -exec="$$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .		