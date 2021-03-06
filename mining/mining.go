// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2016 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package mining

import (
	"time"

	"github.com/UtopiaCoinOrg/ucd/blockchain/stake"
	"github.com/UtopiaCoinOrg/ucd/chaincfg/chainhash"
	"github.com/UtopiaCoinOrg/ucd/ucutil"
)

const (
	// MinHighPriority is the minimum priority value that allows a
	// transaction to be considered high priority.
	MinHighPriority = ucutil.AtomsPerCoin * 144.0 / 250
)

// TxDesc is a descriptor about a transaction in a transaction source along with
// additional metadata.
type TxDesc struct {
	// Tx is the transaction associated with the entry.
	Tx *ucutil.Tx

	// Type is the type of the transaction associated with the entry.
	Type stake.TxType

	// Added is the time when the entry was added to the source pool.
	Added time.Time

	// Height is the block height when the entry was added to the the source
	// pool.
	Height int64

	// Fee is the total fee the transaction associated with the entry pays.
	Fee int64
}

// VoteDesc is a descriptor about a vote transaction in a transaction source
// along with additional metadata.
type VoteDesc struct {
	VoteHash       chainhash.Hash
	TicketHash     chainhash.Hash
	ApprovesParent bool
}

// TxSource represents a source of transactions to consider for inclusion in
// new blocks.
//
// The interface contract requires that all of these methods are safe for
// concurrent access with respect to the source.
type TxSource interface {
	// LastUpdated returns the last time a transaction was added to or
	// removed from the source pool.
	LastUpdated() time.Time

	// MiningDescs returns a slice of mining descriptors for all the
	// transactions in the source pool.
	MiningDescs() []*TxDesc

	// HaveTransaction returns whether or not the passed transaction hash
	// exists in the source pool.
	HaveTransaction(hash *chainhash.Hash) bool

	// HaveAllTransactions returns whether or not all of the passed
	// transaction hashes exist in the source pool.
	HaveAllTransactions(hashes []chainhash.Hash) bool

	// VoteHashesForBlock returns the hashes for all votes on the provided
	// block hash that are currently available in the source pool.
	VoteHashesForBlock(hash *chainhash.Hash) []chainhash.Hash

	// VotesForBlocks returns a slice of vote descriptors for all votes on
	// the provided block hashes that are currently available in the source
	// pool.
	VotesForBlocks(hashes []chainhash.Hash) [][]VoteDesc

	// IsRegTxTreeKnownDisapproved returns whether or not the regular
	// transaction tree of the block represented by the provided hash is
	// known to be disapproved according to the votes currently in the
	// source pool.
	IsRegTxTreeKnownDisapproved(hash *chainhash.Hash) bool
}
