package ucutil

import (
	"bytes"
	"github.com/UtopiaCoinOrg/ucd/chaincfg/chainec"
	"github.com/UtopiaCoinOrg/ucd/chaincfg/chainhash"
	"github.com/UtopiaCoinOrg/ucd/wire"
	"github.com/UtopiaCoinOrg/ucwallet/errors"
)



type FlashTxVote struct {
	msgFlashTxVote *wire.MsgFlashTxVote
}

func NewFlashTxVote(vote *wire.MsgFlashTxVote) *FlashTxVote {
	return &FlashTxVote{
		msgFlashTxVote: vote,
	}
}

func (flashTxVote *FlashTxVote) Hash() *chainhash.Hash {
	return flashTxVote.msgFlashTxVote.Hash()
}

func (flashTxVote *FlashTxVote) MsgFlashTxVote() *wire.MsgFlashTxVote {
	return flashTxVote.msgFlashTxVote
}

func (flashTxVote *FlashTxVote) GetPubKey() []byte{
	return flashTxVote.msgFlashTxVote.PubKey
}

type FlashTx struct {
	Tx
}

// MsgTx returns the underlying wire.MsgTx for the transaction.
func (t *FlashTx) MsgFlashTx() *wire.MsgFlashTx {
	// Return the cached transaction.
	return wire.NewMsgFlashTxFromMsgTx(t.msgTx)
}

func NewFlashTx(msgFlashTx *wire.MsgFlashTx) *FlashTx {
	return &FlashTx{
		Tx: Tx{
			hash:    msgFlashTx.TxHash(),
			msgTx:   &msgFlashTx.MsgTx,
			txTree:  wire.TxTreeUnknown,
			txIndex: TxIndexUnknown,
		},
	}
}

func NewFlashTxFromTx(tx *Tx) *FlashTx {
	return &FlashTx{
		Tx: *tx,
	}
}





// VerifyMessage verifies that sig is a valid signature of msg and was created
// using the secp256k1 private key for addr.
func VerifyMessage(msg string, addr Address, sig []byte) (bool, error) {
	const op errors.Op = "wallet.VerifyMessage"
	// Validate the signature - this just shows that it was valid for any pubkey
	// at all. Whether the pubkey matches is checked below.
	var buf bytes.Buffer
	wire.WriteVarString(&buf, 0, "Utopia Signed Message:\n")
	wire.WriteVarString(&buf, 0, msg)
	expectedMessageHash := chainhash.HashB(buf.Bytes())
	pk, wasCompressed, err := chainec.Secp256k1.RecoverCompact(sig,
		expectedMessageHash)
	if err != nil {
		return false, errors.E(op, err)
	}

	// Reconstruct the address from the recovered pubkey.
	var serializedPK []byte
	if wasCompressed {
		serializedPK = pk.SerializeCompressed()
	} else {
		serializedPK = pk.SerializeUncompressed()
	}
	recoveredAddr, err := NewAddressSecpPubKey(serializedPK, ActiveNet)
	if err != nil {
		return false, errors.E(op, err)
	}

	// Return whether addresses match.
	return recoveredAddr.Address() == addr.Address(), nil
}