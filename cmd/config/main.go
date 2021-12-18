package main

import (
	_ "xlddz/cmd/config/business"
	_ "xlddz/cmd/config/conf"
	"xlddz/pkg/gate"
)

func main() {
	gate.Start("config")
}
