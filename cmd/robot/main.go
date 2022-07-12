package main

import (
	_ "mango/cmd/robot/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("robot")
}
