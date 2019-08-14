### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying ucd via ucctl](#ucctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Network Configuration](#NetworkConfig)
    2. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [Community](#ContactCommunity)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [Go Modules](#GoModules)
    4. [Module Hierarchy](#ModuleHierarchy)

<a name="About" />

### 1. About

ucd is a full node Utopia implementation written in [Go](https://golang.org),
and is licensed under the [copyfree](http://www.copyfree.org) ISC License.

This software is currently under active development.  It is extremely stable and
has been in production use since February 2016.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block.  It
ensures all individual transactions admitted to the pool follow the rules
required into the block chain and also includes the vast majority of the more
strict checks which filter transactions based on miner requirements ("standard"
transactions).

<a name="GettingStarted" />

### 2. Getting Started

<a name="Installation" />

**2.1 Installation**<br />

The first step is to install ucd.  The installation instructions can be found
[here](https://github.com/UtopiaCoinOrg/ucd/tree/master/README.md#Installation).

<a name="Configuration" />

**2.2 Configuration**<br />

ucd has a number of [configuration](https://godoc.org/github.com/UtopiaCoinOrg/ucd)
options, which can be viewed by running: `$ ucd --help`.

<a name="ucctlConfig" />

**2.3 Controlling and Querying ucd via ucctl**<br />

ucctl is a command line utility that can be used to both control and query ucd
via [RPC](https://www.wikipedia.org/wiki/Remote_procedure_call).  ucd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* ucd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* ucctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ ucctl --help`

<a name="Mining" />

**2.4 Mining**<br />
ucd supports the [getwork](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/json_rpc_api.mediawiki#getwork)
RPC.  The limited user cannot access this RPC.<br />

**1. Add the payment addresses with the `miningaddr` option.**<br />

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=DsExampleAddress1
miningaddr=DsExampleAddress2
```

**2. Add ucd's RPC TLS certificate to system Certificate Authority list.**<br />

`cgminer` uses [curl](https://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `ucd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**<br />

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.ucd/rpc.cert /usr/share/ca-certificates/ucd.crt`<br />
2. Add ucd.crt to /etc/ca-certificates.conf: `# echo ucd.crt >> /etc/ca-certificates.conf`<br />
3. Update the CA certificate list: `# update-ca-certificates`<br />

**3. Set your mining software url to use https.**<br />

`$ cgminer -o https://127.0.0.1:10509 -u rpcuser -p rpcpassword`

<a name="Help" />

### 3. Help

<a name="NetworkConfig" />

**3.1 Network Configuration**<br />
* [What Ports Are Used by Default?](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring ucd with Tor](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />

**3.2 Wallet**<br />

ucd was intentionally developed without an integrated wallet for security
reasons.  Please see [ucwallet](https://github.com/UtopiaCoinOrg/ucwallet) for more
information.

<a name="Contact" />

### 4. Contact

<a name="ContactCommunity" />

**4.1 Community**<br />

If you have any further questions you can find us at:

https://utopia.org/community

<a name="DeveloperResources" />

### 5. Developer Resources

<a name="ContributionGuidelines" />

**5.1 Code Contribution Guidelines**

* [Code Contribution Guidelines](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/code_contribution_guidelines.md)

<a name="JSONRPCReference" />

**5.2 JSON-RPC Reference**

* [JSON-RPC Reference](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/json_rpc_api.mediawiki)
* [RPC Examples](https://github.com/UtopiaCoinOrg/ucd/tree/master/docs/json_rpc_api.mediawiki#8-example-code)

<a name="GoModules" />

**5.3 Go Modules**

The following versioned modules are provided by ucd repository:

* [rpcclient/v3](https://github.com/UtopiaCoinOrg/ucd/tree/master/rpcclient) - Implements
  a robust and easy to use Websocket-enabled Utopia JSON-RPC client
* [ucjson/v3](https://github.com/UtopiaCoinOrg/ucd/tree/master/ucjson) - Provides
  infrastructure for working with Utopia JSON-RPC APIs
* [rpc/jsonrpc/types](https://github.com/UtopiaCoinOrg/ucd/tree/master/rpc/jsonrpc/types) -
  Provides concrete types via ucjson for the chain server JSON-RPC commands,
  return values, and notifications
* [wire](https://github.com/UtopiaCoinOrg/ucd/tree/master/wire) - Implements the
  Utopia wire protocol
* [peer](https://github.com/UtopiaCoinOrg/ucd/tree/master/peer) - Provides a common
  base for creating and managing Utopia network peers
* [blockchain](https://github.com/UtopiaCoinOrg/ucd/tree/master/blockchain) -
  Implements Utopia block handling and chain selection rules
  * [stake/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/blockchain/stake) -
    Provides an API for working with stake transactions and other portions
    related to the Proof-of-Stake (PoS) system
* [txscript/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/txscript) -
  Implements the Utopia transaction scripting language
* [ucec](https://github.com/UtopiaCoinOrg/ucd/tree/master/ucec) - Provides constants
  for the supported cryptographic signatures supported by Utopia scripts
  * [secp256k1](https://github.com/UtopiaCoinOrg/ucd/tree/master/ucec/secp256k1) -
    Implements the secp256k1 elliptic curve
  * [edwards/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/ucec/edwards) -
    Implements the edwards25519 twisted Edwards curve
* [database/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/database) -
  Provides a database interface for the Utopia block chain
* [mempool/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/mempool) - Provides a
  policy-enforced pool of unmined Utopia transactions
* [ucutil/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/ucutil) - Provides
  Utopia-specific convenience functions and types
* [chaincfg/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/chaincfg) - Defines
  chain configuration parameters for the standard Utopia networks and allows
  callers to define their own custom Utopia networks for testing puproses
  * [chainhash](https://github.com/UtopiaCoinOrg/ucd/tree/master/chaincfg/chainhash) -
    Provides a generic hash type and associated functions that allows the
    specific hash algorithm to be abstracted
* [certgen](https://github.com/UtopiaCoinOrg/ucd/tree/master/certgen) - Provides a
  function for creating a new TLS certificate key pair, typically used for
  encrypting RPC and websocket communications
* [addrmgr](https://github.com/UtopiaCoinOrg/ucd/tree/master/addrmgr) - Provides a
  concurrency safe Utopia network address manager
* [connmgr](https://github.com/UtopiaCoinOrg/ucd/tree/master/connmgr) - Implements a
  generic Utopia network connection manager
* [hdkeychain/v2](https://github.com/UtopiaCoinOrg/ucd/tree/master/hdkeychain) -
  Provides an API for working with  Utopia hierarchical deterministic extended
  keys
* [gcs](https://github.com/UtopiaCoinOrg/ucd/tree/master/gcs) - Provides an API for
  building and using Golomb-coded set filters useful for light clients such as
  SPV wallets
* [fees](https://github.com/UtopiaCoinOrg/ucd/tree/master/fees) - Provides methods for
  tracking and estimating fee rates for new transactions to be mined into the
  network
* [lru](https://github.com/UtopiaCoinOrg/ucd/tree/master/lru) - Implements a generic
  concurrent safe least-recently-used cache with near O(1) perf

<a name="ModuleHierarchy" />

**5.4 Module Hierarchy**

The following diagram shows an overview of the hierarchy for the modules
provided by the ucd repository.

![Module Hierarchy](./assets/module_hierarchy.svg)
