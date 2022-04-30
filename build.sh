echo 'building wasm...'
GOOS=js GOARCH=wasm go build -o ./static/go/broccoli.wasm
echo 'done!'
