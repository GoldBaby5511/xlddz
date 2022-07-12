package main

import (
	_ "mango/cmd/table/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("table")
}
