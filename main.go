package main

import (
	"command_ssh/app/controller"
	"command_ssh/app/utils"
)

func main() {
	utils.ReadFile()
	controller.InitApp()
}
