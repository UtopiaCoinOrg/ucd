// Copyright (c) 2019 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpctest

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/UtopiaCoinOrg/ucd/blockchain"
	"github.com/UtopiaCoinOrg/ucd/blockchain/stake"
	"github.com/UtopiaCoinOrg/ucd/chaincfg"
	"github.com/UtopiaCoinOrg/ucd/chaincfg/chainhash"
	"github.com/UtopiaCoinOrg/ucd/ucec"
	"github.com/UtopiaCoinOrg/ucd/ucec/secp256k1"
	"github.com/UtopiaCoinOrg/ucd/ucjson/v2"
	"github.com/UtopiaCoinOrg/ucd/ucutil"
	"github.com/UtopiaCoinOrg/ucd/rpcclient/v3"
	"github.com/UtopiaCoinOrg/ucd/txscript"
	"github.com/UtopiaCoinOrg/ucd/wire"
)

var (
	// feeRate used when sending voting wallet transactions.
	feeRate = ucutil.Amount(1e4)

	// hardcodedPrivateKey used for all signing operations.
	hardcodedPrivateKey = []byte{
		0x79, 0xa6, 0x1a, 0xdb, 0xc6, 0xe5, 0xa2, 0xe1,
		0x39, 0xd2, 0x71, 0x3a, 0x54, 0x6e, 0xc7, 0xc8,
		0x75, 0x63, 0x2e, 0x75, 0xf1, 0xdf, 0x9c, 0x3f,
		0xa6, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	// nullPay2SSTXChange is the pkscript used on sstxchange outputs of the
	// tickets purchased by the voting wallet. This sends all change into a
	// null address, effectively discarding it.
	nullPay2SSTXChange = []byte{
		0xbd, 0xa9, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x87,
	}

	// stakebaseOutPoint is the outpoint that needs to be used in stakebase
	// inputs of vote transactions.
	stakebaseOutPoint = wire.OutPoint{Index: math.MaxUint32}

	// commitAmountMultiplier is a multiplier for the minimum stake difficulty,
	// used to fund inputs used in purchasing tickets. This needs to be high
	// enough that (minimumStakeDifficulty*commitAmountMultiplier) -
	// minimumStakeDifficulty is grater than the dust limit and will allow the
	// ticket to be relayed on the network.
	commitAmountMultiplier = int64(4)
)

type blockConnectedNtfn struct {
	blockHeader  []byte
	transactions [][]byte
}

type winningTicketsNtfn struct {
	blockHash      *chainhash.Hash
	blockHeight    int64
	winningTickets []*chainhash.Hash
}

type ticketInfo struct {
	ticketPrice int64
}

type utxoInfo struct {
	outpoint wire.OutPoint
	amount   int64
}

// VotingWallet stores the state for a simulated voting wallet. Once it is
// started, it will receive notifications from the associated harness, purchase
// tickets and vote on blocks as necessary to keep the chain going.
//
// This currently only implements the bare minimum requirements for maintaining
// a functioning voting wallet and does not handle reorgs, multiple voting and
// ticket buying wallets, setting vote bits, expired/missed votes, etc.
//
// All operations (after initial funding) are done solely via stake
// transactions, so no additional regular transactions are published. This is
// ideal for use in test suites that require a large (greater than SVH) number
// of blocks.
type VotingWallet struct {
	hn         *Harness
	privateKey *secp256k1.PrivateKey
	address    ucutil.Address
	c          *rpcclient.Client

	blockConnectedNtfnChan chan blockConnectedNtfn
	winningTicketsNtfnChan chan winningTicketsNtfn
	quitChan               chan struct{}

	p2sstx           []byte
	commitmentScript []byte
	p2pkh            []byte
	voteScript       []byte
	voteReturnScript []byte

	errorReporter func(error)

	// miner is a function responsible for generating new blocks. If
	// specified, then this function is used instead of directly calling
	// the underlying harness' Generate().
	miner func(uint32) ([]*chainhash.Hash, error)

	subsidyCache *blockchain.SubsidyCache

	// utxos are the unspent outpoints not yet locked into a ticket.
	utxos []utxoInfo

	// tickets map the outstanding unspent tickets
	tickets map[chainhash.Hash]ticketInfo

	// maturingVotes tracks the votes maturing at each (future) block height,
	// which will be available for purchasing new tickets.
	maturingVotes map[int64][]utxoInfo
}

// NewVotingWallet creates a new minimal voting wallet for the given harness.
// This wallet should be able to maintain the chain generated by the miner node
// of the harness working after it has passed SVH (Stake Validation Height) by
// continuously buying tickets and voting on them.
func NewVotingWallet(hn *Harness) (*VotingWallet, error) {

	priv, pub := secp256k1.PrivKeyFromBytes(hardcodedPrivateKey)
	serPub := pub.SerializeCompressed()
	hashPub := ucutil.Hash160(serPub)
	addr, err := ucutil.NewAddressPubKeyHash(hashPub, hn.ActiveNet,
		ucec.STEcdsaSecp256k1)
	if err != nil {
		return nil, fmt.Errorf("unable to generate address for pubkey: %v", err)
	}

	p2sstx, err := txscript.PayToSStx(addr)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare p2sstx script: %v", err)
	}

	p2pkh, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare p2pkh script: %v", err)
	}

	commitAmount := ucutil.Amount(hn.ActiveNet.MinimumStakeDiff * commitAmountMultiplier)
	limit := uint16(0x0058)
	commitmentScript, err := txscript.GenerateSStxAddrPush(addr, commitAmount, limit)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare commitment script: %v", err)
	}

	voteScript, err := txscript.GenerateSSGenVotes(0x0001)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare vote script: %v", err)
	}

	voteReturnScript, err := txscript.PayToSSGen(addr)
	if err != nil {
		return nil, fmt.Errorf("unable to generate vote return script: %v", err)
	}

	// Hints for the initial sizing of the tickets and maturing votes maps.
	// Given we have a deterministic purchase process, this should allow us to
	// size these maps only once at setup time.
	hintTicketsCap := requiredTicketCount(hn.ActiveNet)
	hintMaturingVotesCap := int(hn.ActiveNet.CoinbaseMaturity)

	// Buffer length for notification channels. As long as we don't get
	// notifications faster than this, we should be fine.
	bufferLen := 20

	w := &VotingWallet{
		hn:                     hn,
		privateKey:             priv,
		address:                addr,
		p2sstx:                 p2sstx,
		p2pkh:                  p2pkh,
		commitmentScript:       commitmentScript,
		voteScript:             voteScript,
		voteReturnScript:       voteReturnScript,
		subsidyCache:           blockchain.NewSubsidyCache(0, hn.ActiveNet),
		tickets:                make(map[chainhash.Hash]ticketInfo, hintTicketsCap),
		maturingVotes:          make(map[int64][]utxoInfo, hintMaturingVotesCap),
		blockConnectedNtfnChan: make(chan blockConnectedNtfn, bufferLen),
		winningTicketsNtfnChan: make(chan winningTicketsNtfn, bufferLen),
		quitChan:               make(chan struct{}),
	}

	handlers := &rpcclient.NotificationHandlers{
		OnBlockConnected: w.onBlockConnected,
		OnWinningTickets: w.onWinningTickets,
	}

	rpcConf := hn.RPCConfig()
	for i := 0; i < 20; i++ {
		if w.c, err = rpcclient.New(&rpcConf, handlers); err != nil {
			time.Sleep(time.Duration(i) * 50 * time.Millisecond)
			continue
		}
		break
	}
	if w.c == nil {
		return nil, fmt.Errorf("unable to connect to miner node")
	}

	if err = w.c.NotifyBlocks(); err != nil {
		return nil, fmt.Errorf("unable to subscribe to block notifications: %v", err)
	}
	if err = w.c.NotifyWinningTickets(); err != nil {
		return nil, fmt.Errorf("unable to subscribe to winning tickets notification: %v", err)
	}

	return w, nil
}

