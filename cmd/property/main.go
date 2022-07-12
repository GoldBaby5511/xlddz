package main

import (
	_ "mango/cmd/property/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("property")
}
