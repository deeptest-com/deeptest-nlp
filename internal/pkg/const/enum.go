package _const

type ResultCode int

const (
	ResultSuccess ResultCode = 1
	ResultFail    ResultCode = 0
)

func (c ResultCode) Int() int {
	return int(c)
}

type BuildType string

const (
	AppiumTest   BuildType = "appium_test"
	SeleniumTest BuildType = "selenium_test"
	UnitTest     BuildType = "unit_test"
)

func (e BuildType) ToString() string {
	return string(e)
}

type ValidMethod string

const (
	ValidProjectPath ValidMethod = "validProjectPath"
	ValidDictCode    ValidMethod = "validDictCode"
)

func (e ValidMethod) ToString() string {
	return string(e)
}
