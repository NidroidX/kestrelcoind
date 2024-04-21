package grpcclient

import (
	"github.com/NidroidX/kestrelcoind/infrastructure/logger"
	"github.com/NidroidX/kestrelcoind/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
