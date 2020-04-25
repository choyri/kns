package store

import (
	"fmt"
	"github.com/choyri/kns/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type MySQL struct {
	*gorm.DB
	DBName string
}

var (
	mysql     MySQL
	mysqlOnce sync.Once
)

func InitMySQL() error {
	var err error

	mysqlOnce.Do(func() {
		mysql, err = newMySQL(config.GetMySQL())
	})

	return err
}

func GetMySQL() *MySQL {
	err := InitMySQL()
	if err != nil {
		panic(err)
	}

	return &mysql
}

func newMySQL(cfg config.MySQL) (MySQL, error) {
	var (
		err error
		ret MySQL
		db  *gorm.DB
	)

	db, err = gorm.Open("mysql", cfg.GetDSN())
	if err != nil {
		return ret, fmt.Errorf("gorm Open 失败：%w", err)
	}

	db.LogMode(cfg.EnableLog)
	db.DB().SetMaxIdleConns(cfg.MaxIdle)
	db.DB().SetMaxOpenConns(cfg.MaxOpen)
	db.DB().SetConnMaxLifetime(time.Hour)

	ret = MySQL{
		DB:     db,
		DBName: cfg.Database,
	}

	return ret, nil
}
