package main

import (
	_ "mango/cmd/gateway/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("gateway")
}
