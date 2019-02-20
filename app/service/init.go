package service

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"nway/nway_400/app/entity"
	"time"
)

var (
	o           orm.Ormer
	tablePrefix string      // 表前缀
	SmService   *smService  // 短信查询
	CdrService  *cdrService // 话单查询
)

func Init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	tablePrefix = beego.AppConfig.String("db.prefix")

	if dbPort == "" {
		dbPort = "5432"
	}
	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	if timezone != "" {
		local2, _ := time.LoadLocation(timezone)
		orm.DefaultTimeLoc = local2
	}
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dsn)

	orm.RegisterModelWithPrefix(tablePrefix,
		new(entity.Sm),
		new(entity.Cdr),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	// 初始化服务对象
	initService()
}

func initService() {
	SmService = &smService{}
	CdrService = &cdrService{}
}

// 返回真实表名
func tableName(name string) string {
	return tablePrefix + name
}

func debug(v ...interface{}) {
	beego.Debug(v...)
}

func concatenateError(err error, stderr string) error {
	if len(stderr) == 0 {
		return err
	}
	return fmt.Errorf("%v: %s", err, stderr)
}

func DBVersion() string {
	var lists []orm.ParamsList
	o.Raw("SELECT VERSION()").ValuesList(&lists)
	return lists[0][0].(string)
}
