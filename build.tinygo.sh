echo 'building wasm with tinygo...'
tinygo build -o ./static/go/broccoli.tinygo.wasm -target wasm ./main.go
echo 'done!'
