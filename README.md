# WalletConnect
golang wallet connect 

## 1. DAPP work flow

### 1.1 genrate wc URI and show QRCode for wallet

wc:{topic...}@{version...}?bridge={url...}&key={key...}

**topic** is a uuid string

```text
wc:9d9c764d-e6ca-4163-9320-779402e880d3@1?bridge=https%3A%2F%2Fy.bridge.walletconnect.org&key=5384d68cb83a80b8f1cce8ff37415a54d52a079b77940cb70b6dab2f12fc68f6
```

### 1.2 connect to bridge

DAPP -> bridge

connect to bridge by websocket

```http request
wss://0.bridge.walletconnect.org/?env=browser&host=app.uniswap.org&protocol=wc&version=1

Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,zh-TW;q=0.5
Cache-Control: no-cache
Connection: Upgrade
Host: c.bridge.walletconnect.org
Origin: https://app.uniswap.org
Pragma: no-cache
Sec-WebSocket-Extensions: permessage-deflate; client_max_window_bits
Sec-WebSocket-Key: 2HVyAUunvvZDCgqnPLZFJw==
Sec-WebSocket-Version: 13
Upgrade: websocket
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.47
```

### 1.3 pub/sub message 

- send topic and encrypted payload (AES-256-CBC + HMAC + IV)

```json
{
  "topic":"9d9c764d-e6ca-4163-9320-779402e880d3",
  "type":"pub",
  "payload":"{\"data\":\"42eb083fa4e59bc0e184aa3d938af7a01bbc943345f96bc4d17d39b267312a9362e27baa43a488d6ed08d7e29a6fe4bf63e4a1c19e06bfedee142f3e9bf6d1308871efeea93a8044814fb02c3581b80933b290ab730d2da978aee35d36ee2f9153736413b875ec9f8255245f33dbf8aed11e3c710d3b3de5c0445078c60377745fb56208838e43cb8564108c21380bb6d9a6beb1a97172a562dd07d2d0eeaeae2320115932518c63be41578d15c76e0ce2d53ad43974a00512a6e8c9e280e51a38da98534c034a31e789a2acc2bd6656a4972da81d8dcea83254160e39a46f39753fa450fc6cd68d0086116c8cb22a360cbe3598528b686865c59e33bd918301bcb6040959bac57fc612f6428d9e36a95cbc11223b055da7aa94571537f424ae44035bbd70cd6f8310f134b73e258ccc20716432589bd90e03a3818d8a5d9159ed7647c574180689c80b4639ab598ebab5fb2f2798f391f5080fa1493595ff936264b028c6c06a3c37c7e5f3135974b36a2d3fe85acdf497976a1856ece47f2578a3d0772faf6fcd3d47a20803f1d100d41697ba62927bfc58efa67e369530acde5a75fd53765d30e406fc4ca5557fcd2833835c8a2e0c49d2a9124b2f1bf13b\",\"hmac\":\"c8bf96199abe0059fd1a8869448516d871e56201e6759f71c03f39be2e2bbd2a\",\"iv\":\"0b0026d95c2b6704b9e02c9bd34dcfa0\"}",
  "silent":true
}
```

- payload decrypt
  
  mode: AES-256-CBC 

  key:  5384d68cb83a80b8f1cce8ff37415a54d52a079b77940cb70b6dab2f12fc68f6 (hex)

  iv:   0b0026d95c2b6704b9e02c9bd34dcfa0 (hex)

```json
{
  "id":1660726015060055,
  "jsonrpc":"2.0",
  "method":"wc_sessionRequest",
  "params":[
    {
      "peerId":"1d1208ac-03de-4a1e-825e-84433d3a90f2",
      "peerMeta":{
        "description":"Swap or provide liquidity on the Uniswap Protocol",
        "url":"https://app.uniswap.org",
        "icons":[
          "https://app.uniswap.org/./favicon.png",
          "https://app.uniswap.org/./images/192x192_App_Icon.png",
          "https://app.uniswap.org/./images/512x512_App_Icon.png"
        ],
        "name":"Uniswap Interface"
      },
      "chainId":1
    }
  ]
}
```

- subscribe events

DAPP -> bridge

```json
{
  "topic":"28f17b68-4878-4473-8bc3-15f6af81a773",
  "type":"sub",
  "payload":"",
  "silent":true
}
```
