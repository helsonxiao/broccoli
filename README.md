# Broccoli

Golang wasm wrapper for [brotli](https://github.com/google/brotli) ;)

## Benchmark

👉 [Let's go!](https://helsonxiao.github.io/broccoli/benchmark)

⚠️ You probably DONT need a wasm wrapper for brotli. [Google's JS version](https://github.com/google/brotli/tree/master/js) is fast enough in the browser!

| URLSearchParams | Type                  | Note                          |
| --------------- | --------------------- | ----------------------------- |
| language        | `js`  / `go` / `rust` | Pure JS or wasm               |
| times           | number                | Number of decompression times |
| messages        | number                | Number of messages            |
