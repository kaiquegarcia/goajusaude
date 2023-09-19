package ajusaude

type ProtocolStatus string

const (
	STATUS_PENDING  ProtocolStatus = "PENDENTE"
	STATUS_REVIEWED ProtocolStatus = "BAIXADO"
	STATUS_REJECTED ProtocolStatus = "NEGADO"
	STATUS_CANCELED ProtocolStatus = "CANCELADO"
)
