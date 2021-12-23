package main

import (
	_ "xlddz/cmd/room/business"
	"xlddz/pkg/gate"
)

func main() {
	gate.Start("room")
}
