package main

import (
	_ "mango/cmd/center/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("center")
}
