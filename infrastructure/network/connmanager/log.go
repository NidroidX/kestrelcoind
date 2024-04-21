package connmanager

import (
	"github.com/NidroidX/kestrelcoind/infrastructure/logger"
	"github.com/NidroidX/kestrelcoind/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
