import "../go/wasm_exec.js";
// import "../go/wasm_exec.tinygo.js";
// https://github.com/google/brotli/blob/master/js/decode.js
import { BrotliDecode as jsBrotliDecode } from "./decode.js";
// build manually using https://github.com/httptoolkit/brotli-wasm
// https://github.com/helsonxiao/brotli-wasm/commit/e6e3bcc67ef855af0fa8e0b759699a3db38d83c8
import initRustBrotli, * as rustBrotli from "../rust/brotli_wasm.js";
// build manually using https://github.com/ustclug-dev/brotli-dec-wasm
// https://github.com/helsonxiao/brotli-dec-wasm/commit/dae05636b3dc7d3d671d6b60c8deed7153d73174
import initRustBrotliDec, * as rustBrotliDec from "../rust/brotli_dec_wasm.js";

const params = new URLSearchParams(location.search);
const language = params.get("language") || "js";
console.log("[language]", language);
const times = parseInt(params.get("times")) || 1000;
console.log("[times]", times);
const messages = parseInt(params.get("messages")) || 1;
console.log("[messages]", messages);

const textEncoder = new TextEncoder();
const textDecoder = new TextDecoder();

const msg =
  '{"type":"message","event":"new","room_id":262451765,"user":{"user_id":9804407,"username":"迎宾机器猫小小川ゞ","iconurl":"http://static.missevan.com/avatars/202112/08/e40e541df9ca26088e9ddf26f0a30595234301.png","titles":[{"type":"level","level":49},{"type":"username","color":"#FF8686","colors":"#FF8686"},{"type":"noble","name":"新秀","level":2},{"type":"medal","name":"彼岸","super_fan":{"expire_time":1648915200},"level":18}]},"msg_id":"34ea0c75-a3dd-4149-8bf7-4d4f0caec46c","message":"@孙夫人s 来啦~ 小心炸弹！(╯‵□′)╯•••*～●，哈 骗你的啦~　"}'.repeat(
    messages
  );

if (language === "js" || language === "go") {
  if (!WebAssembly.instantiateStreaming) {
    if (!WebAssembly.instantiate) {
      throw new Error("WebAssembly.instantiate is not available");
    }
    // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
      const source = await (await resp).arrayBuffer();
      return await WebAssembly.instantiate(source, importObject);
    };
  }

  const go = new Go();
  WebAssembly.instantiateStreaming(
    fetch("go/broccoli.wasm", {}),
    // fetch("go/broccoli.tinygo.wasm", {}),
    go.importObject
  ).then((result) => {
    go.run(result.instance);

    const encodedData = BroccoliEncode(msg);
    if (language === "js") {
      console.log(textDecoder.decode(jsBrotliDecode(encodedData)));
    } else if (language === "go") {
      console.log(textDecoder.decode(BroccoliDecode(encodedData)));
    }

    console.profile(language);
    console.time(language);
    if (language === "js") {
      for (let i = 0; i < times; i += 1) {
        jsBrotliDecode(encodedData);
      }
    } else if (language === "go") {
      for (let i = 0; i < times; i += 1) {
        BroccoliDecode(encodedData);
        // BroccoliDummyDecode(encodedData);
      }
    }
    console.timeEnd(language);
    console.profileEnd(language);
  });
} else if (language === "rust") {
  initRustBrotli().then(async () => {
    await initRustBrotliDec();

    const encodedData = rustBrotli.compress(textEncoder.encode(msg));
    console.log(textDecoder.decode(rustBrotli.decompress(encodedData)));
    console.log(textDecoder.decode(rustBrotliDec.brotliDec(encodedData)));

    console.profile(language);
    console.time(language);
    for (let i = 0; i < times; i += 1) {
      // rustBrotli.decompress(encodedData);
      rustBrotliDec.brotliDec(encodedData);
    }
    console.timeEnd(language);
    console.profileEnd(language);
  });
}
