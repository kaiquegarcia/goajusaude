package internal

import (
	"fmt"
	"goajusaude/pkg/shared/ajusaude"
)

func CheckProtocol(ajuSaudeClient ajusaude.AjuSaudeClient, document string) error {
	response, err := ajuSaudeClient.SearchProtocolByDocument(document)
	if err != nil {
		return err
	}

	fmt.Printf("%d protocolo(s) encontrado(s):\n", len(response.Protocols))
	fmt.Println("| # | Descrição + Nome | Situação/Condição |")
	fmt.Println("--------------------------------------------")
	for index := 0; index < len(response.Protocols); index++ {
		protocol := response.Protocols[index]
		fmt.Printf("| %d | %s para %s | %s/%s |\n", (index + 1), protocol.Description, protocol.NameAbbreviation, protocol.Status, protocol.Condition)
	}
	fmt.Println("--------------------------------------------")
	return nil
}
