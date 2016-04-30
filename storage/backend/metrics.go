package backend

import (
	app_mertics "github.com/yubuylov/gokitpetprj/storage/metrics"
	"time"
)

func metricsMiddleware(metrics app_mertics.AppMetrics) ServiceMW {
	return func(next Service) Service {
		return metricsMW{metrics, next}
	}
}

type metricsMW struct {
	metrics app_mertics.AppMetrics
	Service
}

func (mw metricsMW) GetNodeEntity(nid NodeId, id EntityId) (m Entity, err error) {
	defer func(begin time.Time) {
		mw.metrics.Access.GetNodeEntity.Add(1)
		mw.metrics.Timers.GetNodeEntity.CatchOverTime(time.Since(begin), time.Second)
	}(time.Now())
	m, err = mw.Service.GetNodeEntity(nid, id)
	return
}

func (mw metricsMW) GetNodeEntities(nid NodeId) (list []Entity, err error) {
	defer func(begin time.Time) {
		mw.metrics.Access.GetNodeEntities.Add(1)
		mw.metrics.Timers.GetNodeEntities.CatchOverTime(time.Since(begin), time.Second)
	}(time.Now())
	list, err = mw.Service.GetNodeEntities(nid)
	return
}

func (mw metricsMW) GetNodeEntitiesCount(nid NodeId) (cnt int64, err error) {
	defer func(begin time.Time) {
		mw.metrics.Access.GetNodeEntitiesCount.Add(1)
		mw.metrics.Timers.GetNodeEntitiesCount.CatchOverTime(time.Since(begin), time.Second)
	}(time.Now())
	cnt, err = mw.Service.GetNodeEntitiesCount(nid)
	return
}