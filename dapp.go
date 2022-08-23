package walletconnect

//import "github.com/civet148/jsonrpc"

type DApp struct {
	wc string
}

// NewDApp creates a new DApp with the given wc uri
func NewDApp(wc string) *DApp {

	return &DApp{
		wc: wc,
		//relay: jsonrpc.NewRelayClient(),
	}
}

func (d *DApp) Publish() error {

	return nil
}


func (d *DApp) Subscribe() error {

	return nil
}
