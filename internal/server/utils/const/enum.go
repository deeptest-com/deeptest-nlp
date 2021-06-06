package serverConst

type NluSlotType string

const (
	Synonym NluSlotType = "synonym"
	Lookup  NluSlotType = "lookup"
	Regex   NluSlotType = "regex"
	Text    NluSlotType = "text"
)

type NluAction string

const (
	Create  ServiceStatus = "create"
	Compile ServiceStatus = "compile"
)

type ServiceStatus string

const (
	StartService ServiceStatus = "start_service"
	StopService  ServiceStatus = "stop_service"
)

type TrainingStatus string

const (
	StartTraining  ServiceStatus = "start_training"
	EndTraining    ServiceStatus = "end_training"
	CancelTraining ServiceStatus = "cancel_training"
)
