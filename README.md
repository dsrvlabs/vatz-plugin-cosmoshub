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
2022-11-23T08:44:08Z INF Register module=grpc
2022-11-23T08:44:08Z INF Start 127.0.0.1 9091 module=sdk
2022-11-23T08:44:08Z INF Start module=grpc
2022-11-23T08:44:09Z INF Execute module=grpc
2022-11-23T08:44:09Z INF previous block height: 0, latest block height: 11662932 module=plugin
2022-11-23T08:44:09Z DBG block height increasing module=plugin
2022-11-23T08:44:13Z INF Execute module=grpc
2022-11-23T08:44:13Z INF previous block height: 11662932, latest block height: 11662933 module=plugin
2022-11-23T08:44:13Z DBG block height increasing module=plugin
2022-11-23T08:44:18Z INF Execute module=grpc
2022-11-23T08:44:19Z INF previous block height: 11662933, latest block height: 11662935 module=plugin
2022-11-23T08:44:19Z DBG block height increasing module=plugin
```
```
$ node_is_alived
2022-09-14T08:36:29+02:00 INF Register module=grpc
2022-09-14T08:36:29+02:00 INF Start 0.0.0.0 9098 module=sdk
2022-09-14T08:36:29+02:00 INF Start module=grpc
2022-09-14T08:36:29+02:00 INF Execute module=grpc
2022-09-14T08:36:29+02:00 INF gaiad Process alive process=up
```
```
$ node_peer_count
2022-09-28T09:21:22Z INF Register module=grpc
2022-09-28T09:21:22Z INF Start 127.0.0.1 9095 module=sdk
2022-09-28T09:21:22Z INF Start module=grpc
2022-09-28T09:21:32Z INF Execute module=grpc
2022-09-28T09:21:32Z INF [cosmos-mainnet-sentry2-do]Good: peer_count is 56 moudle=plugin
```
```
$ node_active_status -valoperAddr <VALIDATOR_OPERATOR_ADDRESS>
2022-10-26T11:42:35+02:00 INF Register module=grpc
2022-10-26T11:42:35+02:00 INF Start 127.0.0.1 9100 module=sdk
2022-10-26T11:42:35+02:00 INF Start module=grpc
2022-10-26T11:42:42+02:00 INF Execute module=grpc
2022-10-26T11:42:42+02:00 DBG Validator bonded. included active set module=plugin
```
```
# Your node have to enable API configuration ({HOME_DIR}/config/app.toml)
$ node_governance_alarm -port 9093 -apiPort <API server port{default is 1317}> -voterAddr <Account Address>
Lastest proposal is #195 module=plugin
2022-12-23T07:52:52Z INF Proposal #N: Dit not vote module=plugin
2022-12-23T07:52:52Z INF Register module=grpc
2022-12-23T07:52:52Z INF Start 127.0.0.1 9093 module=sdk
2022-12-23T07:52:52Z INF Start module=grpc
2022-12-23T07:53:15Z INF Execute module=grpc
2022-12-23T07:53:15Z DBG DEBUG : tmp == proposalId module=plugin
2022-12-23T07:53:15Z INF Proposal #N: Dit not vote module=plugin
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
