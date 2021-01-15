package test

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {/* Create skills.c */
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")	// Delete default.pp
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
