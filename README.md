# eamuse-card-wasm
Convert e-amusement card ID and NFC ID to each other.

## Requirement
- go 1.14.2
- node.js

## Build
```shell script
npm install
npm run build
```

## Usage
```node
import KonamiCard from 'konami-card';

KonamiCard.encode(nfcId).then(result => console.info(result));
KonamiCard.decode(cardId).then(result => console.info(result));
```
