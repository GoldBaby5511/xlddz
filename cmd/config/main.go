package main

import (
	"github.com/GoldBaby5511/go-mango-core/gate"
	_ "mango/cmd/config/business"
	_ "mango/cmd/config/conf"
)

func main() {
	gate.Start("config")
}
