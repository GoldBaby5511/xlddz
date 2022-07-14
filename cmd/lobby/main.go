package main

import (
	"github.com/GoldBaby5511/go-mango-core/gate"
	_ "mango/cmd/lobby/business"
)

func main() {
	gate.Start("login")
}
