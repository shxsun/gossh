package models

import (
	"errors"
	"log"

	//_ "github.com/lib/pq" //当某时间字段表现为0001-01-01 07:36:42+07:36:42形式的时候 会读不出数据
	//_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Engine *xorm.Engine
)

const (
	DbName = "./sqlite.db"
	dbtype = "sqlite"
)

func init() {
	_, err := SetEngine()
	if err != nil {
		log.Fatal(err)
	}
}

func XConDb() (*xorm.Engine, error) {
	switch {
	case dbtype == "sqlite":
		return xorm.NewEngine("sqlite3", DbName)

	case dbtype == "mysql":
		return xorm.NewEngine("mysql", "user=mysql password=jn!@#9^&* dbname=mysql")

	case dbtype == "pgsql":
		return xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")
	}
	return nil, errors.New("dbtype not specified")
}

func SetEngine() (*xorm.Engine, error) {
	var err error
	Engine, err = XConDb()
	//Engine.Mapper = xorm.SameMapper{}
	//Engine.SetMaxConns(5)
	//Engine.ShowSQL = true

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)

	return Engine, err
}

func CreateDb() {
	err := Engine.Sync(new(Record))
	if err != nil {
		log.Fatal(err)
	}
}
