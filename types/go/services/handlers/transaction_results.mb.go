// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsHandler

type TransactionResultsHandler interface {
	HandleTransactionResults(input *HandleTransactionResultsInput) (*HandleTransactionResultsOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsInput (non serializable)

type HandleTransactionResultsInput struct {
	TransactionReceipt []*protocol.TransactionReceipt
	BlockHeight uint64
	Timestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsOutput (non serializable)

type HandleTransactionResultsOutput struct {
}

