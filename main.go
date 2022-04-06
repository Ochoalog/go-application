package main

import (
	"app-command/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Início")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
