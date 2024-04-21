package protowire

import (
	"github.com/NidroidX/kestrelcoind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *kestrelcoindMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "kestrelcoindMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *kestrelcoindMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
