cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static/go
go get .
go install github.com/agnivade/wasmbrowsertest@latest
