package enums

const (
	AcknowledgedStatus = iota + 1
	InProgressStatus
	PendingStatus
	EscalatedStatus
	EscalatedProgressedStatus
	ResolvedStatus
)

const (
	CriticalPriority = iota + 1
	HighPriority
	MediumPriority
	LowPriority
)

const (
	CriticalSeverity = iota + 1
	MajorSeverity
	MinorSeverity
)

const (
	OperationChannel = iota + 1
	SalesChannel
	SupportChannel
	BillingChannel
	HRChannel
	FinanceChannel
)

const (
	IncidentType = iota + 1
	ComplainType
	RequestType
)

const (
	Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	Length  = 4
)
