package walletconnect

import (
	"encoding/hex"
	"github.com/civet148/log"
	"net/url"
	"strings"
)

type WalletConnectURI struct {
	Scheme  string `json:"scheme"`
	Version string `json:"version"`
	Bridge  string `json:"bridge"`
	Topic   string `json:"topic"`
	Key     string `json:"key"`
}

//wc:38d80e38-1c9a-470f-aed1-27b8ced3338d@1?bridge=https://6.bridge.walletconnect.org&key=d8089b579aeef1ba4692d72bd3a0bb870f8bbceafaab325f23739635e2e54f05
func ParseWC(wc string) (*WalletConnectURI, error) {
	if !strings.HasPrefix(wc, "wc:") {
		return nil, log.Errorf("wc uri invalid")
	}
	wc = strings.Replace(wc, "wc:", "wc://", -1)
	u, err := url.Parse(wc)
	if err != nil {
		return nil, log.Errorf("wc uri parse error: %s", err)
	}
	params := u.Query()
	var bridge, key string
	if vs, ok := params["bridge"]; ok {
		bridge = vs[0]
	}
	if vs, ok := params["key"]; ok {
		key = vs[0]
	}
	return &WalletConnectURI{
		Scheme:  u.Scheme,
		Version: u.Host,
		Bridge:  bridge,
		Topic:   u.User.Username(),
		Key:     key,
	}, nil
}

func (m *WalletConnectURI) GetBytesKey() []byte {
	key, _ := hex.DecodeString(m.Key)
	return key
}