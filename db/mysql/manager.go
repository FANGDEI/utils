package mysql

import (
	"database/sql"
	"fmt"
	"utils/config"
	"utils/logs"

	_ "github.com/go-sql-driver/mysql"
)

var M *Manager

func init() {
	var err error
	M, err = New()
	if err != nil {
		logs.M.FatallnError(err)
	}
}

type Manager struct {
	handler *sql.DB
}

func New() (*Manager, error) {
	handler, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
			config.M.Db.Username,
			config.M.Db.Password,
			config.M.Db.Addr,
			config.M.Db.DBName,
		),
	)
	if err != nil {
		return nil, err
	}
	m := &Manager{
		handler: handler,
	}
	return m, m.handler.Ping()
}
