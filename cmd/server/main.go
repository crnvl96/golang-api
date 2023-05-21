package main

import (
	"fmt"

	"github.com/crnvl96/go-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
