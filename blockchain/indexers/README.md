indexers
========

[![Build Status](https://travis-ci.org/Utopia/ucd.png?branch=master)](https://travis-ci.org/Utopia/ucd)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://godoc.org/github.com/UtopiaCoinOrg/ucd/blockchain/indexers?status.png)](https://godoc.org/github.com/UtopiaCoinOrg/ucd/blockchain/indexers)

Package indexers implements optional block chain indexes.

These indexes are typically used to enhance the amount of information available
via an RPC interface.

## Supported Indexers

- Transaction-by-hash (txbyhashidx) Index
  - Creates a mapping from the hash of each transaction to the block that
    contains it along with its offset and length within the serialized block
- Transaction-by-address (txbyaddridx) Index
  - Creates a mapping from every address to all transactions which either credit
    or debit the address
  - Requires the transaction-by-hash index
- Address-ever-seen (existsaddridx) Index
  - Stores a key with an empty value for every address that has ever existed 
    and was seen by the client
  - Requires the transaction-by-hash index
- Committed Filter (cfindexparentbucket) Index
  - Stores all committed filters and committed filter headers for all blocks in
    the main chain

## Installation

```bash
$ go get -u github.com/UtopiaCoinOrg/ucd/blockchain/indexers
```

## License

Package indexers is licensed under the [copyfree](http://copyfree.org) ISC
License.
