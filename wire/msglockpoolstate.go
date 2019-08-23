// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers 
// Copyright (c) 2018-2020 The Hc developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"github.com/UtopiaCoinOrg/ucd/chaincfg/chainhash"
	"io"
)

// MaxMSBlocksAtHeadPerMsg is the maximum number of block hashes allowed
// per message.
const MaxFlashTx = 20
const MaxFlashTxVote = 100

// MsgMiningState implements the Message interface and represents a mining state
// message.  It is used to request a list of blocks located at the chain tip
// along with all votes for those blocks.  The list is returned is limited by
// the maximum number of blocks per message and the maximum number of votes per
// message.
type MsgLockPoolState struct {
	FlashTxHashes     []*chainhash.Hash
	FlashTxVoteHashes []*chainhash.Hash
}

// AddBlockHash adds a new block hash to the message.
func (msg *MsgLockPoolState) AddFlashTxHash(hash *chainhash.Hash) error {
	if len(msg.FlashTxHashes)+1 > MaxFlashTx {
		str := fmt.Sprintf("too many flashtx hashes for message [max %v]",
			MaxFlashTx)
		return messageError("MsgLockPoolState.AddBlockHash", str)
	}

	msg.FlashTxHashes = append(msg.FlashTxHashes, hash)
	return nil
}

// AddVoteHash adds a new vote hash to the message.
func (msg *MsgLockPoolState) AddFlashTxVoteHash(hash *chainhash.Hash) error {
	if len(msg.FlashTxVoteHashes)+1 > MaxFlashTxVote {
		str := fmt.Sprintf("too many vote hashes for message [max %v]",
			MaxFlashTxVote)
		return messageError("MsgLockPoolState.AddVoteHash", str)
	}

	msg.FlashTxVoteHashes = append(msg.FlashTxVoteHashes, hash)
	return nil
}

// BtcDecode decodes r using the protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgLockPoolState) BtcDecode(r io.Reader, pver uint32) error {
	// Read num block hashes and limit to max.
	count, err := ReadVarInt(r, pver)
	if err != nil {
		return err
	}
	if count > MaxFlashTx {
		str := fmt.Sprintf("too many flashTx hashes for message "+
			"[count %v, max %v]", count, MaxFlashTx)
		return messageError("MsgLockPoolState.BtcDecode", str)
	}

	msg.FlashTxHashes = make([]*chainhash.Hash, 0, count)
	for i := uint64(0); i < count; i++ {
		hash := chainhash.Hash{}
		err := readElement(r, &hash)
		if err != nil {
			return err
		}
		msg.AddFlashTxHash(&hash)
	}

	// Read num vote hashes and limit to max.
	count, err = ReadVarInt(r, pver)
	if err != nil {
		return err
	}
	if count > MaxFlashTxVote {
		str := fmt.Sprintf("too many vote hashes for message "+
			"[count %v, max %v]", count, MaxFlashTxVote)
		return messageError("MsgLockPoolState.BtcDecode", str)
	}

	msg.FlashTxVoteHashes = make([]*chainhash.Hash, 0, count)
	for i := uint64(0); i < count; i++ {
		hash := chainhash.Hash{}
		err := readElement(r, &hash)
		if err != nil {
			return err
		}
		err = msg.AddFlashTxVoteHash(&hash)
		if err != nil {
			return err
		}
	}

	return nil
}

// BtcEncode encodes the receiver to w using the protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgLockPoolState) BtcEncode(w io.Writer, pver uint32) error {
	// Write block hashes.
	count := len(msg.FlashTxHashes)
	if count > MaxFlashTx {
		str := fmt.Sprintf("too many flashTx hashes for message "+
			"[count %v, max %v]", count, MaxFlashTx)
		return messageError("MsgLockPoolState.BtcEncode", str)
	}

	err := WriteVarInt(w, pver, uint64(count))
	if err != nil {
		return err
	}

	for _, hash := range msg.FlashTxHashes {
		err = writeElement(w, hash)
		if err != nil {
			return err
		}
	}

	// Write vote hashes.
	count = len(msg.FlashTxVoteHashes)
	if count > MaxFlashTxVote {
		str := fmt.Sprintf("too many vote hashes for message "+
			"[count %v, max %v]", count, MaxFlashTxVote)
		return messageError("MsgLockPoolState.BtcEncode", str)
	}

	err = WriteVarInt(w, pver, uint64(count))
	if err != nil {
		return err
	}

	for _, hash := range msg.FlashTxVoteHashes {
		err = writeElement(w, hash)
		if err != nil {
			return err
		}
	}

	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgLockPoolState) Command() string {
	return CmdLockPoolState
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgLockPoolState) MaxPayloadLength(pver uint32) uint32 {
	//  + num block hashes (varInt) +
	// block hashes + num vote hashes (varInt) + vote hashes
	return MaxVarIntPayload + (MaxFlashTx *
		chainhash.HashSize) + MaxVarIntPayload + (MaxFlashTxVote *
		chainhash.HashSize)
}

// NewMsgMiningState returns a new hcd MsgLockPoolState message that conforms to
// the Message interface using the defaults for the fields.
func NewMsgLockPoolState() *MsgLockPoolState {
	return &MsgLockPoolState{
		FlashTxHashes:     make([]*chainhash.Hash, 0, MaxFlashTx),
		FlashTxVoteHashes: make([]*chainhash.Hash, 0, MaxFlashTxVote),
	}
}
