package walletconnect

import (
	"context"
	"encoding/json"
	"github.com/civet148/jsonrpc"
	"github.com/civet148/log"
	"github.com/civet148/walletconnect/proto"
)

type CallbackFunc func(c context.Context, topic, typ string, msg []byte) bool

type DApp struct {
	wc  string
	cb  CallbackFunc
	uri *WalletConnectURI
	ws  *jsonrpc.WebSocketClient
}

// NewDApp creates a new DApp with the given wc uri
func NewDApp(wc string) (*DApp, error) {
	uri, err := ParseWC(wc)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	ws, err := jsonrpc.NewWebSocketClient(uri.Bridge, nil)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return &DApp{
		wc:  wc,
		uri: uri,
		ws:  ws,
	}, nil
}

func (d *DApp) PublishTopic(sd proto.SessionData) error {
	data, err := json.Marshal(sd)
	if err != nil {
		return log.Errorf("unmarshal error [%s]", err)
	}
	return d.ws.Send(data)
}

func (d *DApp) SubscribeTopic(topic string, cb CallbackFunc) error {
	d.cb = cb
	var req = proto.SessionData{
		Topic:   topic,
		Type:    proto.PubType_Sub,
		Payload: "",
		Silent:  true,
	}
	data, err := json.Marshal(&req)
	if err != nil {
		return log.Errorf(err.Error())
	}
	log.Debugf("subscribed to topic [%s] data [%s]", topic, data)
	err = d.ws.Subscribe(context.Background(), data, d.subscriber)
	if err != nil {
		return log.Errorf("subscribe error [%s]", err.Error())
	}
	return nil
}

func (d *DApp) Close() {
	d.ws.Close()
}

func (d *DApp) subscriber(ctx context.Context, msg []byte) bool {
	log.Debugf("subscriber: [%s]", string(msg))
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
	data, ok := payload.DecodeByHexKey(d.uri.Key)
	if !ok {
		log.Errorf("decrypt payload failed")
		return false
	}
	if !d.cb(ctx, sd.Topic, sd.Type, data) {
		return false
	}
	return true
}
