# vatz-plugin-cosmoshub
Vatz plugin for cosmoshub node monitoring

## Plugins
- cosmos-sdk-blocksync : monitor block sync status
- is_alived : monitor `gaiad` running status
- peer_count : monitor the number of peers

## Installation and Usage
> Please make sure [Vatz](https://github.com/dsrvlabs/vatz) is running with proper configuration. [Vatz Installation Guide](https://github.com/dsrvlabs/vatz/blob/main/docs/installation.md)

### Install
```
$ make install
```

### Run
> Run as default config
```
$ cosmos-sdk-blocksync
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
TBD
```

## Command line arguments
- cosmos-sdk-blocksync
```
Usage of cosmos-sdk-blocksync:
  -addr string
    	IP Address(e.g. 0.0.0.0, 127.0.0.1) (default "127.0.0.1")
  -port int
    	Port number, default 9091 (default 9091)
```
- is_alived
```
Usage of is_alived:
TBD
```
- peer_count
```
Usage of peer_count:
TBD
```
