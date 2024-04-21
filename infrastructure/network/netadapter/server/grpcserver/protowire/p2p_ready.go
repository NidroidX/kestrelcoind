package protowire

import (
	"github.com/NidroidX/kestrelcoind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *kestrelcoindMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "kestrelcoindMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *kestrelcoindMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
