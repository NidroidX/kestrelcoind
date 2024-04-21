package rpchandlers

import (
	"github.com/NidroidX/kestrelcoind/app/appmessage"
	"github.com/NidroidX/kestrelcoind/app/rpc/rpccontext"
	"github.com/NidroidX/kestrelcoind/domain/consensus/utils/constants"
	"github.com/NidroidX/kestrelcoind/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kestrelcoind is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSiumSupply, err := context.UTXOIndex.GetCirculatingSiumSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSium,
		circulatingSiumSupply,
	)

	return response, nil
}
