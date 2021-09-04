package domain

import consts "github.com/utlai/utl/internal/comm/const"

type InstructSelenium struct {
	Instruct string `json:"instruct"`

	Opt consts.SeleniumIntent `json:"opt"`
}
