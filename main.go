package main

import (
	"bytes"
	"io/ioutil"
	"syscall/js"

	"github.com/andybalholm/brotli"
)

const ERR_PREFIX = "[BROCCOLI ERROR] "

// TODO: expose options
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

func encodeWrapper(this js.Value, args []js.Value) interface{} {
	data := args[0].String()
	encodedData, err := encode([]byte(data))
	if err != nil {
		return ERR_PREFIX + err.Error()
	}
	buffer := js.Global().Get("Uint8Array").New(len(encodedData))
	js.CopyBytesToJS(buffer, encodedData)
	return buffer
}

func decodeWrapper(this js.Value, args []js.Value) interface{} {
	encodedData := make([]byte, args[0].Length())
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
	global.Set("BroccoliEncode", js.FuncOf(encodeWrapper))
	global.Set("BroccoliDecode", js.FuncOf(decodeWrapper))
	<-done
}
