package ajusaude

import (
	"encoding/json"
	"fmt"
	"goajusaude/pkg/shared"
)

type AjuSaudeClient interface {
	SearchProtocolByDocument(document string) (*SearchResponse, error)
}

type ajuSaudeClient struct {
	client shared.HttpClient
}

func New(client shared.HttpClient) AjuSaudeClient {
	return ajuSaudeClient{client: client}
}

type SearchResponse struct {
	Protocols   []Protocol `json:"lista"`
	RequestDate string     `json:"dataHora"`
}

func (client ajuSaudeClient) SearchProtocolByDocument(document string) (*SearchResponse, error) {
	url := fmt.Sprintf("/listaesperapublica?tipo=TODAS&situacao=0&documento=%s", document)
	responseBytes, err := client.client.Get(url)
	if err != nil {
		return nil, err
	}

	var response SearchResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
