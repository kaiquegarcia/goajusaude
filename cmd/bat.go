package cmd

import (
	"fmt"
	"goajusaude/internal"
	"goajusaude/pkg/shared/ajusaude"
)

func RunBat(ajuSaudeClient ajusaude.AjuSaudeClient, args []string) {
	if len(args) == 0 {
		fmt.Println("Você precisa informar um ou mais documentos para executarmos a consulta")
		return
	}

	for index := 0; index < len(args); index++ {
		err := internal.CheckProtocol(ajuSaudeClient, args[index])
		if err != nil {
			fmt.Printf("Um erro aconteceu:\n%s\n----\n", err)
		}
		fmt.Println("--------")
	}
}
