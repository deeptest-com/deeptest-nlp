package consts

type AnalyzerType string

const (
	Rasa    AnalyzerType = "rasa"
	Pattern AnalyzerType = "pattern"
)

func (e AnalyzerType) ToString() string {
	return string(e)
}

type AgentStatus string

const (
	AgentBusy  AgentStatus = "busy"
	AgentReady AgentStatus = "ready"
)

func (e AgentStatus) ToString() string {
	return string(e)
}

type SeleniumIntent string

const (
	LogPrint SeleniumIntent = "打印日志"
	TimeWait SeleniumIntent = "等待"

	BrowserOpen  SeleniumIntent = "打开浏览器"
	BrowserClose SeleniumIntent = "关闭浏览器"

	NavForward SeleniumIntent = "前进"
	NavBack    SeleniumIntent = "后退"
	NavRefresh SeleniumIntent = "刷新"

	PageLoad      SeleniumIntent = "加载"
	PageGetSource SeleniumIntent = "获取源"

	WindowsConfirm SeleniumIntent = "确定"
	WindowsCancel  SeleniumIntent = "取消"

	CookieSet SeleniumIntent = "Cookie设置"
	CookieGet SeleniumIntent = "Cookie获取"

	VariableSetToExpression         SeleniumIntent = "设置变量为表达式"
	VariableSetSetToElementProperty SeleniumIntent = "设置变量为元素属性"

	IterationAlways   SeleniumIntent = "迭代永久"
	IterationTime     SeleniumIntent = "迭代次数"
	IterationData     SeleniumIntent = "迭代数据"
	IterationBreak    SeleniumIntent = "迭代退出"
	IterationContinue SeleniumIntent = "迭代继续"

	ElementClick       SeleniumIntent = "点击"
	ElementDoubleClick SeleniumIntent = "双击"
	ElementRightClick  SeleniumIntent = "右击"
	ElementMouseDown   SeleniumIntent = "按下"
	ElementMouseUp     SeleniumIntent = "抬起"
	ElementMouseOver   SeleniumIntent = "悬停"
	ElementDragDrop    SeleniumIntent = "拖动"

	ElementInput SeleniumIntent = "输入"
	ElementClear SeleniumIntent = "清除"

	LogicAssert SeleniumIntent = "断言"
	LogicIf     SeleniumIntent = "判断"
)

func (e SeleniumIntent) ToString() string {
	return string(e)
}

type NluPlaceholder string

const (
	Variable    NluPlaceholder = "variable"
	Value       NluPlaceholder = "value"
	Version     NluPlaceholder = "version"
	Expression  NluPlaceholder = "expression"
	Expression2 NluPlaceholder = "expression2"
	Data        NluPlaceholder = "data"
	Element     NluPlaceholder = "element"
	Property    NluPlaceholder = "property"
)

func (e NluPlaceholder) ToString() string {
	return string(e)
}

type NluSynonym string

const (
	Assert   NluSynonym = "assert"
	Print    NluSynonym = "print"
	Run      NluSynonym = "run"
	Script   NluSynonym = "script"
	If       NluSynonym = "if"
	Iterator NluSynonym = "iterator"
	Exit     NluSynonym = "exit"
	Continue NluSynonym = "continue"
	Wait     NluSynonym = "wait"
	Get      NluSynonym = "get"
	Set      NluSynonym = "set"
	Launch   NluSynonym = "launch"
	Stop     NluSynonym = "stop"
)

func (e NluSynonym) ToString() string {
	return string(e)
}

type NluLookUp string

const ()

func (e NluLookUp) ToString() string {
	return string(e)
}

type NluRegex string

const ()

func (e NluRegex) ToString() string {
	return string(e)
}
