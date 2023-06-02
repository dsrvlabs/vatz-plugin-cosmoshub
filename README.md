# vatz-plugin-cosmoshub
Vatz plugin for cosmoshub node monitoring

## Plugins
- node_block_sync : monitor block sync status
- node_is_alived : monitor process running status
- node_peer_count : monitor the number of peers
- node_active_status : monitor the validator include in active set
- node_governance_alarm : monitor the new governance proposal and whether or not to vote

## Installation and Usage
> Please make sure [Vatz](https://github.com/dsrvlabs/vatz) is running with proper configuration. [Vatz Installation Guide](https://github.com/dsrvlabs/vatz/blob/main/docs/installation.md)

### Install Plugins
- Install with source
```
$ git clone https://github.com/dsrvlabs/vatz-plugin-cosmoshub.git
$ cd vatz-plugin-cosmoshub
$ make install
```
- Install with Vatz CLI command
```
$ vatz plugin install --help
Install new plugin

Usage:
   plugin install [flags]

Examples:
vatz plugin install github.com/dsrvlabs/<somewhere> name

Flags:
  -h, --help   help for install
```
> please make sure install path for the plugins repository URL.
```
$ vatz plugin install github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_block_sync vatz_block_sync
$ vatz plugin install github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_is_alived vatz_node_is_alived
$ vatz plugin install github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_peer_count vatz_peer_count
$ vatz plugin install github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_active_status vatz_active_status
$ vatz plugin install github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_governance_alarm vatz_gov_alarm
```
- Check plugins list with Vatz CLI command
```
$ vatz plugin list
2023-01-03T09:21:34Z INF List plugins module=plugin
2023-01-03T09:21:34Z INF List module=plugin
2023-01-03T09:21:34Z INF newReader /root/.vatz/vatz.db module=db
2023-01-03T09:21:34Z INF Create DB Instance module=db
2023-01-03T09:21:34Z INF List Plugin module=db
+---------------------+---------------------+-------------------------------------------------------------------------+---------+
| NAME                | INSTALL DATA        | REPOSITORY                                                              | VERSION |
+---------------------+---------------------+-------------------------------------------------------------------------+---------+
| vatz_block_sync     | 2023-01-02 09:13:19 | github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_block_sync       | latest  |
| vatz_node_is_alived | 2023-01-02 09:13:43 | github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_is_alived        | latest  |
| vatz_peer_count     | 2023-01-02 09:14:05 | github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_peer_count       | latest  |
| vatz_active_status  | 2023-01-02 09:14:41 | github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_active_status    | latest  |
| vatz_gov_alarm      | 2023-01-02 09:15:00 | github.com/dsrvlabs/vatz-plugin-cosmoshub/plugins/node_governance_alarm | latest  |
+---------------------+---------------------+-------------------------------------------------------------------------+---------+
```

### Run
> Run as default config or option flags
```
$ node_block_sync
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10001 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z INF previous block height: 0, latest block height: 5969512 module=plugin
2023-05-31T07:08:10Z DBG block height increasing module=plugin
```
```
$ node_is_alived
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10002 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z INF HEALTHY process=up
2023-05-31T07:08:40Z INF Execute module=grpc
2023-05-31T07:08:40Z INF HEALTHY process=up
```
```
$ node_peer_count
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10003 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z INF Good: peer_count is 50 moudle=plugin
2023-05-31T07:08:40Z INF Execute module=grpc
2023-05-31T07:08:40Z INF Good: peer_count is 50 moudle=plugin
```
```
$ node_active_status -valoperAddr <VALIDATOR_OPERATOR_ADDRESS>
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10004 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z DBG Validator bonded. included active set module=plugin
2023-05-31T07:08:40Z INF Execute module=grpc
2023-05-31T07:08:40Z DBG Validator bonded. included active set module=plugin
```
```
# Your node have to enable API configuration ({HOME_DIR}/config/app.toml)
$ node_governance_alarm -apiPort <API server port{default is 1317}> -voterAddr <Account Address>
2023-05-31T07:07:36Z INF Register module=grpc
2023-05-31T07:07:36Z INF Start 127.0.0.1 10005 module=sdk
2023-05-31T07:07:36Z INF Start module=grpc
2023-05-31T07:08:10Z INF Execute module=grpc
2023-05-31T07:08:10Z DBG DEBUG : tmp == proposalId module=plugin
2023-05-31T07:08:10Z INF Lastest proposal is #51
 module=plugin
```
## Command line arguments
- node_block_sync
```
Usage of node_block_sync:
  -addr string
	Listening address (default "127.0.0.1")
  -critical int
	block height stucked count to raise critical level of alert (default 3)
  -port int
	Listening port (default 10001)
  -rpcURI string
	Tendermint RPC URI Address (default "http://localhost:26657")
```
- node_is_alived
```
Usage of node_is_alived:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number (default 10002)
  -rpcAddr string
    	RPC addrest:port (e.g. http://127.0.0.1:26667) (default "http://localhost:26657")
```
- node_peer_count
```
Usage of node_peer_count:
  -addr string
        IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -minPeer int
        minimum peer count, default 5 (default 5)
  -port int
        Port number (default 10003)
  -rpcAddr string
    	RPC Address, default http://localhost:26657 (default "https://localhost:26657")
```
- node_active_status
```
Usage of node_active_status:
  -addr string
    	Listening address (default "127.0.0.1")
  -port int
    	Listening port (default 10004)
  -rpcURI string
    	CosmosHub RPC URI Address (default "http://localhost:1317")
  -valoperAddr string
    	CosmosHub validator operator address
```
- node_governance_alarm
```
Usage of node_governance_alarm:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -apiPort uint
    	Need to know proposal id (default 1317)
  -port int
    	Port number (default 10005)
  -proposalId uint
    	Need to know last proposal id
  -voterAddr string
    	Need to voter address (default "address")
```

## TroubleShooting
1. Encountered issue related with `Device or Resource Busy` or `Too many open files` error.
 - Check your open file limit and recommended to increase it.
 ```
 $ ulimit -n
 1000000
 ```
