package main

import (
	_ "baseServer/cmd/daemon/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("daemon")
}
