package main

import (
	"github.com/civet148/log"
	"github.com/civet148/walletconnect/proto"
	"github.com/civet148/walletconnect/rpc"
)

//https://relay.walletconnect.com/?apiKey=176974580ede0fe0f6cc6b5fed0bd307

func main() {
	//qrcode.QrencodeTerminal("https://relay.walletconnect.com/?apiKey=176974580ede0fe0f6cc6b5fed0bd307", qrencode.ECLevelL)
	ch := make(chan interface{})

	wc := rpc.NewWCClient()
	if !wc.Init(&proto.ParamsInit{
		Controller: false,
		Metadata: proto.AppMetadata{
			Name:        "civet148",
			Description: "my app",
			Url:         "#",
			Icons:       nil,
		},
		RelayProvider: "wss://relay.walletconnect.com",
	}) {
		log.Errorf("initiate connect failed")
		return
	}
	log.Debugf("connect ok")
	<-ch
	log.Infof("program exit...")
}
