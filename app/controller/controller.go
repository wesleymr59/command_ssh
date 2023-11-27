package controller

import (
	"command_ssh/app/utils"
	"fmt"
	"log"

	"github.com/sfreiberg/simplessh"
)

func InitApp() {
	log.Println("Aplicação iniciada")
	var client *simplessh.Client
	var err error

	if client, err = simplessh.ConnectWithPassword("localhost:500", "root", "passwd"); err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	commandText := utils.ReadFile()
	for _, command := range commandText {
		if _, err := client.Exec(command); err != nil {
			log.Println(err)
		}
		a := fmt.Sprintf("O Comando %s, Foi executado com sucesso", command)
		log.Println(a)

	}
}
