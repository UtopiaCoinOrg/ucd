// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2019 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"math/big"
	"time"

	"github.com/UtopiaCoinOrg/ucd/chaincfg/chainhash"
	"github.com/UtopiaCoinOrg/ucd/wire"
)

// TestNet3Params return the network parameters for the test currency network.
// This network is sometimes simply called "testnet".
// This is the third public iteration of testnet.
func TestNet3Params() *Params {
	// testNetPowLimit is the highest proof of work value a Utopia block
	// can have for the test network.  It is the value 2^232 - 1.
	testNetPowLimit := new(big.Int).Sub(new(big.Int).Lsh(bigOne, 236), bigOne)

	// genesisBlock defines the genesis block of the block chain which serves as
	// the public transaction ledger for the test network (version 3).
	genesisBlock := wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:   6,
			PrevBlock: chainhash.Hash{},
			// MerkleRoot: Calculated below.
			Timestamp:    time.Unix(1546819200, 0), // 2018-08-06 00:00:00 +0000 UTC
			Bits:         0x1e0fffff,               // Difficulty 1 [00000fffff000000000000000000000000000000000000000000000000000000]
			SBits:        20000000,
			Nonce:        0x18aea41a,
			StakeVersion: 6,
		},
		Transactions: []*wire.MsgTx{{
			SerType: wire.TxSerializeFull,
			Version: 1,
			TxIn: []*wire.TxIn{{
				// Fully null.
				PreviousOutPoint: wire.OutPoint{
					Hash:  chainhash.Hash{},
					Index: 0xffffffff,
					Tree:  0,
				},
				SignatureScript: hexDecode("0000"),
				Sequence:        0xffffffff,
				BlockHeight:     wire.NullBlockHeight,
				BlockIndex:      wire.NullBlockIndex,
				ValueIn:         wire.NullValueIn,
			}},
			TxOut: []*wire.TxOut{{
				Version: 0x0000,
				Value:   0x00000000,
				PkScript: hexDecode("801679e98561ada96caec2949a5d41c4cab3851e" +
					"b740d951c10ecbcf265c1fd9"),
			}},
			LockTime: 0,
			Expiry:   0,
		}},
	}
	// NOTE: This really should be TxHashFull, but it was defined incorrectly.
	//
	// Since the field is not used in any validation code, it does not have any
	// adverse effects, but correcting it would result in changing the block
	// hash which would invalidate the entire test network.  The next test
	// network should set the value properly.
	genesisBlock.Header.MerkleRoot = genesisBlock.Transactions[0].TxHash()

	return &Params{
		Name:        "testnet3",
		Net:         wire.TestNet3,
		DefaultPort: "11508",
		DNSSeeds: []DNSSeed{
			{"testnetseed.utopiacoin.org", true},
			{"testnetseed1.utopiacoin.org", true},
			{"testnetseed2.utopiacoin.org", true},
			{"testnetseed3.utopiacoin.org", true},
			{"testnetseed4.utopiacoin.org", true},
			{"testnetseed5.utopiacoin.org", true},
			{"testnetseed.utopiacoin.network", true},
			{"testnetseed.utopiacoin.io", true},
			{"testnetseed.utopiacoin.org", true},
		},

		// Chain parameters
		GenesisBlock:             &genesisBlock,
		GenesisHash:              genesisBlock.BlockHash(),
		PowLimit:                 testNetPowLimit,
		PowLimitBits:             0x1e0fffff,
		ReduceMinDifficulty:      false,
		MinDiffReductionTime:     0, // ~99.3% chance to be mined before reduction
		GenerateSupported:        true,
		MaximumBlockSizes:        []int{1310720},
		MaxTxSize:                1000000,
		TargetTimePerBlock:       time.Second * 40,
		WorkDiffAlpha:            1,
		WorkDiffWindowSize:       144,
		WorkDiffWindows:          20,
		TargetTimespan:           time.Second * 40 * 144, // TimePerBlock * WindowSize
		RetargetAdjustmentFactor: 4,

		// Subsidy parameters.
		BaseSubsidy:              250000000000, //2500 coins
		MulSubsidy:               9943,
		DivSubsidy:               10000,
		SubsidyReductionInterval: 20480,
		WorkRewardProportion:     5,
		StakeRewardProportion:    4,
		BlockTaxProportion:       1,


		//flash tx
		FlashSendConfirmationsRequired:6,
		// Checkpoints ordered from oldest to newest.
		Checkpoints: []Checkpoint{
		//	{83520, newHashFromStr("0000000001e6244d95feae8b598e854905158c7bc781daf874afff88675ef0c8")},
		},

		// Consensus rule change deployments.
		//
		// The miner confirmation window is defined as:
		//   target proof of work timespan / target proof of work spacing
		RuleChangeActivationQuorum:     2520, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
		RuleChangeActivationMultiplier: 3,    // 75%
		RuleChangeActivationDivisor:    4,
		RuleChangeActivationInterval:   5040, // 1 week
		Deployments: map[uint32][]ConsensusDeployment{
		/*	7: {{
				Vote: Vote{
					Id:          VoteIDFixLNSeqLocks,
					Description: "Modify sequence lock handling as defined in DCP0004",
					Mask:        0x0006, // Bits 1 and 2
					Choices: []Choice{{
						Id:          "abstain",
						Description: "abstain voting for change",
						Bits:        0x0000,
						IsAbstain:   true,
						IsNo:        false,
					}, {
						Id:          "no",
						Description: "keep the existing consensus rules",
						Bits:        0x0002, // Bit 1
						IsAbstain:   false,
						IsNo:        true,
					}, {
						Id:          "yes",
						Description: "change to the new consensus rules",
						Bits:        0x0004, // Bit 2
						IsAbstain:   false,
						IsNo:        false,
					}},
				},
				StartTime:  1548633600, // Jan 28th, 2019
				ExpireTime: 1580169600, // Jan 28th, 2020
			}},
			*/
		},

		// Enforce current block version once majority of the network has
		// upgraded.
		// 51% (51 / 100)
		// Reject previous block versions once a majority of the network has
		// upgraded.
		// 75% (75 / 100)
		BlockEnforceNumRequired: 51,
		BlockRejectNumRequired:  75,
		BlockUpgradeNumToCheck:  100,

		// AcceptNonStdTxs is a mempool param to either accept and relay non
		// standard txs to the network or reject them
		AcceptNonStdTxs: true,

		// Address encoding magics
		NetworkAddressPrefix: "T",
		PubKeyAddrID:         [2]byte{0x28, 0xf7}, // starts with Tk
		PubKeyHashAddrID:     [2]byte{0x0e, 0xc0}, // starts with TC
		PKHEdwardsAddrID:     [2]byte{0x0f, 0x01}, // starts with Te
		PKHSchnorrAddrID:     [2]byte{0x0f, 0x20}, // starts with Ts
		ScriptHashAddrID:     [2]byte{0x0f, 0x12}, // starts with Tm
		PrivateKeyID:         [2]byte{0x23, 0x0e}, // starts with Pt

		// BIP32 hierarchical deterministic extended key magics
		HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
		HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

		// BIP44 coin type used in the hierarchical deterministic path for
		// address generation.
		SLIP0044CoinType: 401,  // SLIP0044, Testnet (all coins)
		LegacyCoinType:   11, // for backwards compatibility

		// Utopia PoS parameters
		MinimumStakeDiff:        50000 * 1e8, // 50000 Coin
		TicketPoolSize:          820,
		TicketsPerBlock:         5,
		TicketMaturity:          16,
		TicketExpiry:            4100, // 5*TicketPoolSize
		CoinbaseMaturity:        16,
		SStxChangeMaturity:      1,
		TicketPoolSizeWeight:    4,
		StakeDiffAlpha:          1,
		StakeDiffWindowSize:     144,
		StakeDiffWindows:        20,
		StakeVersionInterval:    144 * 2 * 7, // ~1 week
		MaxFreshStakePerBlock:   20,          // 4*TicketsPerBlock
		StakeEnabledHeight:      16 + 16,     // CoinbaseMaturity + TicketMaturity
		StakeValidationHeight:   768,         // Arbitrary
		StakeBaseSigScript:      []byte{0x00, 0x00},
		StakeMajorityMultiplier: 3,
		StakeMajorityDivisor:    4,

		// Utopia organization related parameters.
		// Organization address is TCaowBdDEu9877qkATPju4Ug1L4ASnnC7kn.
		OrganizationPkScript:        hexDecode("76a9147ef6513c048101bf44823ad391e2f1eb0b6b0d3988ac"),
		OrganizationPkScriptVersion: 0,
		BlockOneLedger:              BlockOneLedgerTestNet3,
	}
}