// Start stars the goroutines necessary for this voting wallet to function.
func (w *VotingWallet) Start() error {
	value := w.hn.ActiveNet.MinimumStakeDiff * commitAmountMultiplier

	// Create enough outputs to perform the voting, each with twice the amount
	// of the minimum ticket price.
	//
	// The number of required outputs is twice the coinbase maturity, since
	// we buy TicketsPerBlock tickets per block, starting at SVH-TM. At SVH,
	// TicketsPerBlock tickets will mature and be selected to vote (given they
	// are the only ones in the live ticket pool).
	//
	// Every following block we purchase the same amount of tickets, such that
	// TicketsPerBlock are maturing.
	nbOutputs := requiredTicketCount(w.hn.ActiveNet)
	outputs := make([]*wire.TxOut, nbOutputs)

	for i := 0; i < nbOutputs; i++ {
		outputs[i] = wire.NewTxOut(value, w.p2pkh)
	}

	txid, err := w.hn.SendOutputs(outputs, feeRate)
	if err != nil {
		return fmt.Errorf("unable to fund voting wallet: %v", err)
	}

	// Build the outstanding utxos for ticket buying. These will be the first
	// nbOutputs outputs from txid (assuming the SendOutputs() from above always
	// sends the change last).
	utxos := make([]utxoInfo, nbOutputs)
	for i := 0; i < nbOutputs; i++ {
		utxos[i] = utxoInfo{
			outpoint: wire.OutPoint{Hash: *txid, Index: uint32(i), Tree: wire.TxTreeRegular},
			amount:   value,
		}
	}
	w.utxos = utxos

	go w.handleNotifications()

	return nil
}

