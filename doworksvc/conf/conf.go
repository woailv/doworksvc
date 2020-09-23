package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	DBPwd  string
	DBUser string
	DBName string
)

var (
	WYEmailUser     string
	WYEmailPassword string
)

func init() {
	file, err := os.Open("./conf/conf.yaml")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件异常:%s", err))
	}
	defer file.Close()
	m := map[string]string{}
	err = yaml.NewDecoder(file).Decode(&m)
	if err != nil {
		panic(fmt.Sprintf("解析配置文件异常:%s", err))
	}
	DBPwd = m["db_pwd"]
	DBUser = m["db_user"]
	DBName = m["db_name"]
	WYEmailUser = m["wy_email_user"]
	WYEmailPassword = m["wy_email_password"]
}
