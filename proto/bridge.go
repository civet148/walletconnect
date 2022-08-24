package proto

import "encoding/json"

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