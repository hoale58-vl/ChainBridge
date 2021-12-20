module github.com/ChainSafe/ChainBridge

go 1.15

replace github.com/ChainSafe/chainbridge-utils => /Users/black/Workspace/tokoin/chain/bridge/ChainBridge/chainbridge-utils

require (
	github.com/ChainSafe/chainbridge-utils v1.0.6
	github.com/ChainSafe/log15 v1.0.0
	github.com/centrifuge/go-substrate-rpc-client/v4 v4.0.0
	github.com/ethereum/go-ethereum v1.10.12
	github.com/prometheus/client_golang v1.4.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
)
