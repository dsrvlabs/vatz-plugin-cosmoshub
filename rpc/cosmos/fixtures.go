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

const fixtureHealth string = `{
  "id": 0,
  "jsonrpc": "2.0",
  "result": {}
}`

const fixtureBondStatus string = `{
  "validator": {
    "operator_address": "cosmosvaloper1wlagucxdxvsmvj6330864x8q3vxz4x02rmvmsu",
    "consensus_pubkey": {
      "@type": "/cosmos.crypto.ed25519.PubKey",
      "key": "efOai5jzck+C46Zt8ruUcD1w2E7wnDnL9u2ATsODIPg="
    },
    "jailed": false,
    "status": "BOND_STATUS_BONDED",
    "tokens": "36504214653",
    "delegator_shares": "36504214653.000000000000000000",
    "description": {
      "moniker": "DSRV",
      "identity": "CC434B6FE536F51B",
      "website": "https://dsrvlabs.com",
      "security_contact": "",
      "details": ""
    },
    "unbonding_height": "12485537",
    "unbonding_time": "2022-11-07T15:27:59.955383084Z",
    "commission": {
      "commission_rates": {
        "rate": "0.100000000000000000",
        "max_rate": "0.900000000000000000",
        "max_change_rate": "0.100000000000000000"
      },
      "update_time": "2021-02-26T15:02:19.028351586Z"
    },
    "min_self_delegation": "1"
  }
}`

const fixturePeerCount string = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "listening": true,
    "listeners": [
      "Listener(@)"
    ],
    "n_peers": "41",
    "peers": [
      {
        "node_info": {
          "protocol_version": {
            "p2p": "8",
            "block": "11",
            "app": "0"
          },
          "id": "af77fd96cb62c6011272ee67390e540504b47fd9",
          "listen_addr": "tcp://51.222.42.205:26656",
          "network": "agoric-3",
          "version": "0.34.19",
          "channels": "40202122233038606100",
          "moniker": "FWQREb8H2n",
          "other": {
            "tx_index": "on",
            "rpc_address": "tcp://0.0.0.0:26657"
          }
        },
        "is_outbound": true,
        "connection_status": {
          "Duration": "526436041912849",
          "SendMonitor": {
            "Start": "2022-11-02T02:43:38.1Z",
            "Bytes": "4499155655",
            "Samples": "1459495",
            "InstRate": "1518",
            "CurRate": "18985",
            "AvgRate": "8546",
            "PeakRate": "1183830",
            "BytesRem": "0",
            "Duration": "526436040000000",
            "Idle": "380000000",
            "TimeRem": "0",
            "Progress": 0,
            "Active": true
          },
          "RecvMonitor": {
            "Start": "2022-11-02T02:43:38.1Z",
            "Bytes": "4827793450",
            "Samples": "1382543",
            "InstRate": "1391",
            "CurRate": "27642",
            "AvgRate": "9171",
            "PeakRate": "2933610",
            "BytesRem": "0",
            "Duration": "526436040000000",
            "Idle": "320000000",
            "TimeRem": "0",
            "Progress": 0,
            "Active": true
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "638"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "6",
              "RecentlySent": "6739"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "17400"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "7",
              "RecentlySent": "55331"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "24"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "6",
              "RecentlySent": "0"
            },
            {
              "ID": 96,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 97,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "3",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        },
        "remote_ip": "51.222.42.205"
      },
	  {
        "node_info": {
          "protocol_version": {
            "p2p": "8",
            "block": "11",
            "app": "0"
          },
          "id": "20dcba485cde15ee5e1ae93e36280c3acf77ef91",
          "listen_addr": "tcp://0.0.0.0:26656",
          "network": "agoric-3",
          "version": "0.34.19",
          "channels": "40202122233038606100",
          "moniker": "infstones",
          "other": {
            "tx_index": "on",
            "rpc_address": "tcp://0.0.0.0:26657"
          }
        },
        "is_outbound": false,
        "connection_status": {
          "Duration": "421506461547522",
          "SendMonitor": {
            "Start": "2022-11-03T07:53:27.94Z",
            "Bytes": "4140742572",
            "Samples": "1189362",
            "InstRate": "2717",
            "CurRate": "11792",
            "AvgRate": "9824",
            "PeakRate": "1443030",
            "BytesRem": "0",
            "Duration": "421506460000000",
            "Idle": "120000000",
            "TimeRem": "0",
            "Progress": 0,
            "Active": true
          },
          "RecvMonitor": {
            "Start": "2022-11-03T07:53:27.94Z",
            "Bytes": "3142893617",
            "Samples": "1008980",
            "InstRate": "56",
            "CurRate": "7735",
            "AvgRate": "7456",
            "PeakRate": "878180",
            "BytesRem": "0",
            "Duration": "421506460000000",
            "Idle": "720000000",
            "TimeRem": "0",
            "Progress": 0,
            "Active": true
          },
          "Channels": [
            {
              "ID": 48,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "911"
            },
            {
              "ID": 64,
              "SendQueueCapacity": "1000",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 32,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "6",
              "RecentlySent": "6748"
            },
            {
              "ID": 33,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "10",
              "RecentlySent": "17284"
            },
            {
              "ID": 34,
              "SendQueueCapacity": "100",
              "SendQueueSize": "0",
              "Priority": "7",
              "RecentlySent": "66884"
            },
            {
              "ID": 35,
              "SendQueueCapacity": "2",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            },
            {
              "ID": 56,
              "SendQueueCapacity": "1",
              "SendQueueSize": "0",
              "Priority": "6",
              "RecentlySent": "0"
            },
            {
              "ID": 96,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "5",
              "RecentlySent": "0"
            },
            {
              "ID": 97,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "3",
              "RecentlySent": "0"
            },
            {
              "ID": 0,
              "SendQueueCapacity": "10",
              "SendQueueSize": "0",
              "Priority": "1",
              "RecentlySent": "0"
            }
          ]
        },
        "remote_ip": "18.116.45.237"
      }
    ]
  }
}`
