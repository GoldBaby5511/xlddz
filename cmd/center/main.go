package main

import (
	_ "xlddz/cmd/center/business"
	"xlddz/pkg/gate"
)

func main() {
	gate.Start("center")
}
