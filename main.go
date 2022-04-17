package main

import (
	"bytes"
	"io/ioutil"
	"syscall/js"

	"github.com/andybalholm/brotli"
)

// TODO: expose encode func
func encode(data []byte /* , options brotli.WriterOptions */) ([]byte, error) {
	var buf bytes.Buffer
	// writer := brotli.NewWriterOptions(&buf, options)
	writer := brotli.NewWriter(&buf)
	_, err := writer.Write(data)
	if closeErr := writer.Close(); err == nil {
		err = closeErr
	}
	return buf.Bytes(), err
}

func decode(encodedData []byte) ([]byte, error) {
	r := brotli.NewReader(bytes.NewReader(encodedData))
	return ioutil.ReadAll(r)
}

const ERR_PREFIX = "[BROCCOLI ERROR] "

func decodeWrapper(this js.Value, args []js.Value) interface{} {
	var encodedData []byte = make([]byte, args[0].Length())
	js.CopyBytesToGo(encodedData, args[0])
	data, err := decode(encodedData)
	if err != nil {
		return ERR_PREFIX + err.Error()
	}
	return string(data)
}

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("BroccoliDecode", js.FuncOf(decodeWrapper))
	<-done
}
