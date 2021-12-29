package main

import (
	_ "mango/cmd/login/business"
	"mango/pkg/gate"
)

func main() {
	gate.Start("login")
}
