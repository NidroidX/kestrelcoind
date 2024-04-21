package rpchandlers

import (
	"github.com/NidroidX/kestrelcoind/app/appmessage"
	"github.com/NidroidX/kestrelcoind/app/rpc/rpccontext"
	"github.com/NidroidX/kestrelcoind/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
