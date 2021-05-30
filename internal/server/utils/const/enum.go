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
	StartTraining  NluAction = "startTraining"
	EndTraining    NluAction = "endTraining"
	CancelTraining NluAction = "cancelTraining"
	LaunchService  NluAction = "launchService"
	StartService   NluAction = "startService"
	StopService    NluAction = "stopService"
)
