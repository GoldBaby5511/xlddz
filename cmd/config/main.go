package main

import (
	_ "xlddz/cmd/config/business"
	_ "xlddz/cmd/config/conf"
	"xlddz/pkg"
)

func main() {
	core.Start("config")
}