// Stop signals all goroutines from this wallet to stop their functions.
func (w *VotingWallet) Stop() {
	close(w.quitChan)
}

// SetErrorReporting allows users of the voting wallet to specify a function
// that will be called whenever an error happens while purchasing tickets or
// generating votes.
func (w *VotingWallet) SetErrorReporting(f func(err error)) {
	w.errorReporter = f
}

// SetMiner allows users of the voting wallet to specify a function that will
// be used to mine new blocks instead of using the regular Generate function of
// the configured rpcclient.
//
// This allows callers to use a custom function to generate blocks, such as one
// that allows faster mining in simnet.
func (w *VotingWallet) SetMiner(f func(uint32) ([]*chainhash.Hash, error)) {
	w.miner = f
}

// GenerateBlocks generates blocks while ensuring the chain will continue past
// SVH indefinitely. This will generate a block then wait for the votes from
// this wallet to be sent and tickets to be purchased before either generating
// the next block or returning.
//
// This function will either return the hashes of the generated blocks or an
// error if, after generating a candidate block, votes and tickets aren't
// submitted in a timely fashion.
func (w *VotingWallet) GenerateBlocks(nb uint32) ([]*chainhash.Hash, error) {
	_, startHeight, err := w.c.GetBestBlock()
	if err != nil {
		return nil, err
	}

	nbVotes := int(w.hn.ActiveNet.TicketsPerBlock)
	hashes := make([]*chainhash.Hash, nb)

	miner := w.c.Generate
	if w.miner != nil {
		miner = w.miner
	}

	for i := uint32(0); i < nb; i++ {
		// genHeight is the height of the _next_ block (the one that will be
		// generated once we call generate()).
		genHeight := startHeight + int64(i) + 1

		h, err := miner(1)
		if err != nil {
			return nil, fmt.Errorf("unable to generate block at height %d: %v",
				genHeight, err)
		}
		hashes[i] = h[0]

		needsVotes := genHeight >= (w.hn.ActiveNet.StakeValidationHeight - 1)
		needsTickets := genHeight >= ticketPurchaseStartHeight(w.hn.ActiveNet)

		timeout := time.After(time.Second * 5)
		testTimeout := time.After(time.Millisecond * 2)
		gotAllReqs := !needsVotes && !needsTickets
		for !gotAllReqs {
			select {
			case <-timeout:
				mempoolTickets, _ := w.c.GetRawMempool(ucjson.GRMTickets)
				mempoolVotes, _ := w.c.GetRawMempool(ucjson.GRMVotes)
				var notGot []string
				if len(mempoolVotes) != nbVotes {
					notGot = append(notGot, "votes")
				}
				if len(mempoolTickets) != nbVotes {
					notGot = append(notGot, "tickets")
				}

				return nil, fmt.Errorf("timeout waiting for %s "+
					"at height %d", strings.Join(notGot, ","), genHeight)
			case <-w.quitChan:
				return nil, fmt.Errorf("wallet is stopping")
			case <-testTimeout:
				mempoolTickets, _ := w.c.GetRawMempool(ucjson.GRMTickets)
				mempoolVotes, _ := w.c.GetRawMempool(ucjson.GRMVotes)

				gotAllReqs = (!needsTickets || (len(mempoolTickets) >= nbVotes)) &&
					(!needsVotes || (len(mempoolVotes) >= nbVotes))
				testTimeout = time.After(time.Millisecond * 2)
			}
		}
	}

	return hashes, nil
}

