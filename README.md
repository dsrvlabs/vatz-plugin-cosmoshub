# vatz-plugin-cosmoshub
Vatz plugin for cosmoshub node monitoring

## Plugins
- node_block_sync : monitor block sync status
- is_alived : monitor `gaiad` running status
- peer_count : monitor the number of peers
- active_status : monitor the validator include in active set

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
2022-09-14T08:35:27+02:00 INF Start main=statusCollector
2022-09-14T08:35:27+02:00 INF Register module=grpc
2022-09-14T08:35:27+02:00 INF Start 127.0.0.1 9091 module=sdk
2022-09-14T08:35:27+02:00 INF Start module=grpc
2022-09-14T08:35:29+02:00 INF Execute module=grpc
2022-09-14T08:35:29+02:00 INF pluginFeature: string_value:"getBlockSync" main=main
2022-09-14T08:35:29+02:00 INF estimate policy=blockSyncEstimator
```
```
$ is_alived
2022-09-14T08:36:29+02:00 INF Register module=grpc
2022-09-14T08:36:29+02:00 INF Start 0.0.0.0 9098 module=sdk
2022-09-14T08:36:29+02:00 INF Start module=grpc
2022-09-14T08:36:29+02:00 INF Execute module=grpc
2022-09-14T08:36:29+02:00 INF gaiad Process alive process=up
```
```
$ peer_count
2022-09-28T09:21:22Z INF Register module=grpc
2022-09-28T09:21:22Z INF Start 127.0.0.1 9095 module=sdk
2022-09-28T09:21:22Z INF Start module=grpc
2022-09-28T09:21:32Z INF Execute module=grpc
2022-09-28T09:21:32Z INF [cosmos-mainnet-sentry2-do]Good: peer_count is 56 moudle=plugin
```
```
$ active_status -valoperAddr <VALIDATOR_OPERATOR_ADDRESS>
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
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number, default 9091 (default 9091)
```
- is_alived
```
Usage of is_alived:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number, default 9098 (default 9098)
```
- peer_count
```
Usage of peer_count:
  -addr string
        IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -minPeer int
        minimum peer count, default 5 (default 5)
  -port int
        Port number, default 9091 (default 9091)
```
- active_status
```
Usage of active_status:
  -addr string
    	Listening address (default "127.0.0.1")
  -port int
    	Listening port (default 9100)
  -rpcURI string
    	CosmosHub RPC URI Address (default "http://localhost:1317")
  -valoperAddr string
    	CosmosHub validator operator address
```
