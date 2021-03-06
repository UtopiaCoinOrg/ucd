// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2015-2016 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ucutil

const (
	// AtomsPerCent is the number of atomic units in one coin cent.
	AtomsPerCent = 1e6

	// AtomsPerCoin is the number of atomic units in one coin.
	AtomsPerCoin = 1e8

	// MaxAmount is the maximum transaction amount allowed in atoms.
	// Utopia - Changeme for release
	MaxAmount = 100e8 * AtomsPerCoin
)
