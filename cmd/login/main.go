package main

import (
	_ "xlddz/cmd/login/business"
	"xlddz/pkg/gate"
)

func main() {
	gate.Start("login")
}
