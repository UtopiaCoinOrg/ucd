jsonrpc/types
=============

[![Build Status](https://travis-ci.org/Utopia/ucd.png?branch=master)](https://travis-ci.org/Utopia/ucd)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/Utopia/ucd/rpc/jsonrpc/types)

Package types implements concrete types for marshalling to and from the ucd
JSON-RPC commands, return values, and notifications.  A comprehensive suite of
tests is provided to ensure proper functionality.

The provided types are automatically registered with
[ucjson](https://github.com/Utopia/ucd/tree/master/ucjson) when the package
is imported.  Although this package was primarily written for ucd, it has
intentionally been designed so it can be used as a standalone package for any
projects needing to marshal to and from ucd JSON-RPC requests and responses.

## Installation and Updating

```bash
$ go get -u github.com/Utopia/ucd/rpc/jsonrpc/types
```

## License

Package types is licensed under the [copyfree](http://copyfree.org) ISC License.
