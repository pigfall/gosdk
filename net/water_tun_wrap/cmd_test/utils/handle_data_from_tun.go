package utils

import (
	"context"

	"github.com/pigfall/gosdk/log"
)

func HandleDataFromTun(ctx context.Context, logger log.LoggerLite, msgReadFromTun chan MsgReadFromTun, msgWillWriteToConn chan []byte) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("HandlDataFromTun context done")
		case msg := <-msgReadFromTun:
			logger.Info("Read one msg  from Tun Channel")
			msgWillWriteToConn <- msg
		}
	}
}
