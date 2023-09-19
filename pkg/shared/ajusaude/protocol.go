package ajusaude

import "encoding/json"

type Protocol struct {
	ID               json.Number       `json:"protocolo"`
	Position         int               `json:"posicao"`
	CreatedAt        string            `json:"inclusao"`
	NameAbbreviation string            `json:"nome"`
	BirthDate        string            `json:"dataNascimento"`
	Age              int               `json:"idade"`
	Priority         ProtocolPriority  `json:"prioridade"`
	Status           ProtocolStatus    `json:"situacao"`
	Condition        ProtocolCondition `json:"condicao"`
	City             string            `json:"municipioEndereco"`
	Code             int               `json:"codigo"`
	Description      string            `json:"descricao"`
}
