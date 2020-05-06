"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const crypto_1 = __importDefault(require("crypto"));
crypto_1.default['getRandomValues'] = (buf) => {
    buf = Buffer.from(buf);
    let bytes = crypto_1.default.randomBytes(buf.length);
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
exports.encode = function (nfcId) {
    return new Promise((resolve) => {
        // @ts-ignore
        KONAMI_CARD_ENCODE(nfcId, function (result) {
            resolve(result);
        });
    });
};
exports.decode = function (cardId) {
    return new Promise((resolve) => {
        // @ts-ignore
        KONAMI_CARD_DECODE(cardId, function (result) {
            resolve(result);
        });
    });
};
exports.default = { encode: exports.encode, decode: exports.decode };
//# sourceMappingURL=konami-card.js.map