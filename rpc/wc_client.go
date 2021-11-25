package rpc

import (
	"github.com/civet148/log"
	"github.com/civet148/walletconnect/proto"
	"github.com/ybbus/jsonrpc"
)

type WCClient struct {
	jsonrpc.RPCClient
	params *proto.ParamsInit
}

func NewWCClient() *WCClient {
	return &WCClient{}
}

// initializes the client with persisted storage and a network connection
func (c *WCClient) Init(params *proto.ParamsInit) bool {
	c.RPCClient = jsonrpc.NewClient(params.RelayProvider)
	if c.RPCClient == nil {
		log.Errorf("initiate connect failed")
		return false
	}
	c.params = params
	return true
}

// for proposer to propose a session to a responder
func (c *WCClient) Connect(params *proto.ParamsConnect) {

}

// for responder to receive a session proposal from a proposer
func (c *WCClient) Pair(params *proto.ParamsPair) {

}

// for responder to approve a session proposal
func (c *WCClient) Approve(params *proto.ParamsApprove) {

}

// for responder to reject a session proposal
func (c *WCClient) Reject(params *proto.ParamsReject) {

}

// for responder to upgrade session permissions
func (c *WCClient) Upgrade(params *proto.ParamsUpgrade) {

}

// for responder to update session state
func (c *WCClient) Update(params *proto.ParamsUpdate) {

}

// for proposer to request JSON-RPC
func (c *WCClient) Request(params *proto.ParamsRequest) {

}

// for responder to respond JSON-RPC
func (c *WCClient) Respond(params *proto.ParamsRespond) {

}

// for either to ping and verify peer is online
func (c *WCClient) Ping(params *proto.ParamsPing) {

}

// for either to send notifications
func (c *WCClient) Notify(params *proto.ParamsNotify) {

}

// for either to disconnect a session
func (c *WCClient) Disconnect(params *proto.ParamsDisconnect) {

}
