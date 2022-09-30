# Tixcash Core integration/ repository

EVM-compatible chain secured by the Lachesis consensus algorithm.

Tixcash (TXH) is a sustainable cryptocurrency modeled after Satoshi Nakamoto’s vision for Bitcoin. It is a decentralized, peer-to-peer transactional currency designed to offer a solution to the problem posed by the exponential increase in energy consumed by Bitcoin and other proof-of-work currencies. Proof-of-work mining is environmentally unsustainable due to the electricity used by high-powered mining hardware.

Tixcash utilizes an energy efficient proof-of-stake algorithm, can be mined on any computer or server, and will never require specialized mining equipment. The Green Protocol offers a simple solution to Bitcoin sustainability issues and provides a faster, more scalable blockchain that is better suited for daily transactional use.

Tixcash is based on the latest 3 generation blockchain technology, and its fast transaction features guaranteed instant confirmation transactions. Decentralized blockchain voting providing for consensus based advancement of the current technology used to secure the network and provide the above features, each full node is secured. and the smart contract pledge mining method will be used to obtain income; the international common floating pledge mining mode is being adopted. 

With the advanced EVM model, the Tixcash chain will run in a more sustanable and more efficient way, since the EVM technology, adopted by ETH and more 3rd generation coins, provides a system with lower gas fee and better transaction efficiency. 

Furthermore, the standard hash system ensures compatibility when it comes to interactions with other mainstream coins.

## Tixcash  shares allocation

The distribution plan of a total of 100 billion Tixcash coins will be announced worldwide in the near future. The distribution method will include node mining rewards10%, pledge rewards50%, foundations13%, team holdings10%, community development and ecological construction funds15%, pre-mined2%,etc.
![Tixcash](https://user-images.githubusercontent.com/65695574/193205589-f7c7f9eb-9059-43d2-b405-468f73293b37.PNG)

## Building the source

Building opera requires both a Go (version 1.14 or later) and a C compiler. You can install them using your favourite package manager. Once the dependencies are installed, run

make opera
The build output is build/opera executable.

## Running `opera`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `opera` instance.

### Configuration

As an alternative to passing the numerous flags to the `opera` binary, you can also pass a
configuration file via:

```shell
$ opera --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ opera --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `opera validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ opera --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`opera` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ opera --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ opera --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ opera --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ opera --datadir /path/to/datadir account new # create new account
$ opera --datadir /path/to/datadir attach # attach to IPC
```
### Operating a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ opera --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ opera --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ opera --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ opera --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```
