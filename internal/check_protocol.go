package internal

import (
	"fmt"
	"goajusaude/pkg/shared/ajusaude"
)

const PaginationLimit = 20

func CheckProtocol(ajuSaudeClient ajusaude.AjuSaudeClient, document string) error {
	page := 0
	protocols := make([]ajusaude.Protocol, 0)
	for {
		page++
		response, err := ajuSaudeClient.SearchProtocolByDocument(document, page, PaginationLimit)
		if err != nil {
			return err
		}

		protocols = append(protocols, response.Protocols...)
		if page*PaginationLimit >= response.TotalCount {
			break
		}
	}

	if len(protocols) == 0 {
		fmt.Printf("Nenhum protocolo encontrado ao buscar por %s\n", document)
		return nil
	}

	fmt.Printf("%d protocolo(s) encontrado(s) ao buscar por %s:\n", len(protocols), document)
	fmt.Println("| # | Descrição + Nome | Situação/Condição |")
	fmt.Println("--------------------------------------------")
	for index, protocol := range protocols {
		fmt.Printf("| %d | %s para %s | %s/%s |\n", (index + 1), protocol.Description, protocol.NameAbbreviation, protocol.Status, protocol.Condition)
	}
	fmt.Println("--------------------------------------------")
	return nil
}
