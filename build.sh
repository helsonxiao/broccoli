echo 'building wasm...'
GOOS=js GOARCH=wasm go build -o ./static/broccoli.wasm
echo 'done!'
