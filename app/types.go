package app

import (
	"bharvest.io/beramon/utils"
)

type Config struct {
	General struct {
		APIListenPort int  `toml:"api_listen_port"`
		Period        uint `toml:"period"`
	} `toml:"general"`
	Tg utils.TgConfig `toml:"tg"`
	CL struct {
		RPC              string `toml:"rpc"`
		ValidatorAddress string `toml:"validator_address"`
		MissThreshold    uint64 `toml:"miss_threshold"`
	} `toml:"cl"`
	EL struct {
		JsonRPC               string `toml:"json_rpc"`
		PeerThreshold         uint64 `toml:"peer_threshold"`
		TxpoolQueuedThreshold uint64 `toml:"txpool_queued_threshold"`
	} `toml:"el"`
}

type BaseApp struct {
	cfg *Config
}
