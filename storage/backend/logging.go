package backend

import (
	app_logs "github.com/yubuylov/gokitpetprj/storage/loggers"
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

func (mw logmw) GetNodeEntity(nid NodeId, id EntityId) (m Entity, err error) {
	defer func(begin time.Time) {
		_ = mw.logs.Access.Log(
			"method", "GetNodeEntity",
			"id", id,
			"Model", m.toString(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	m, err = mw.Service.GetNodeEntity(nid, id)
	return
}

func (mw logmw) GetNodeEntities(nid NodeId) (list []Entity, err error) {
	defer func(begin time.Time) {
		_ = mw.logs.Access.Log(
			"method", "GetNodeEntities",
			"id", nid,
			"result_len", len(list),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	list, err = mw.Service.GetNodeEntities(nid)
	return
}

func (mw logmw) GetNodeEntitiesCount(nid NodeId) (cnt int64, err error) {
	defer func(begin time.Time) {
		_ = mw.logs.Access.Log(
			"method", "GetEntitiesCount",
			"id", nid,
			"cnt", cnt,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	cnt, err = mw.Service.GetNodeEntitiesCount(nid)
	return
}