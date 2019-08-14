// Copyright (c) 2015-2019 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ucutil_test

import (
	"fmt"
	"math"

	"github.com/UtopiaCoinOrg/ucd/chaincfg"
	"github.com/UtopiaCoinOrg/ucd/ucec"
	"github.com/UtopiaCoinOrg/ucd/ucutil"
)

func ExampleAmount() {

	a := ucutil.Amount(0)
	fmt.Println("Zero Atom:", a)

	a = ucutil.Amount(1e8)
	fmt.Println("100,000,000 Atoms:", a)

	a = ucutil.Amount(1e5)
	fmt.Println("100,000 Atoms:", a)
	// Output:
	// Zero Atom: 0 UC
	// 100,000,000 Atoms: 1 UC
	// 100,000 Atoms: 0.001 UC
}

func ExampleNewAmount() {
	amountOne, err := ucutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := ucutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := ucutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := ucutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 UC
	// 0.01234567 UC
	// 0 UC
	// invalid coin amount
}

func ExampleAmount_unitConversions() {
	amount := ucutil.Amount(44433322211100)

	fmt.Println("Atom to kCoin:", amount.Format(ucutil.AmountKiloCoin))
	fmt.Println("Atom to Coin:", amount)
	fmt.Println("Atom to MilliCoin:", amount.Format(ucutil.AmountMilliCoin))
	fmt.Println("Atom to MicroCoin:", amount.Format(ucutil.AmountMicroCoin))
	fmt.Println("Atom to Atom:", amount.Format(ucutil.AmountAtom))

	// Output:
	// Atom to kCoin: 444.333222111 kUC
	// Atom to Coin: 444333.222111 UC
	// Atom to MilliCoin: 444333222.111 mUC
	// Atom to MicroCoin: 444333222111 μUC
	// Atom to Atom: 44433322211100 Atom
}

// This example demonstrates decoding addresses, determining their underlying
// type, and displaying their associated underlying hash160 and digitial
// signature algorithm.
func ExampleDecodeAddress() {
	// Ordinarily addresses would be read from the user or the result of a
	// derivation, but they are hard coded here for the purposes of this
	// example.
	mainNetParmas := chaincfg.MainNetParams()
	addrsToDecode := []string{
		"DsRUvfCwTMrKz29dDiQBJhZii9GDN3bVx6Q", // pay-to-pubkey-hash ecdsa
		"DSpf9Sru9MarMKQQnuzTiQ9tjWVJA3KSm2d", // pay-to-pubkey-hash schnorr
	}
	for idx, encodedAddr := range addrsToDecode {
		addr, err := ucutil.DecodeAddress(encodedAddr, mainNetParmas)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("addr%d hash160: %x\n", idx, *addr.Hash160())

		// The example addresses are pay-to-pubkey-hash with different signature
		// algorithms, so this code is limited to that type
		switch a := addr.(type) {
		case *ucutil.AddressPubKeyHash:
			// Determine and display the digitial signature algorithm.
			algo := "unknown"
			switch a.DSA() {
			case ucec.STEcdsaSecp256k1:
				algo = "ECDSA"
			case ucec.STSchnorrSecp256k1:
				algo = "Schnorr"
			}
			fmt.Printf("addr%d DSA: %v\n", idx, algo)

		default:
			fmt.Println("Unexpected test address type")
			return
		}
	}

	// Output:
	// addr0 hash160: 05ad744deacf5334671d3e62db86230af1891f71
	// addr0 DSA: ECDSA
	// addr1 hash160: e280cb6e66b96679aec288b1fbdbd4db08077a1b
	// addr1 DSA: Schnorr
}
