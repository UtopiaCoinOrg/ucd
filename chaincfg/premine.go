// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2018 The Utopia developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.

var BlockOneLedgerMainNet = []*TokenPayout{
	{"UCLf91Euh2a9baoeCyoAmuRnYXLbczGbShy", 5000000 * 1e8},
	{"UCQGvbm7ghGuY1UCnbXC35vgRL1HD5RdUSG", 5000000 * 1e8},
	{"UCem9jrdHQ7fniFptdeP3dYQMQSxRetssCz", 5000000 * 1e8},
	{"UCNnwds81m9CtCcMRrL8mkNZrPRXEXTH8jc", 5000000 * 1e8},
	{"UCVt6qiQ34VyrAQ6n6rDGb68phT9e2MahyP", 5000000 * 1e8},
	{"UCdjU6mEgiC55A2ZPBVC18uzcRwrWs6Sjeo", 5000000 * 1e8},
	{"UCjBYf4djWSgyiqPnFnoQMRL44hRRASmyEk", 5000000 * 1e8},
	{"UCVvv4a8DvEdEYPQVtmhpEvdTCxCGHr87Km", 5000000 * 1e8},
	{"UCT4zACLxUmk6uXhmEYaPoVMLrVQWBM5Hdu", 500000 * 1e8},
	{"UCRyfGoG45n3LE8TytYHqYxYsXzMdyYEHec", 500000 * 1e8},
	{"UCTPBVxSwSys1KE9y8mMD6adfHQ854cke59", 500000 * 1e8},
	{"UCXyfQk4AAsR88TChE1F7vAoq9msiPZE7Nm", 500000 * 1e8},
	{"UCQvTHi2tNgeCs9GP47HCEZJrPMCvzHvaZE", 958 * 1e14},
}

// BlockOneLedgerTestNet3 is the block one output ledger for testnet version 3.
var BlockOneLedgerTestNet3 = []*TokenPayout{
	{"TCe6sr1wNcce13JC624taF4jR9NVdnJEJsr", 2 * 1e16},
	{"TCPUBwKcPD8fdFuiQgHJQrmRunAtP7cVh5S", 2 * 1e16},
	{"TCYSkvH9iqTR4b2gcqEN5v8m1gxTyFZE1kN", 2 * 1e16},
	{"TCg5zdr3kcmUijYBTCd5iPXsGtRp3Gk8iUE", 2 * 1e16},
	{"TCez1JEsHo65iLNWirHgUCGVXBRusB5wuQt", 2 * 1e16},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network.  See "Utopia organization related parameters" in simnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"SChTs8bcyNoffM9hZEQCyWXtuEKWFPJ1nK4", 1 * 1e17},
}

// BlockOneLedgerRegNet is the block one output ledger for the regression test
// network.  See "Utopia organization related parameters" in regnetparams.go for
// information on how to spend these outputs.
var BlockOneLedgerRegNet = []*TokenPayout{
	{"RsKrWb7Vny1jnzL1sDLgKTAteh9RZcRr5g6", 1 * 1e16},
	{"Rs8ca5cDALtsMVD4PV3xvFTC7dmuU1juvLv", 1 * 1e16},
	{"RsHzbGt6YajuHpurtpqXXHz57LmYZK8w9tX", 8 * 1e16},
}
