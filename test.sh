GOOS=js GOARCH=wasm go test -exec="$(go env GOPATH)/bin/wasmbrowsertest"
