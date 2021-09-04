package _const

import "os"

const (
	PthSep = string(os.PathSeparator)

	UserTokenExpireTime = 365 * 24 * 60 * 60 * 1000

	RpcPort = 8086

	UploadDir = "uploads"
	PageSize  = 10
)
