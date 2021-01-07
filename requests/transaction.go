package requests

import (
	"github.com/uber/gonduit/entities"
)

// TransactionSearchRequest represents a request to
// transaction.search API method.
type TransactionSearchRequest struct {
	ObjectIdentifier string `json:"objectIdentifier"`
	// Constraints contains additional filters for results. Applied on top of
	// query if provided.
	Constraints *TransactionSearchConstraints `json:"constraints,omitempty"`

	*entities.Cursor
	Request
}

// TransactionSearchConstraints describes search criteria for request.
type TransactionSearchConstraints struct {
	PHIDs       []string `json:"phids,omitempty"`
	AuthorPHIDs []string `json:"authorPHIDs,omitempty"`
}

type Transactions struct {
	ObjectIdentifier string `json:"objectIdentifier"`
	Transactions     []Transaction
}

type Transaction struct {
	Type  string
	Value interface{}
}
