package datas

import (
	"context"
	"gin-vue-admin/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"github.com/goxiaoy/go-saas/common"
	sg "github.com/goxiaoy/go-saas/gorm"
	"gorm.io/gorm"
	"os"
)

func InitMysqlData(ctx context.Context,db *gorm.DB) {
	tenantId := common.FromCurrentTenant(ctx).Id
	tenant:=sg.MultiTenancy{
		TenantId: sg.NewTenantId(tenantId),
	}
	InitSysApi(db)
	InitSysUser(tenant,db)
	InitExaCustomer(tenant,db)
	InitCasbinModel(db)
	InitSysAuthority(tenant,db)
	InitSysBaseMenus(db)
	InitAuthorityMenu(tenant,db)
	InitSysDictionary(tenant,db)
	InitSysAuthorityMenus(db)
	InitSysDataAuthorityId(db)
	InitSysDictionaryDetail(tenant,db)
	InitExaFileUploadAndDownload(tenant,db)
	InitWkProcess(db)
}

func InitMysqlTables(ctx context.Context,db *gorm.DB) {
	var err error
	if !db.Migrator().HasTable("casbin_rule") {
		err = db.Migrator().CreateTable(&gormadapter.CasbinRule{})
	}
	err = db.AutoMigrate(
		model.SysApi{},
		model.SysUser{},
		model.ExaFile{},
		model.ExaCustomer{},
		model.SysBaseMenu{},
		model.SysAuthority{},
		model.JwtBlacklist{},
		model.ExaFileChunk{},
		model.SysDictionary{},
		model.ExaSimpleUploader{},
		model.SysOperationRecord{},
		model.SysDictionaryDetail{},
		model.SysBaseMenuParameter{},
		model.ExaFileUploadAndDownload{},
		model.WorkflowProcess{},
		model.WorkflowNode{},
		model.WorkflowEdge{},
		model.WorkflowStartPoint{},
		model.WorkflowEndPoint{},
	)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化数据表失败,err: %v\n", err)
		os.Exit(0)
	}
	color.Info.Println("[Mysql]-->初始化数据表成功")
}
