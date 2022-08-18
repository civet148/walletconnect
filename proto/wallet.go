package proto

type ParamsInit struct {
	Controller    bool        `json:"controller"`
	Metadata      AppMetadata `json:"metadata"`
	RelayProvider string      `json:"relay_provider"`
}

type ParamsConnect struct {
	Permissions SessionPermissions `json:"permissions"`
	Pairing     Sequence           `json:"pairing"`
}

type ParamsPair struct {
	Uri string `json:"uri"`
}

type ParamsApprove struct {
	Proposal SessionProposal `json:"proposal"`
	Response SessionResponse `json:"response"`
}

type ParamsReject struct {
	Proposal SessionProposal `json:"proposal"`
	Reason   Reason          `json:"reason"`
}

type ParamsUpgrade struct {
	Topic       string             `json:"topic"`
	Permissions SessionPermissions `json:"permissions"`
}

type ParamsUpdate struct {
	Topic string       `json:"topic"`
	State SessionState `json:"state"`
}

type ParamsRequest struct {
	Topic   string           `json:"topic"`
	Request RequestArguments `json:"request"`
	ChainId string           `json:"chain_id"`
}

type ParamsRespond struct {
	Topic    string          `json:"topic"`
	Response JsonRpcResponse `json:"response"`
}

type ParamsPing struct {
	Topic string `json:"topic"`
}

type ParamsNotify struct {
	Topic        string       `json:"topic"`
	Notification Notification `json:"notification"`
}

type ParamsDisconnect struct {
	Topic  string `json:"topic"`
	Reason Reason `json:"reason"`
}
