package protowire

import (
	"github.com/NidroidX/kestrelcoind/app/appmessage"
	"github.com/pkg/errors"
)

func (x *kestrelcoindMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "kestrelcoindMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *kestrelcoindMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
