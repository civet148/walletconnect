package walletconnect

import (
	"github.com/civet148/jsonrpc"
	"github.com/civet148/log"
)

type DApp struct {
	wc string
	uri *WalletConnectURI
	ws *jsonrpc.WebSocketClient
}

// NewDApp creates a new DApp with the given wc uri
func NewDApp(wc string) *DApp {
    uri, err := ParseWC(wc)
    if err != nil {
    	log.Errorf(err.Error())
    	return nil
	}
	return &DApp{
		wc: wc,
		uri: uri,
		ws: jsonrpc.NewWebSocketClient(uri.Bridge, nil),
	}
}

func (d *DApp) Publish() error {

	return nil
}


func (d *DApp) Subscribe(topic string) error {

	//d.ws.Subscribe()
	return nil
}
