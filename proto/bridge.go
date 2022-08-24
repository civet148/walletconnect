package proto

import (
	"encoding/hex"
	"encoding/json"
	"github.com/civet148/gotools/cryptos/goaes"
	_ "github.com/civet148/gotools/cryptos/goaes/cbc"
	"github.com/civet148/log"
)

const(
	PubType_Pub = "pub"
	PubType_Sub = "sub"
	PubType_Ack = "ack"
)

type SessionData struct {
	Topic   string `json:"topic"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Silent  bool   `json:"silent"`
}

func (m *SessionData) GetPayload() (*SessionPayload, error) {
	payload := &SessionPayload{}
	err := json.Unmarshal([]byte(m.Payload), payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

type SessionPayload struct {
	Data string `json:"data"`
	Hmac string `json:"hmac"`
	Iv   string `json:"iv"`
}

func (m *SessionPayload) GetBytesIV() []byte {
	iv, _ := hex.DecodeString(m.Iv)
	return iv
}

func (m *SessionPayload) GetIV() string {
	return m.Iv
}

func (m *SessionPayload) GetData() string {
	return m.Data
}

func (m *SessionPayload) GetHmac() string {
	return m.Hmac
}

func (m *SessionPayload) Decode(key []byte) ([]byte, bool) {
	aes := goaes.NewCryptoAES(goaes.AES_Mode_CBC, key, m.GetBytesIV())
	if aes == nil {
		log.Errorf("create AES-256-CBC object failed")
		return nil, false
	}
	data, err := aes.DecryptHex(m.Data)
	if err != nil {
		log.Errorf("decrypt payload error [%s]", err)
		return nil, false
	}
	return data, true
}

func (m *SessionPayload) DecodeByHexKey(strHexKey string) ([]byte, bool) {
	key, _ := hex.DecodeString(strHexKey)
	return m.Decode(key)
}
