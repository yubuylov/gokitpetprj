package metrics

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"
	app_config "github.com/yubuylov/gokitpetprj/storage/config"
	"time"
)

type AppMetrics struct {
	Access AccessMetrics
	Timers TimeMetrics
}

type AccessMetrics struct {
	GetNodeEntity        metrics.Counter
	GetNodeEntities      metrics.Counter
	GetNodeEntitiesCount metrics.Counter
}

type TimeMetrics struct {
	OverTimeCounter      metrics.Counter
	GetNodeEntity        MethodTimeMetric
	GetNodeEntities      MethodTimeMetric
	GetNodeEntitiesCount MethodTimeMetric
}

type MethodTimeMetric struct {
	th       metrics.TimeHistogram
	overtime metrics.Counter
}

func (m MethodTimeMetric)CatchOverTime(dur time.Duration, max time.Duration) {
	if dur > max {
		m.overtime.Add(1)
	}
	m.th.Observe(dur)
}

func Load(cfg app_config.AppConfig) AppMetrics {
	var quantiles = []int{50, 90, 95, 99}
	appMetrics := AppMetrics{
		Access: AccessMetrics{
			GetNodeEntity: expvar.NewCounter("access_GetNodeEntity"),
			GetNodeEntities: expvar.NewCounter("access_GetNodeEntities"),
			GetNodeEntitiesCount: expvar.NewCounter("access_GetNodeEntitiesCount"),
		},
		Timers: TimeMetrics{
			GetNodeEntity: MethodTimeMetric{
				metrics.NewTimeHistogram(time.Millisecond,
					expvar.NewHistogram("duration_ms_GetEntity", 0, 10000, 3, quantiles...), ),
				expvar.NewCounter("overtime_GetEntity"),
			},
			GetNodeEntities: MethodTimeMetric{
				metrics.NewTimeHistogram(time.Millisecond,
					expvar.NewHistogram("duration_ms_GetNodeEntities", 0, 10000, 3, quantiles...), ),
				expvar.NewCounter("overtime_GetNodeEntities"),
			},
			GetNodeEntitiesCount: MethodTimeMetric{
				metrics.NewTimeHistogram(time.Millisecond,
					expvar.NewHistogram("duration_ms_GetNodeEntitiesCount", 0, 10000, 3, quantiles...), ),
				expvar.NewCounter("overtime_GetNodeEntitiesCount"),
			},
		},
	}

	return appMetrics
}