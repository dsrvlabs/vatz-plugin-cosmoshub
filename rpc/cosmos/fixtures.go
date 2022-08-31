package common

const fixtureStatus string = `{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "7",
        "block": "10",
        "app": "0"
      },
      "id": "64c922566ac244ecc84307ff35ad493ec79b5796",
      "listen_addr": "tcp://0.0.0.0:26656",
      "network": "heimdall-137",
      "version": "0.32.7",
      "channels": "4020212223303800",
      "moniker": "knuth",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://127.0.0.1:26657"
      }
    },
    "sync_info": {
      "latest_block_hash": "D6739DE74A18EB39097E96647E6D4B47E819F558310FFEF755BC860CFB89530A",
      "latest_app_hash": "D99D073ACB8D57EC1B6C4E9311DCF84606CBAA4EA2FBEEC48C791EE8FE9E83E8",
      "latest_block_height": "9869741",
      "latest_block_time": "2022-06-23T18:02:34.128600247Z",
      "catching_up": false
    },
    "validator_info": {
      "address": "6DE3B706B342382E652472A9C03BBC7106C17ACF",
      "pub_key": {
        "type": "tendermint/PubKeySecp256k1",
        "value": "BIsYhK0Cu722NjcjMs5aFNAbOEeQczOrhkvPBB5yd3v05Z6NvJVjnioDlC6KjJBT5oZB5qC74vRcXtZbkg95N2s="
      },
      "voting_power": "0"
    }
  }
}`
