package initialize

import (
	"context"
	"gin-vue-admin/global"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
	sg "github.com/goxiaoy/go-saas/gorm"
	sgm "github.com/goxiaoy/go-saas/management/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMultiTenancyDb()  {
	initDbProvider()
	initTenantStore()
}

func MigrateTenantManagementAndHostTable()  {
	hostDb := sgm.GetDb(context.Background(),global.GVA_DB_PROVIDER)
	err :=sgm.AutoMigrate(nil,hostDb)
	if err!=nil{
		panic(err)
	}
	MysqlTables(hostDb) // 初始化表
}

func initDbProvider()(*sg.DefaultDbProvider, sg.DbClean) {
	m := global.GVA_CONFIG.Mysql
	cfg :=sg.Config{
		Debug:        m.LogMode,
		Dialect: func(s string) gorm.Dialector {
			mysqlConfig := mysql.Config{
				DSN:                       s,   // DSN data source name
				DefaultStringSize:         191,   // string 类型字段的默认长度
				DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
				DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
				DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
				SkipInitializeWithVersion: false, // 根据版本自动配置
			}
			return mysql.New(mysqlConfig)
		},
		Cfg:          GormConfig(m.LogMode),
		MaxOpenConns: m.MaxOpenConns,
		MaxIdleConns: m.MaxOpenConns,
	}
	conn := make(data.ConnStrings,1)

	//default to host
	conn.SetDefault(m.Dsn)
	ct := common.ContextCurrentTenant{}
	mr := common.NewMultiTenancyConnStrResolver(ct, func() common.TenantStore {
		return global.GVA_TENANT_STORE
	},data.ConnStrOption{
		Conn: conn,
	})
	r ,close := sg.NewDefaultDbProvider(mr,cfg)
	global.GVA_DB_PROVIDER=r
	global.GVA_DB_CLEAN=close
	return r,close
}

func initTenantStore()  {
	tenantRepo := &sgm.GormTenantRepo{DbProvider: global.GVA_DB_PROVIDER}
	global.GVA_TENANT_STORE=sgm.NewGormTenantStore(tenantRepo)
}