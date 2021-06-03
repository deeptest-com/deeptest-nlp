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
	Create         NluAction = "create"
	Compile        NluAction = "compile"
	StartTraining  NluAction = "start_training"
	EndTraining    NluAction = "end_training"
	CancelTraining NluAction = "cancel_training"
	LaunchService  NluAction = "launch_service"
	StartService   NluAction = "start_service"
	StopService    NluAction = "stop_service"
)
