echo 'building wasm...'
GOOS=js GOARCH=wasm go build -o ./static/go/broccoli.wasm
echo 'compressing wasm...'
rm ./static/go/broccoli.br.wasm
brotli -o ./static/go/broccoli.br.wasm ./static/go/broccoli.wasm
echo 'done!'
