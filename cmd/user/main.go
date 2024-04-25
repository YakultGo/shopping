package main

import (
	"shopping/config"
)

func main() {
	config.InitConfig()
	config.InitLogger("user")
}
