// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2018 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.

var BlockOneLedgerMainNet = []*TokenPayout{
	{"UCMiFJAa1aQ3r6RyTpZjWyL2RSdmRCgrNaa", 5000000 * 1e8},
	{"UCXa1L4rmMThrX1wvna45dpvAxSprVJosuC", 5000000 * 1e8},
	{"UCTrSQHeJXtwhzPXzUTBLXmaghcCvGYs24U", 5000000 * 1e8},
	{"UCWvLSvNSoviCsniynZ5dLqHD4c6uXP8D1J", 5000000 * 1e8},
	{"UCdq8HUMpmP8K5h5qqHBUGcYQNW4hSTy2DH", 5000000 * 1e8},
	{"UCPpXiCDhhBNhqN8rTLxg8MnD2Rm4q776mN", 5000000 * 1e8},
	{"UCc5SK87FAKHryfZ4Z7UimLZbHxLUmbvTxi", 5000000 * 1e8},
	{"UCSdJp6Rr2rYBdLNtEoqFPLW3oRYKWDtyZG", 5000000 * 1e8},
	{"UCdRCS1JzGb2yJvmYEgKbLemeNnvTzKfg7q", 500000 * 1e8},
	{"UCZJcB23f6G6ohEBhvtbYfcMvNqvc2QUwrr", 500000 * 1e8},
	{"UCe7PCoYHwbLPHqoKEAa57N2HcT3bQMdtNp", 500000 * 1e8},
	{"UCVEpxhbwSqoyaQusBGjMn39PYMrLToX9fN", 500000 * 1e8},
	{"UCVEpxhbwSqoyaQusBGjMn39PYMrLToX9fN", 958 * 1e14},
}

// BlockOneLedgerTestNet3 is the block one output ledger for testnet version 3.
var BlockOneLedgerTestNet3 = []*TokenPayout{
	{"TCUZB3H9UvrE6gN5TNfQP7gC996LABTYMT8", 2 * 1e16},
	{"TCY9ZaNxEvUCqChfUYg9nWAsyj7pZrZ2kQW", 2 * 1e16},
	{"TCbaskANHaMA5B2YybyyhsGR4cFjsK5evUa", 2 * 1e16},
	{"TCc2LFBRLyZkFkzpU97ZsDFUqD1JvHiFgfK", 2 * 1e16},
	{"TCTBjWcY1AgYd2gt3vWY8GcjDmAwN36JrJv", 2 * 1e16},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network.  See "Utopia organization related parameters" in simnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"SCiaSErr5pS11tSK1RN1n7tnJfetUK7ZLkh", 1 * 1e17},
}

// BlockOneLedgerRegNet is the block one output ledger for the regression test
// network.  See "Utopia organization related parameters" in regnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerRegNet = []*TokenPayout{
	{"RsKrWb7Vny1jnzL1sDLgKTAteh9RZcRr5g6", 1 * 1e16},
	{"Rs8ca5cDALtsMVD4PV3xvFTC7dmuU1juvLv", 1 * 1e16},
	{"RsHzbGt6YajuHpurtpqXXHz57LmYZK8w9tX", 8 * 1e16},
}
