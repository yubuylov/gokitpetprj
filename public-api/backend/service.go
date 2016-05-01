package backend

import (
	"sync"
	cfg "github.com/yubuylov/gokitpetprj/public-api/config"
	logs "github.com/yubuylov/gokitpetprj/public-api/loggers"
	metrics "github.com/yubuylov/gokitpetprj/public-api/metrics"

	"net/http"

	"strings"
	"net/url"
	"io/ioutil"
	"fmt"
)

type ServiceMW func(Service) Service
type Service interface {
	CreateNodeEntity(nid int64, uid int64) (bool, error)
}

type service struct {
	cfg    cfg.AppConfig
	logger logs.AppLogs
	mtx    sync.RWMutex
}

func (s *service) CreateNodeEntity(nid int64, uid int64) (success bool, err error) {

	resCh := make(chan []byte, 4)
	// retrieve some data from relation resources
	go callResource(resCh, s.cfg.Relations.Storage, "/api/v1/1/entities") // Sleep(10
	go callResource(resCh, s.cfg.Relations.Storage, "/api/v1/1/entities/count") // Sleep(5
	go callResource(resCh, s.cfg.Relations.Storage, "/api/v1/1/entities/1") // Sleep(2
	go callResource(resCh, s.cfg.Relations.Storage, "/api/v1/1/entities/2") // Sleep(2

	s.logger.Debug.Log("res1", string(<-resCh))
	s.logger.Debug.Log("res2", string(<-resCh))
	s.logger.Debug.Log("res3", string(<-resCh))
	s.logger.Debug.Log("res4", string(<-resCh))
	// process results

	success = true
	return
}

func InitService(appCfg cfg.AppConfig, appLogs logs.AppLogs, appMetrics metrics.AppMetrics) Service {
	var svc Service
	{
		svc = &service{
			logger:appLogs,
			cfg:appCfg,
		}
		svc = loggingMiddleware(appLogs)(svc)
		svc = metricsMiddleware(appMetrics)(svc)
	}
	return svc
}

func callResource(resCh chan []byte, instance string, path string) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance); if err != nil { return }
	u.Path = path
	resp, err := http.Get(u.String())
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("recieved:", path)
	resCh <- body
	return
}