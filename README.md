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
> Run as default config
```
$ node_block_sync
```
2023-01-02T09:26:47Z INF Register module=grpc
2023-01-02T09:26:47Z INF Start 127.0.0.1 9091 module=sdk
2023-01-02T09:26:47Z INF Start module=grpc
2023-01-02T09:26:56Z INF Execute module=grpc
2023-01-02T09:26:56Z INF previous block height: 0, latest block height: 3791025 module=plugin
2023-01-02T09:26:56Z DBG block height increasing module=plugin
2023-01-02T09:27:25Z INF Execute module=grpc
2023-01-02T09:27:25Z INF previous block height: 3791025, latest block height: 3791030 module=plugin
2023-01-02T09:27:25Z DBG block height increasing module=plugin
```
$ node_is_alived
```
2023-01-02T09:26:47Z INF Register module=grpc
2023-01-02T09:26:47Z INF Start 127.0.0.1 9098 module=sdk
2023-01-02T09:26:47Z INF Start module=grpc
2023-01-02T09:26:56Z INF Execute module=grpc
2023-01-02T09:26:56Z INF HEALTHY process=up
2023-01-02T09:27:25Z INF Execute module=grpc
2023-01-02T09:27:25Z INF HEALTHY process=up
```
```
$ node_peer_count
2023-01-02T09:26:47Z INF Register module=grpc
2023-01-02T09:26:47Z INF Start 127.0.0.1 9091 module=sdk
2023-01-02T09:26:47Z INF Start module=grpc
2023-01-02T09:26:56Z INF Execute module=grpc
2023-01-02T09:26:56Z INF Good: peer_count is 10 moudle=plugin
2023-01-02T09:27:25Z INF Execute module=grpc
2023-01-02T09:27:25Z INF Good: peer_count is 10 moudle=plugin
```
```
$ node_active_status -valoperAddr <VALIDATOR_OPERATOR_ADDRESS>
2023-01-02T09:26:47Z INF Register module=grpc
2023-01-02T09:26:47Z INF Start 127.0.0.1 9100 module=sdk
2023-01-02T09:26:47Z INF Start module=grpc
2023-01-02T09:26:56Z INF Execute module=grpc
2023-01-02T09:26:56Z DBG Validator bonded. included active set module=plugin
2023-01-02T09:27:25Z INF Execute module=grpc
2023-01-02T09:27:26Z DBG Validator bonded. included active set module=plugin
```
```
# Your node have to enable API configuration ({HOME_DIR}/config/app.toml)
$ node_governance_alarm -port 9097 -apiPort <API server port{default is 1317}> -voterAddr <Account Address>
2023-01-02T09:27:32Z INF Register module=grpc
2023-01-02T09:27:32Z INF Start 127.0.0.1 9097 module=sdk
2023-01-02T09:27:32Z INF Start module=grpc
2023-01-02T09:27:55Z INF Execute module=grpc
2023-01-02T09:28:04Z DBG DEBUG : tmp == proposalId module=plugin
2023-01-02T09:28:04Z INF Lastest proposal is #2
 module=plugin
2023-01-02T09:28:25Z INF Execute module=grpc
2023-01-02T09:28:26Z DBG DEBUG : tmp == proposalId module=plugin
2023-01-02T09:28:26Z INF Lastest proposal is #2
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
	Listening port (default 9091)
  -rpcURI string
	Tendermint RPC URI Address (default "http://localhost:26657")
```
- node_is_alived
```
Usage of node_is_alived:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number, default 9098 (default 9098)
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
        Port number, default 9091 (default 9091)
  -rpcAddr string
    	RPC Address, default http://localhost:26657 (default "https://localhost:26657")
```
- node_active_status
```
Usage of node_active_status:
  -addr string
    	Listening address (default "127.0.0.1")
  -port int
    	Listening port (default 9100)
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
    	Port number, default 9091 (default 9091)
  -proposalId uint
    	Need to know last proposal id
  -voterAddr string
    	Need to voter address (default "address")
```
