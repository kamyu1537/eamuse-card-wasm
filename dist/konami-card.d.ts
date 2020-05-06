/// <reference types="node" />
declare module 'crypto' {
    function getRandomValues(buf: Array<number> | Uint8Array | Buffer): Uint8Array;
}
export declare const encode: (nfcId: string) => Promise<unknown>;
export declare const decode: (cardId: string) => Promise<unknown>;
declare const _default: {
    encode: (nfcId: string) => Promise<unknown>;
    decode: (cardId: string) => Promise<unknown>;
};
export default _default;
//# sourceMappingURL=konami-card.d.ts.map