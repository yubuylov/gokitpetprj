package metrics

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"
	app_config "github.com/yubuylov/gokitpetprj/public-api/config"
	"time"
)

type AppMetrics struct {
	Access AccessMetrics
	Timers TimeMetrics
}

type AccessMetrics struct {
	CreateNodeEntity metrics.Counter
}

type TimeMetrics struct {
	OverTimeCounter  metrics.Counter
	CreateNodeEntity MethodTimeMetric
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
			CreateNodeEntity: expvar.NewCounter("access_CreateNodeEntity"),
		},
		Timers: TimeMetrics{
			CreateNodeEntity: MethodTimeMetric{
				metrics.NewTimeHistogram(time.Millisecond,
					expvar.NewHistogram("duration_ms_CreateNodeEntity", 0, 10000, 3, quantiles...), ),
				expvar.NewCounter("overtime_CreateNodeEntity"),
			},
		},
	}

	return appMetrics
}