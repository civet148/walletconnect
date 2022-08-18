package wc

import (
	"github.com/civet148/log"
	"github.com/civet148/walletconnect/proto"
	"github.com/ybbus/jsonrpc"
)

type WCWallet struct {
	jsonrpc.RPCClient
	params *proto.ParamsInit
}

func NewWCWallet() *WCWallet {
	return &WCWallet{}
}

// initializes the client with persisted storage and a network connection
func (c *WCWallet) Init(params *proto.ParamsInit) bool {
	c.RPCClient = jsonrpc.NewClient(params.RelayProvider)
	if c.RPCClient == nil {
		log.Errorf("initiate connect failed")
		return false
	}
	c.params = params
	return true
}

// for proposer to propose a session to a responder
func (c *WCWallet) Connect(params *proto.ParamsConnect) {

}

// for responder to receive a session proposal from a proposer
func (c *WCWallet) Pair(params *proto.ParamsPair) {

}

// for responder to approve a session proposal
func (c *WCWallet) Approve(params *proto.ParamsApprove) {

}

// for responder to reject a session proposal
func (c *WCWallet) Reject(params *proto.ParamsReject) {

}

// for responder to upgrade session permissions
func (c *WCWallet) Upgrade(params *proto.ParamsUpgrade) {

}

// for responder to update session state
func (c *WCWallet) Update(params *proto.ParamsUpdate) {

}

// for proposer to request JSON-RPC
func (c *WCWallet) Request(params *proto.ParamsRequest) {

}

// for responder to respond JSON-RPC
func (c *WCWallet) Respond(params *proto.ParamsRespond) {

}

// for either to ping and verify peer is online
func (c *WCWallet) Ping(params *proto.ParamsPing) {

}

// for either to send notifications
func (c *WCWallet) Notify(params *proto.ParamsNotify) {

}

// for either to disconnect a session
func (c *WCWallet) Disconnect(params *proto.ParamsDisconnect) {

}
