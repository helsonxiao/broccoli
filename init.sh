cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
go get .
go install github.com/agnivade/wasmbrowsertest@latest
