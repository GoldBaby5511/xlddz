package main

import (
	_ "mango/cmd/list/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("list")
}
