package backend

import (
	cfg "github.com/yubuylov/gokitpetprj/storage/config"
	logs "github.com/yubuylov/gokitpetprj/storage/loggers"
	"sync"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"fmt"
)

type storage struct {
	dbWrite *gorp.DbMap
	dbRead  *gorp.DbMap
	tblName string
	mtx     sync.RWMutex
	logger  *logs.AppLogs
}

// Persistence storage
type Storage interface {
	Get(id EntityId) (Entity, error)
}

func (s *storage) Get(id EntityId) (m Entity, err error) {
	var row Unsafe
	//query := fmt.Sprintf("select * from %s where ID=?", s.tblName)
	// err := s.dbRead.SelectOne(&row, query, id)
	row.Id = new(EntityId); *row.Id = id
	row.NodeID = new(NodeId); *row.NodeID = 1
	row.Payload = new(string); *row.Payload = "asdf payload"
	s.logErr(err)
	m.ApplyRawIO(row)
	return
}

func (s *storage) logErr(err error) {
	if err != nil {
		s.logger.Error.Log("mysql", err)
	}
}

func NewStorage(appCfg cfg.AppConfig, appLogs *logs.AppLogs) Storage {

	appLogs.Info.Log("mysql", "Init storage...")

	c := appCfg.Mysql

	dSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Pass, c.Host, c.Port, c.DBName)
	appLogs.Info.Log("mysql:dSN", dSN)
	db, err := sql.Open("mysql", dSN); if err != nil {
		appLogs.Info.Log("mysqlErr", err.Error())
	}

	// switch of table associations
	if false {
		// Entity is a type-safe struct for write
		dbWrite := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		dbWrite.AddTableWithName(Entity{}, appCfg.Mysql.TBLName).SetKeys(true, "id")

		// Unsafe used if read value can be nil
		dbRead := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		dbRead.AddTableWithName(Unsafe{}, appCfg.Mysql.TBLName).SetKeys(true, "id")
	}


	var s Storage
	{
		s = &storage{
			tblName: appCfg.Mysql.TBLName,
			//dbWrite: dbWrite,
			//dbRead: dbRead,
			logger: appLogs,
		}
	}

	return s
}