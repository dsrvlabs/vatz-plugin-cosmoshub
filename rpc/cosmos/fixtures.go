package rpc

const fixtureStatus string = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "8",
        "block": "11",
        "app": "0"
      },
      "id": "efdafdfe19abdf83486ea10eb153815406f60c22",
      "listen_addr": "tcp://0.0.0.0:26656",
      "network": "cosmoshub-4",
      "version": "v0.34.14",
      "channels": "40202122233038606100",
      "moniker": "wcloud",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://127.0.0.1:26657"
      }
    },
    "sync_info": {
      "latest_block_hash": "78A36AA3F36237CB093417A5C9E122EC4FB001BEB4C2BDC4CF8D56A893CBC1A0",
      "latest_app_hash": "2461DE8BDD6B4123DEC9707967F5C0588057F6A63B7D9118249784FCD6DE8867",
      "latest_block_height": "12232219",
      "latest_block_time": "2022-09-28T08:49:58.420923399Z",
      "earliest_block_hash": "18B38FBEC3137550A3B44813AC7A0A52A8929BBB0B204B64950C06F838FCCF28",
      "earliest_app_hash": "6EF9E7F96C2DF723EFAFE69F003C20E185F54578DF9D752AAA10C5D250EB7B67",
      "earliest_block_height": "12152001",
      "earliest_block_time": "2022-09-22T05:48:45.040926776Z",
      "catching_up": false
    },
    "validator_info": {
      "address": "BDF9854E9578A376A58D7335527D17AB68E982A7",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "P61RnQelwb7eXxEZ2UBXTEoZKG7loCoblsphSHSulUE="
      },
      "voting_power": "0"
    }
  }
}`
