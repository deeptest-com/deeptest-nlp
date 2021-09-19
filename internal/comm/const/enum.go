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
	SeleniumStart SeleniumIntent = "启动浏览器"
	SeleniumStop  SeleniumIntent = "停止浏览器"

	Load      SeleniumIntent = "加载"
	GetSource SeleniumIntent = "获取源"

	CookieSet SeleniumIntent = "Cookie设置"
	CookieGet SeleniumIntent = "Cookie获取"

	VariableSetToExpression         SeleniumIntent = "设置变量为表达式"
	VariableSetSetToElementProperty SeleniumIntent = "设置变量为元素属性"

	WindowsConfirm SeleniumIntent = "确定"
	WindowsCancel  SeleniumIntent = "取消"

	IterationAlways   SeleniumIntent = "迭代永久"
	IterationTime     SeleniumIntent = "迭代次数"
	IterationData     SeleniumIntent = "迭代数据"
	IterationBreak    SeleniumIntent = "迭代退出"
	IterationContinue SeleniumIntent = "迭代继续"

	PrintLog SeleniumIntent = "打印日志"

	Click       SeleniumIntent = "点击"
	DoubleClick SeleniumIntent = "双击"
	RightClick  SeleniumIntent = "右击"
	MouseDown   SeleniumIntent = "按下"
	MouseUp     SeleniumIntent = "抬起"
	MouseOver   SeleniumIntent = "悬停"
	DragDrop    SeleniumIntent = "拖动"

	Input SeleniumIntent = "输入"
	Clear SeleniumIntent = "清除"

	Assert SeleniumIntent = "断言"
	If     SeleniumIntent = "判断"

	Forward SeleniumIntent = "前进"
	Back    SeleniumIntent = "后退"
	Refresh SeleniumIntent = "刷新"

	Wait SeleniumIntent = "等待"
)

func (e SeleniumIntent) ToString() string {
	return string(e)
}

type NluPlaceholder string

const ()

func (e NluPlaceholder) ToString() string {
	return string(e)
}

type NluSynonym string

const ()

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
