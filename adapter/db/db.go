package db

import (
	"bass/config"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func New(conf *config.Conf) (*sql.DB, error) {
	cfg := &mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", conf.Db.Host, conf.Db.Port),
		DBName:               conf.Db.DbName,
		User:                 conf.Db.Username,
		Passwd:               conf.Db.Password,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return sql.Open("mysql", cfg.FormatDSN())

}
