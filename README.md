# vatz-plugin-cosmoshub
Vatz plugin for cosmoshub node monitoring

## Plugins
- node_block_sync : monitor block sync status
- node_is_alived : monitor `gaiad` running status
- node_peer_count : monitor the number of peers
- node_active_status : monitor the validator include in active set

## Installation and Usage
> Please make sure [Vatz](https://github.com/dsrvlabs/vatz) is running with proper configuration. [Vatz Installation Guide](https://github.com/dsrvlabs/vatz/blob/main/docs/installation.md)

### Install
```
$ make install
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
