// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

type MsgFlashTx struct {
	MsgTx
}

func NewMsgFlashTx() *MsgFlashTx {
	return &MsgFlashTx{
		MsgTx: *NewMsgTx(),
	}
}

func NewMsgFlashTxFromMsgTx(msgTx *MsgTx) *MsgFlashTx {
	return &MsgFlashTx{
		MsgTx: *msgTx,
	}
}


func (msg *MsgFlashTx) Command() string {
	return CmdFlashTx
}

