package main

import (
	_ "xlddz/cmd/table/business"
	"xlddz/pkg/gate"
)

func main() {
	gate.Start("table")
}
