# WalletConnect
golang wallet connect 

wc:{topic...}@{version...}?bridge={url...}&key={key...}

**topic** is a uuid string


- payload decrypt

  mode: AES-256-CBC

  key:  d8089b579aeef1ba4692d72bd3a0bb870f8bbceafaab325f23739635e2e54f05 (hex)

  iv:   1f432ae9451cc4e27f290ba891426350 (hex)


workflow example
- wc URI (QRCode)

```text
wc:38d80e38-1c9a-470f-aed1-27b8ced3338d@1?bridge=https%3A%2F%2F6.bridge.walletconnect.org&key=d8089b579aeef1ba4692d72bd3a0bb870f8bbceafaab325f23739635e2e54f05
```

- DApp -> bridge (DApp publish)

```json
{"topic":"38d80e38-1c9a-470f-aed1-27b8ced3338d","type":"pub","payload":"{\"data\":\"92807450bee1ecc14432ade00da18f73e97f37859318264f4c2bdc5ee3450559215965e6e8ae5919c04710a138983d54f037df6fb077b6638279c2b84a6f5708572ea6c38e6ddb3669ac06656dfdb75303f11717c9dd4348aa75455338b72cc4abfca9d92bd6ab47dcaa88001e9c3caa7a800d48742a873c33e4533ac934ca630e30281f8fc6e70edcc7e0dcf4b2e55e1d47ffa781dd2544963e9a81441ca17e456714e18f83cfcf2bd13a37a26877df7ec73ffba1451c614e28aff263b88398f192a6e3fc44520dad78f4b175ef488a774b280d2f249aae73a55edfbfbcf611fc431d78a00f73fd8867c06f5381516b84315c913e2df549a70907348c2db4e6b5d3db5fe172734ad83f82858552067ec426bb6afc0990fb9002dcc2d48d1b559960b49e5284b8041e579c0ce3fcf66588002d70f4fb9a22986300e820b2266ac597741436bbefb2a47bdae46c29e67916a98f1171832a4dbefd14cd3059d9465988dee81c362b775182e9bae51a2edc6c5da299e41fad3d49c6d42e3255b45c17719e055b582285b4dadf9c750ddd09e49b337c61619d1f0d4b7caee9671934042be32cd5831b48c2ec8fa4835450b3b50a20aad6c062ce1ec2b682431e78db\",\"hmac\":\"b39735dcc4a57d92f749b4376837fbc3c0cabae38834605d69c6a173f9ccc3fa\",\"iv\":\"1f432ae9451cc4e27f290ba891426350\"}","silent":true}
```
payload
```json
{
    "id":1660732761714425,
    "jsonrpc":"2.0",
    "method":"wc_sessionRequest",
    "params":[
        {
            "peerId":"90e583f2-9078-4f0a-8ace-fb6265d47f50",
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

- DApp -> bridge (subscribe events)

```json
{"topic":"90e583f2-9078-4f0a-8ace-fb6265d47f50","type":"sub","payload":"","silent":true}
```

- bridge -> DApp (wallet approve)

```json
{"topic":"90e583f2-9078-4f0a-8ace-fb6265d47f50","type":"pub","payload":"{\"data\":\"c5a52ca478357941e6d0cb41b35cbd83f2aa86f62b9a328450dc5867998271187f3efd1a39cafacc318ebf886156f64748edd23dbef3c393abeffeafa33dcb89b51c863720c7baee902c22edccf3ce750d4c28ddc3bee589556eaa9410c964d0ca854270eba724db59106eddf80680308b25a0d4f7322d7cde6729e6b26707235f2393826f1d391ae8639ea67f1777571692010cce80517c620331bcbabf9c294ab805bb7bf7c353903c9be665f54d69f5cac666ff6773b3ab4ff6ed1e63ba9bfc3115878ccb72a9d75abaddbfe5d037ce90f8b19fc04ef422e3dda9fd4672d3089091dd95cd0ca0cde8dd1b625fa790fcb57dd6573b3e0de9d73c556f38d5702d7e28112e7dc573608be5ea33d3b120d976f6b88947d23be00ad7e76247872d7bf3b10ee0cc7bdb8cfb7fe32c7b2e1e7ad99b34438fa4f9e72ff1fc2e67923dd09ff52e7acb417357423b6f5558830780354f4d994698575763e288c21bab809527fea0c7162069270c55d83ecdfc6e\",\"hmac\":\"c6bcf7ccbe4c59bc0c951c08db3b2444a2bb59707e6de64ed37b320e8ab15956\",\"iv\":\"88a091f992a62d14ee5a083b793d7721\"}","silent":true}
```
payload
```json
{
  "id":1660732761714425,
  "jsonrpc":"2.0",
  "result":{
    "approved":true,
    "chainId":1,
    "networkId":0,
    "accounts":[
      "0x81A81D8Eef3d7E6023F3e7339A626C2EBE1A39d9"
    ],
    "rpcUrl":"",
    "peerId":"31038ba4-7d48-4792-9e5b-9cd45ca5fb9a",
    "peerMeta":{
      "description":"imToken wallet APP",
      "url":"https://token.im",
      "icons":[
        "https://token.im/images/download/appLogo.svg"
      ],
      "name":"imToken"
    }
  }
}
```

- DApp -> bridge (dapp ack)
 
```json
{"topic":"90e583f2-9078-4f0a-8ace-fb6265d47f50","type":"ack","payload":"","silent":true}
```

- bridge -> DApp (wallet disconnect)

```json
{"topic":"90e583f2-9078-4f0a-8ace-fb6265d47f50","type":"pub","payload":"{\"data\":\"cdc4a90b7783915eeadca95d837e352d8c7781cc6a26120130154d30640108a79e238bf5d03b16c214f725c5d37344ebec31eec3672688c461259f84aa76ca84cd64e108884feac1c33de7cf74c9a826cd5d199a1d8e2a0c361f467a2668d84f2990c3da94cd629bce609b6ef219c7a2392525c5ebeee2978ff5edd96f4344206ac7cec7ef7ed30ec49feca08a907ebd4335135e93ed1568a103287a471b6775\",\"hmac\":\"7d66dd9da8cda7cdf59adf9fef3fdc582e9a95a1b149d60ff8a10df2eea0bc4d\",\"iv\":\"624614af353b33f90c4d60aed31aa03e\"}","silent":true}
```

payload
```json
{
  "id":1660732892855299,
  "jsonrpc":"2.0",
  "method":"wc_sessionUpdate",
  "params":[
    {
      "approved":false,
      "chainId":null,
      "networkId":null,
      "accounts":null
    }
  ]
}
```

- DApp -> bridge (dapp ack)

```json
{"topic":"90e583f2-9078-4f0a-8ace-fb6265d47f50","type":"ack","payload":"","silent":true}
```

