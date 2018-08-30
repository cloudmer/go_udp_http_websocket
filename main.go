package main

import (
	"fmt"
	_ "bnwUdp/init"
	"bnwUdp/share"
	"bnwUdp/library/logger"
)

func main() {
	fmt.Println(share.ShareConfig)
	logger.Info("ffffff")
}
