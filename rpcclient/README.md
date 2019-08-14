rpcclient
=========

[![Build Status](https://img.shields.io/travis/Utopia/ucd.svg)](https://travis-ci.org/Utopia/ucd)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/UtopiaCoinOrg/ucd/rpcclient)

rpcclient implements a Websocket-enabled Utopia JSON-RPC client package written
in [Go](https://golang.org/).  It provides a robust and easy to use client for
interfacing with a Utopia RPC server that uses a ucd compatible Utopia
JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](https://godoc.org/github.com/UtopiaCoinOrg/ucd/rpcclient)
* [ucd Websockets Example](https://github.com/UtopiaCoinOrg/ucd/tree/master/rpcclient/examples/ucdwebsockets)
  Connects to a ucd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [ucwallet Websockets Example](https://github.com/UtopiaCoinOrg/ucd/tree/master/rpcclient/examples/ucwalletwebsockets)  
  Connects to a ucwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign

## Major Features

* Supports Websockets (ucd/ucwallet) and HTTP POST mode (bitcoin core-like)
* Provides callback and registration functions for ucd/ucwallet notifications
* Supports ucd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/UtopiaCoinOrg/ucd/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
