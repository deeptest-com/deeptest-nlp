package consts

type AnalyzerType string

const (
	Rasa    AnalyzerType = "rasa"
	Pattern AnalyzerType = "pattern"
)

func (e AnalyzerType) ToString() string {
	return string(e)
}
