/**
 * 获取数据库连接
 */
package datasource

import (
	"fmt"
	"free/conf"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

//数据库连接

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	c := conf.GetDBConf()
	// path := c.User + ":" + c.Pwd + "@", c.Host + ":" + c.Port + "/" + c.DbName + "?charset=utf8&parseTime=True&loc=Asia/Shanghai"
	path := strings.Join([]string{c.User, ":", c.Pwd, "@(", c.Host, ":", c.Port, ")/", c.DbName, "?charset=utf8&parseTime=true"}, "")
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.MaxIdleConns")) //最大打开的连接数
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.MaxOpenConns")) //设置最大闲置个数
	// db.SingularTable(true)                                      //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(viper.GetBool("app.debug"))
}
