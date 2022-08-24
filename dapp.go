package walletconnect

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/civet148/jsonrpc"
	"github.com/civet148/log"
	"github.com/civet148/gotools/cryptos/goaes"
	_ "github.com/civet148/gotools/cryptos/goaes/cbc"
	"github.com/civet148/walletconnect/proto"
)

type CallbackFn func(c context.Context, msg []byte) bool

type DApp struct {
	wc string
	cb CallbackFn
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

func (d *DApp) PublishTopic(topic, pubType string, data proto.SessionData) error {

	return nil
}


func (d *DApp) SubscribeTopic(topic string, cb CallbackFn) error {
	d.cb = cb
	var req = proto.SessionData{
		Topic:   topic,
		Type:    proto.PubType_Sub,
		Payload: "",
		Silent:  true,
	}
	log.Infof("subscribed to topic [%s]", topic)
	data, err := json.Marshal(&req)
	if err != nil {
		return log.Errorf(err.Error())
	}
	log.Infof("subscribed to topic [%s] data [%s]", topic, data)
	err =  d.ws.Subscribe(context.Background(), string(data), d.subscriber)
	if err != nil {
		return log.Errorf("subscribe error [%s]", err.Error())
	}
	return nil
}

func (d *DApp) subscriber(ctx context.Context, msg []byte) bool {
	log.Infof("subscriber: [%s]", string(msg))
	var sd = &proto.SessionData{}
	err := json.Unmarshal(msg, sd)
	if err != nil {
		log.Errorf("unmarshal [%s] to session data error [%s]", msg, err)
		return false
	}
	payload, err := sd.GetPayload()
	if err != nil {
		log.Errorf("unmarshal payload error [%s]", err)
		return false
	}
	key, _ := hex.DecodeString(d.uri.Key)
	iv, _ := hex.DecodeString(payload.Iv)
	aes := goaes.NewCryptoAES(goaes.AES_Mode_CBC, key, iv)
	if aes == nil {
		log.Errorf("create AES-256-CBC object error [%s]", err)
		return false
	}
	data, err := aes.DecryptHex(payload.Data)
	if err != nil {
		log.Errorf("decrypt payload error [%s]", err)
		return false
	}
	if !d.cb(ctx, data) {
		return false
	}
	return true
}

