package protowire

import (
	"github.com/NidroidX/kestrelcoind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *kestrelcoindMessage_UnbanRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "kestrelcoindMessage_UnbanRequest is nil")
	}
	return x.UnbanRequest.toAppMessage()
}

func (x *kestrelcoindMessage_UnbanRequest) fromAppMessage(message *appmessage.UnbanRequestMessage) error {
	x.UnbanRequest = &UnbanRequestMessage{Ip: message.IP}
	return nil
}

func (x *UnbanRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "UnbanRequestMessage is nil")
	}
	return &appmessage.UnbanRequestMessage{
		IP: x.Ip,
	}, nil
}

func (x *kestrelcoindMessage_UnbanResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "kestrelcoindMessage_UnbanResponse is nil")
	}
	return x.UnbanResponse.toAppMessage()
}

func (x *kestrelcoindMessage_UnbanResponse) fromAppMessage(message *appmessage.UnbanResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.UnbanResponse = &UnbanResponseMessage{
		Error: err,
	}
	return nil
}

func (x *UnbanResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "UnbanResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.UnbanResponseMessage{
		Error: rpcErr,
	}, nil
}
