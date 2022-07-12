package main

import (
	_ "mango/cmd/room/business"
	"github.com/GoldBaby5511/go-mango-core/gate"
)

func main() {
	gate.Start("room")
}
