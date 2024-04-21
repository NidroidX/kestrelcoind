package testapi

import (
	"github.com/NidroidX/kestrelcoind/domain/consensus/model"
	"github.com/NidroidX/kestrelcoind/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
