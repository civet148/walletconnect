package proto

//Subscribe request (Client -> Server)
type WakuSubscribeRequest struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"` //"2.0"
	Method  string `json:"method"`  //"waku_subscribe"
	Params  struct {
		Topic string `json:"topic"`
	} `json:"params"`
}

//Subscribe response (Server -> Client)
type WakuSubscribeResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"` //"2.0"
	Result  string `json:"result"`
}

//Publish request (Client -> Server)
type WakuPublishRequest struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"` //"2.0"
	Method  string `json:"method"`  //"waku_publish"
	Params  struct {
		Topic   string `json:"topic"`
		Message string `json:"message"`
		Ttl     int    `json:"ttl"`
	} `json:"params"`
}

//Publish response (Server -> Client)
type WakuPublishResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"` //"2.0"
	Result  bool   `json:"result"`  //true
}

//Subscription request (Server -> Client)
type WakuSubscriptionRequest struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"` //"2.0"
	Method  string `json:"method"`  //"waku_subscription"
	Params  struct {
		Id   string `json:"id"`
		Data struct {
			Topic   string `json:"topic"`
			Message string `json:"message"`
		} `json:"data"`
	} `json:"params"`
}

//Subscription response (Client -> Server)
type WakuSubscriptionResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"` //"2.0"
	Result  bool   `json:"result"`  //true
}
