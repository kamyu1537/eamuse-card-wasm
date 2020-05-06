import crypto from 'crypto';

declare module 'crypto' {
  function getRandomValues(buf: Array<number> | Uint8Array | Buffer): Uint8Array;
}

crypto['getRandomValues'] = (buf: Array<number> | Uint8Array | Buffer): Uint8Array => {
  buf = Buffer.from(buf);
  let bytes = crypto.randomBytes(buf.length);
  buf.set(bytes);
  return buf;
};

require('./wasm_exec');

//@ts-ignore
const go = new Go();
const path = require('path').join(__dirname, 'konami-card.wasm');
const bytes = require('fs').readFileSync(path);
const wasmModule = new WebAssembly.Module(bytes);
const wasmInstance = new WebAssembly.Instance(wasmModule, go.importObject);
go.run(wasmInstance).then();

export const encode = function(nfcId: string) {
  return new Promise((resolve) => {
    // @ts-ignore
    KONAMI_CARD_ENCODE(nfcId, function(result: string) {
      resolve(result);
    });
  });
};

export const decode = function(cardId: string) {
  return new Promise((resolve) => {
    // @ts-ignore
    KONAMI_CARD_DECODE(cardId, function(result: string) {
      resolve(result);
    });
  });
};

export default { encode, decode };
