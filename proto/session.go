package proto

type Sequence struct {
	Topic string `json:"topic"`
}

type AppMetadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Icons       []string `json:"icons"`
}

type SessionPermissions struct {
	Blockchain struct {
		Chains []string `json:"chains"`
	} `json:"blockchain"`
	Jsonrpc struct {
		Methods []string `json:"methods"`
	} `json:"jsonrpc"`
}

type SessionProposal struct {
	Topic string `json:"topic"`
	Relay struct {
		Protocol string      `json:"protocol"`
		Params   interface{} `json:"params"`
	} `json:"relay"`
	Proposer struct {
		PublicKey string      `json:"public_key"`
		Metadata  AppMetadata `json:"metadata"`
	} `json:"proposer"`
	Signal struct {
		Method string   `json:"method"` //"pairing"
		Params Sequence `json:"params"`
	} `json:"signal"`
	Permissions SessionPermissions `json:"permissions"`
	Ttl         int                `json:"ttl"`
}

type SessionState struct {
	Accounts []string `json:"accounts"`
}

type SessionResponse struct {
	State SessionState `json:"state"`
}

type Reason struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RequestArguments struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type JsonRpcRequest struct {
	Id      int         `json:"id"`
	Jsonrpc string      `json:"jsonrpc"` //"2.0"
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type JsonRpcResult struct {
	Id      int         `json:"id"`
	Jsonrpc string      `json:"jsonrpc"` //"2.0"
	Result  interface{} `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type JsonRpcResponse = JsonRpcResult

type Notification struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type SessionRequest struct {
	Topic   string         `json:"topic"`
	Request JsonRpcRequest `json:"request"`
	ChainId string         `json:"chain_id"`
}

type SessionNotification struct {
	Topic        string       `json:"topic"`
	Notification Notification `json:"notification"`
}

type PairingProposal struct {
	Topic string `json:"topic"`
	Relay struct {
		Protocol string      `json:"protocol"`
		Params   interface{} `json:"params"`
	} `json:"relay"`
	Proposer struct {
		PublicKey string `json:"public_key"`
	} `json:"proposer"`
	Signal struct {
		Method string `json:"method"` //"pairing"
		Params struct {
			Uri string `json:"uri"`
		} `json:"params"`
	} `json:"signal"`
	Permissions struct {
		Jsonrpc struct {
			Methods []string `json:"methods"`
		} `json:"jsonrpc"`
	} `json:"permissions"`
	Ttl int `json:"ttl"`
}
