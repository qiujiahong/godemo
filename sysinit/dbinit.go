package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "mbook/models"
)

//初始化操作
func dbinit(aliases ...string) {
	isDev := ("dev" == beego.AppConfig.String("runmode"))

	if len(aliases) > 0 {
		for _, alias := range aliases {
			registrDatabase(alias)
			if "w" == alias {
				orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registrDatabase("w")
		orm.RunSyncdb("default", false, isDev)
	}

}

func registrDatabase(alias string) {

	if (len(alias) <= 0) {
		return
	}

	dbAlias := alias //default

	if ("w" == alias || "default" == alias || len(alias) <= 0) {
		dbAlias = "default"
		alias = "w"
	}
	//数据名字
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	dbUsername := beego.AppConfig.String("db_" + alias + "_username")
	dbPassword := beego.AppConfig.String("db_" + alias + "_password")
	dbDatabase := beego.AppConfig.String("db_" + alias + "_database")
	//root:1234456@tcp(127.0.0.1:3306)/mbook?charset=utf8
	var db_url =  dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase+"?charset=utf8"
	fmt.Println("db_url",db_url)
	orm.RegisterDataBase(dbAlias, "mysql",db_url , 30)

	isDev := ("dev" == beego.AppConfig.String("runmode"))

	if "w" == alias {
		orm.RunSyncdb("default", false, isDev)
	}

	if isDev {
		orm.Debug = isDev
	}
}