func (w *VotingWallet) logError(err error) {
	if w.errorReporter != nil {
		w.errorReporter(err)
	}
}

func (w *VotingWallet) onBlockConnected(blockHeader []byte, transactions [][]byte) {
	w.blockConnectedNtfnChan <- blockConnectedNtfn{
		blockHeader:  blockHeader,
		transactions: transactions,
	}
}

func (w *VotingWallet) handleBlockConnectedNtfn(ntfn *blockConnectedNtfn) {
	var header wire.BlockHeader
	err := header.FromBytes(ntfn.blockHeader)
	if err != nil {
		w.logError(err)
		return
	}

	blockHeight := int64(header.Height)
	purchaseHeight := ticketPurchaseStartHeight(w.hn.ActiveNet)
	if blockHeight < purchaseHeight {
		// No need to purchase tickets yet.
		return
	}

	// Purchase TicketsPerBlock tickets.
	nbTickets := int(w.hn.ActiveNet.TicketsPerBlock)
	if len(w.utxos) < nbTickets {
		fmt.Println("errrr len utxos < nbTickets")
		w.logError(fmt.Errorf("number of available utxos (%d) less than "+
			"number of tickets to purchase (%d)", len(w.utxos), nbTickets))
		return
	}

	// Use a slightly higher ticket price than the current minimum, to allow us
	// to ignore stakediff changes at exactly the next block (where purchasing
	// at the current value would cause our tickets to be rejected).
	ticketPrice := header.SBits + (header.SBits / 6)
	commitAmount := w.hn.ActiveNet.MinimumStakeDiff * commitAmountMultiplier

	// Select utxos to use and mark them used.
	utxos := make([]utxoInfo, nbTickets)
	copy(utxos, w.utxos[len(w.utxos)-nbTickets:])
	w.utxos = w.utxos[:len(w.utxos)-nbTickets]

	tickets := make([]wire.MsgTx, nbTickets)
	for i := 0; i < nbTickets; i++ {
		changeAmount := utxos[i].amount - commitAmount

		t := &tickets[i]
		t.AddTxIn(wire.NewTxIn(&utxos[i].outpoint, wire.NullValueIn, nil))
		t.AddTxOut(wire.NewTxOut(ticketPrice, w.p2sstx))
		t.AddTxOut(wire.NewTxOut(0, w.commitmentScript))
		t.AddTxOut(wire.NewTxOut(changeAmount, nullPay2SSTXChange))

		prevScript := w.p2pkh
		if utxos[i].outpoint.Tree == wire.TxTreeStake {
			prevScript = w.voteReturnScript
		}

		sig, err := txscript.SignatureScript(t, 0, prevScript, txscript.SigHashAll,
			w.privateKey, true)
		if err != nil {
			w.logError(fmt.Errorf("failed to sign ticket tx: %v", err))
			return
		}
		t.TxIn[0].SignatureScript = sig
	}

	// Submit all tickets to the network.
	promises := make([]rpcclient.FutureSendRawTransactionResult, nbTickets)
	for i := 0; i < nbTickets; i++ {
		promises[i] = w.c.SendRawTransactionAsync(&tickets[i], true)
	}

	for i := 0; i < nbTickets; i++ {
		h, err := promises[i].Receive()
		if err != nil {
			w.logError(fmt.Errorf("unable to send ticket tx: %v", err))
			return
		}

		w.tickets[*h] = ticketInfo{
			ticketPrice: ticketPrice,
		}
	}

	// Mark all maturing votes (if any) as available for spending.
	if maturingVotes, has := w.maturingVotes[blockHeight]; has {
		w.utxos = append(w.utxos, maturingVotes...)
		delete(w.maturingVotes, blockHeight)
	}
}

func (w *VotingWallet) onWinningTickets(blockHash *chainhash.Hash, blockHeight int64,
	winningTickets []*chainhash.Hash) {

	w.winningTicketsNtfnChan <- winningTicketsNtfn{
		blockHash:      blockHash,
		blockHeight:    blockHeight,
		winningTickets: winningTickets,
	}
}

