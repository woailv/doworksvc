package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"doworksvc/conf"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

var mu sync.Mutex

func DB() *xorm.Engine {
	mu.Lock()
	defer mu.Unlock()
	if engine == nil {
		var err error
	START:
		engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DBUser, conf.DBPwd, "localhost:3306", conf.DBName))
		if err == nil && engine.Ping() == nil {
			log.Println("已创建数据库,不用再次创建")
			engine.ShowSQL(true)
			return engine
		}
		if err != nil {
			log.Fatal("未知错误:", err.Error())
		}
		err = engine.Ping()
		if !strings.Contains(err.Error(), "Unknown database") {
			log.Fatalf("数据源初始化数据库连接失败,%s", err)
		}
		log.Println("开始新建数据库:", conf.DBName)
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/mysql", conf.DBUser, conf.DBPwd, "localhost:3306"))
		if err != nil {
			log.Fatalf("链接数据库失败,%s", err)
		}
		defer db.Close()
		stmt, err := db.Prepare(fmt.Sprintf("CREATE SCHEMA %s DEFAULT CHARACTER SET utf8mb4 ;", conf.DBName))
		if err != nil {
			log.Fatalf("创建数据库失败,error:%s", err.Error())
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("新建数据库成功", conf.DBName)
		goto START
	}
	return engine
}
