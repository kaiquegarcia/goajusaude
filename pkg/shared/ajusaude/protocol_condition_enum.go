package ajusaude

type ProtocolCondition string

const (
	CONDITION_PENDING   ProtocolCondition = "Pendente"
	CONDITION_SCHEDULED ProtocolCondition = "Agendado"
	CONDITION_CANCELED  ProtocolCondition = "Cancelado"
	CONDITION_REJECTED  ProtocolCondition = "Negado"
)
