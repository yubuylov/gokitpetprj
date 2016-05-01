package backend

import (
	app_logs "github.com/yubuylov/gokitpetprj/public-api/loggers"
	"time"
)

func loggingMiddleware(logger app_logs.AppLogs) ServiceMW {
	return func(next Service) Service {
		return logmw{logger, next}
	}
}

type logmw struct {
	logs app_logs.AppLogs
	Service
}

func (mw logmw) CreateNodeEntity(nid int64, uid int64) (success bool, err error) {
	defer func(begin time.Time) {
		_ = mw.logs.Access.Log(
			"method", "CreateNodeEntity",
			"nid", nid,
			"uid", nid,
			"success", success,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	success, err = mw.Service.CreateNodeEntity(nid, uid)
	return
}