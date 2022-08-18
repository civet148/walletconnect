package main

import (
	"github.com/civet148/log"
	"github.com/civet148/walletconnect"
	"github.com/civet148/walletconnect/proto"
)

//wc:5c4b71487691b4d6bc8aee257b34ac0769c6fb59bd0a63b04a1196e73d32d437@2?controller=false&publicKey=9d3afdc25d3244678a5d6ee4df887ce02f440ce8f8ca6264834f1b6d738f541f&relay=%7B%22protocol%22%3A%22waku%22%7D
func main() {
	client := wc.NewWCWallet()
	if !client.Init(&proto.ParamsInit{
		Controller: false,
		Metadata: proto.AppMetadata{
			Name:        "civet148",
			Description: "my app",
			Url:         "#",
			Icons:       nil,
		},
		RelayProvider: "wss://bridge.walletconnect.org",
	}) {
		log.Errorf("initiate connect failed")
		return
	}
}
