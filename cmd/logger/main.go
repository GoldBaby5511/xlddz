package main

import (
	_ "xlddz/cmd/logger/business"
	"xlddz/pkg/gate"
)

func main() {
	gate.Start("logger")
}
