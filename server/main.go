package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()          // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库

	initialize.InitMultiTenancyDb()
	initialize.MigrateTenantManagementAndHostTable()
	// 程序结束前关闭数据库链接
	global.GVA_DB_CLEAN()

	core.RunWindowsServer()
}