func (w *VotingWallet) handleWinningTicketsNtfn(ntfn *winningTicketsNtfn) {

	blockRefScript, err := txscript.GenerateSSGenBlockRef(*ntfn.blockHash,
		uint32(ntfn.blockHeight))
	if err != nil {
		w.logError(fmt.Errorf("unable to generate ssgen block ref: %v", err))
		return
	}

	voteScript := w.voteScript
	voteReturnScript := w.voteReturnScript
	stakebaseValue := blockchain.CalcStakeVoteSubsidy(
		w.subsidyCache, ntfn.blockHeight, w.hn.ActiveNet,
	)

	// Create the votes. nbVotes is the number of tickets from the wallet that
	// voted.
	votes := make([]wire.MsgTx, w.hn.ActiveNet.TicketsPerBlock)
	nbVotes := 0

	var (
		ticket   ticketInfo
		myTicket bool
	)

	for _, wt := range ntfn.winningTickets {
		if ticket, myTicket = w.tickets[*wt]; !myTicket {
			continue
		}

		voteReturnValue := ticket.ticketPrice + stakebaseValue

		// Create a corresponding vote transaction.
		vote := &votes[nbVotes]
		nbVotes++
		vote.AddTxIn(wire.NewTxIn(
			&stakebaseOutPoint, stakebaseValue, w.hn.ActiveNet.StakeBaseSigScript,
		))
		vote.AddTxIn(wire.NewTxIn(
			wire.NewOutPoint(wt, 0, wire.TxTreeStake),
			wire.NullValueIn, nil,
		))
		vote.AddTxOut(wire.NewTxOut(0, blockRefScript))
		vote.AddTxOut(wire.NewTxOut(0, voteScript))
		vote.AddTxOut(wire.NewTxOut(voteReturnValue, voteReturnScript))

		sig, err := txscript.SignatureScript(vote, 1, w.p2sstx, txscript.SigHashAll,
			w.privateKey, true)
		if err != nil {
			w.logError(fmt.Errorf("failed to sign ticket tx: %v", err))
			return
		}
		vote.TxIn[1].SignatureScript = sig

		err = stake.CheckSSGen(vote)
		if err != nil {
			w.logError(fmt.Errorf("transaction is not a valid vote: %v", err))
			return
		}
	}

	newUtxos := make([]utxoInfo, nbVotes)

	// Publish the votes.
	promises := make([]rpcclient.FutureSendRawTransactionResult, nbVotes)
	for i := 0; i < nbVotes; i++ {
		promises[i] = w.c.SendRawTransactionAsync(&votes[i], true)
	}
	for i := 0; i < nbVotes; i++ {
		h, err := promises[i].Receive()
		if err != nil {
			w.logError(fmt.Errorf("unable to send vote tx: %v", err))
			return
		}
		newUtxos[i] = utxoInfo{
			outpoint: wire.OutPoint{Hash: *h, Index: 2, Tree: wire.TxTreeStake},
			amount:   votes[i].TxOut[2].Value,
		}
	}

	maturingHeight := ntfn.blockHeight + int64(w.hn.ActiveNet.CoinbaseMaturity)
	w.maturingVotes[maturingHeight] = newUtxos
}

// handleNotifications handles all notifications. This blocks until quitChan
// is closed and MUST be run on a separate goroutine.
func (w *VotingWallet) handleNotifications() {
	for {
		select {
		case <-w.quitChan:
			return
		case ntfn := <-w.blockConnectedNtfnChan:
			w.handleBlockConnectedNtfn(&ntfn)
		case ntfn := <-w.winningTicketsNtfnChan:
			w.handleWinningTicketsNtfn(&ntfn)
		}
	}
}

// ticketPurchaseStartHeight returns the block height where ticket buying
// needs to start so that there will be enough mature tickets for voting
// once SVH is reached.
func ticketPurchaseStartHeight(net *chaincfg.Params) int64 {
	return net.StakeValidationHeight - int64(net.TicketMaturity) - 2
}

// requiredTicketCount returns the number of tickets required to maintain the
// network functioning past SVH, assuming only as many tickets as votes will
// be purchased at every block.
func requiredTicketCount(net *chaincfg.Params) int {
	return int((net.CoinbaseMaturity + net.TicketMaturity + 2) * net.TicketsPerBlock)
}
