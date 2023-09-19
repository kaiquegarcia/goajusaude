package cmd

import (
	"fmt"
	"goajusaude/internal"
	"goajusaude/pkg/shared/ajusaude"
)

func RunCLI(ajuSaudeClient ajusaude.AjuSaudeClient) {
	for {
		var document string
		fmt.Printf("Por favor, digite um documento para buscar no Aracaju Saúde (ou '-' para encerrar): ")
		fmt.Scan(&document)
		fmt.Println("")

		if document == "-" || document == "" {
			return
		}

		err := internal.CheckProtocol(ajuSaudeClient, document)
		if err != nil {
			fmt.Printf("Um erro aconteceu:\n%s\n----\n", err)
		}
	}
}
