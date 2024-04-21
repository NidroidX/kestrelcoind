package prefixmanager

import (
	"github.com/NidroidX/kestrelcoind/infrastructure/logger"
	"github.com/NidroidX/kestrelcoind/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
