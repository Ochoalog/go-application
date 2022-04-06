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

	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
