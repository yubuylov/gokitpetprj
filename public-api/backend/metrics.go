package backend

import (
	app_mertics "github.com/yubuylov/gokitpetprj/public-api/metrics"
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

func (mw metricsMW) CreateNodeEntity(nid int64, uid int64) (success bool, err error) {
	defer func(begin time.Time) {
		mw.metrics.Access.CreateNodeEntity.Add(1)
		mw.metrics.Timers.CreateNodeEntity.CatchOverTime(time.Since(begin), time.Second)
	}(time.Now())
	success, err = mw.Service.CreateNodeEntity(nid, uid)
	return
}