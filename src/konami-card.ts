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

declare class Go {
  public importObject: any;
  run(instance: WebAssembly.Instance): Promise<any>
}

require('./wasm_exec');

const go = new Go();
const path = require('path').join(__dirname, 'konami-card.wasm');
const bytes = require('fs').readFileSync(path);
const wasmModule = new WebAssembly.Module(bytes);
const wasmInstance = new WebAssembly.Instance(wasmModule, go.importObject);
go.run(wasmInstance).then();

declare global {
  namespace NodeJS {
    interface Global {
      __konami_card_encode(nfcId: string, func: (result: string) => void): void;
      __konami_card_decode(cardId: string, func: (result: string) => void): void;
    }
  }
}

export const encode = function(nfcId: string): Promise<string> {
  return new Promise((resolve) => {
    global.__konami_card_encode(nfcId, function(result: string) {
      resolve(result);
    });
  });
};

export const decode = function(cardId: string): Promise<string> {
  return new Promise((resolve) => {
    global.__konami_card_decode(cardId, function(result: string) {
      resolve(result);
    });
  });
};

export default { encode, decode };
