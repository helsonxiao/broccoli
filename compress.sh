echo 'building wasm...'
GOOS=js GOARCH=wasm go build -o ./static/broccoli.wasm
echo 'compressing wasm...'
rm ./static/broccoli.br.wasm
brotli -o ./static/broccoli.br.wasm ./static/broccoli.wasm
echo 'done!'
