package consensus

import (
	"github.com/NidroidX/kestrelcoind/infrastructure/logger"
	"github.com/NidroidX/kestrelcoind/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
