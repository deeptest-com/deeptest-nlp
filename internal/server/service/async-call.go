package service

import (
	"context"
	_logUtils "github.com/utlai/utl/internal/pkg/libs/log"
	serverConst "github.com/utlai/utl/internal/server/utils/const"
	"time"
)

func AsyncCall() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*serverConst.TrainingTimeout))
	defer cancel()
	go func(ctx context.Context) {
		// 发送HTTP请求
	}(ctx)

	select {
	case <-ctx.Done():
		_logUtils.Infof("---async call completed---")
		return
	case <-time.After(time.Duration(time.Millisecond * serverConst.TrainingTimeout)):
		_logUtils.Infof("---async call timeout---")
		return
	}
}
