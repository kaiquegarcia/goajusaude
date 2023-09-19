package main

import (
	"fmt"
	"goajusaude/internal"
	"goajusaude/pkg/shared"
	"goajusaude/pkg/shared/ajusaude"
)

func main() {
	ajuSaudeClient := ajusaude.New(
		shared.NewHttpClient("http://aracajusaude.voipy.com.br:8090/IDSPortalCidadaoWS/rest/portalcidadao"),
	)

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
