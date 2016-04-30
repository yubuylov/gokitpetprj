package backend

import (
	"sync"
	logs "github.com/yubuylov/gokitpetprj/storage/loggers"
	metrics "github.com/yubuylov/gokitpetprj/storage/metrics"
	"errors"
	"time"
)

type ServiceMW func(Service) Service
type Service interface {
	GetNodeEntity(nid NodeId, id EntityId) (Entity, error)
	GetNodeEntities(nid NodeId) ([]Entity, error)
	GetNodeEntitiesCount(nid NodeId) (int64, error)
}

type service struct {
	logger  logs.AppLogs
	mtx     sync.RWMutex
	storage Storage
}

func (s *service) GetNodeEntity(nid NodeId, id EntityId) (Entity, error) {
	var m Entity
	m, err := s.storage.Get(id)
	if err != nil {
		return m, err
	}

	if m.NodeID != nid {
		return m, errors.New("Bad node relation..")
	}

	time.Sleep(1 * time.Millisecond)

	return m, nil
}

func (s *service) GetNodeEntities(nid NodeId) (list []Entity, err error) {
	// storage.find_all entities by NodeId
	time.Sleep(2 * time.Millisecond)
	list = []Entity{{1,1,"asdf payload"},{2,1,"asdf payload"}}
	return
}

func (s *service) GetNodeEntitiesCount(nid NodeId) (cnt int64, err error) {
	// storage.count entities by NodeId
	// use in-memory cache if u need
	time.Sleep(3 * time.Millisecond)
	cnt = 2
	return
}

func InitService(storage Storage, appLogs logs.AppLogs, appMetrics metrics.AppMetrics) Service {
	var svc Service
	{
		svc = &service{
			storage: storage,
			logger:appLogs,
		}
		svc = loggingMiddleware(appLogs)(svc)
		svc = metricsMiddleware(appMetrics)(svc)
	}
	return svc
}