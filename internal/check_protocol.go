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

	if len(response.Protocols) == 0 {
		fmt.Printf("Nenhum protocolo encontrado ao buscar por %s\n", document)
		return nil
	}

	fmt.Printf("%d protocolo(s) encontrado(s) ao buscar por %s:\n", len(response.Protocols), document)
	fmt.Println("| # | Descrição + Nome | Situação/Condição |")
	fmt.Println("--------------------------------------------")
	for index := 0; index < len(response.Protocols); index++ {
		protocol := response.Protocols[index]
		fmt.Printf("| %d | %s para %s | %s/%s |\n", (index + 1), protocol.Description, protocol.NameAbbreviation, protocol.Status, protocol.Condition)
	}
	fmt.Println("--------------------------------------------")
	return nil
}
