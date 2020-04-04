crypto.getRandomValues = (buf) => {
  buf = Buffer.from(buf);
  let bytes = crypto.randomBytes(buf.length);
  buf.set(bytes);
  return buf;
};
require('./wasm_exec');

const go = new Go();
const path = require('path').join(__dirname, 'konami-card.wasm');
const bytes = require('fs').readFileSync(path);
const wasmModule = new WebAssembly.Module(bytes);
const wasmInstance = new WebAssembly.Instance(wasmModule, go.importObject);
go.run(wasmInstance).then();

module.exports.encode = function (nfcId) {
  return new Promise((resolve) => {
    KONAMI_CARD_ENCODE(nfcId, function (result) {
      resolve(result);
    });
  });
};

module.exports.decode = function (cardId) {
  return new Promise((resolve) => {
    KONAMI_CARD_DECODE(cardId, function (result) {
      resolve(result);
    });
  });
};
