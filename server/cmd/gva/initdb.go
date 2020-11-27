/*
Copyright © 2020 SliverHorn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gva

import (
	"context"
	"gin-vue-admin/cmd/datas"
	"gin-vue-admin/core"
	"gin-vue-admin/initialize"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/management/domain"
	"github.com/goxiaoy/go-saas/management/gorm"

	"github.com/gookit/color"

	_ "gin-vue-admin/core"
	"gin-vue-admin/global"

	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "gin-vue-admin初始化数据",
	Long: `gin-vue-admin初始化数据适配数据库情况: 
1. mysql完美适配,
2. postgresql不能保证完美适配,
3. sqlite未适配,
4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		global.GVA_VP = core.Viper(path)
		global.GVA_LOG = core.Zap()           // 初始化zap日志库

		initialize.InitMultiTenancyDb()
		initialize.GetAndMigrateHostDb()
		ctx := context.Background()
		tr := global.GVA_TENATN_REPO
		its := []domain.Tenant{
			{
				ID:          "6045b1e6-b25d-4f97-b3f5-a6289513aae2",
				Name:        "t1",
				DisplayName: "Test1",
			},
			{
				ID:          "34bd2809-c7e2-43f8-b406-9073a23f256d",
				Name:        "t2",
				DisplayName: "Test2",
			},
		}
		for _, it := range its {
			t,err:= tr.FindByIdOrName(ctx,it.ID)
			if err!=nil{
				panic(err)
			}
			if t==nil{
				//create
				err :=tr.Create(ctx,it)
				if err!=nil{
					panic(err)
				}
			}
		}

		_,tenantDbAll,err := tr.GetPaged(ctx,common.Pagination{
			Offset: 0,
			Limit:  int(^uint(0) >> 1),
		})
		if err!=nil{
			panic(err)
		}
		all := []*domain.Tenant{
			//host side
			{ID: "",Name: ""},
		}
		all = append(all,tenantDbAll...)
		for _, tenant := range all {
			color.Info.Printf("[Migration]--> 正在迁移租户: %v %v\n", tenant.ID,tenant.Name)
			newCtx := common.NewCurrentTenant(context.Background(),tenant.ID,tenant.Name)
			db := gorm.GetDb(newCtx,global.GVA_DB_PROVIDER)
			switch global.GVA_CONFIG.System.DbType {
			case "mysql":
				datas.InitMysqlTables(newCtx,db)
				datas.InitMysqlData(newCtx,db)
			case "postgresql":
				color.Info.Println("postgresql功能开发中")
			case "sqlite":
				color.Info.Println("sqlite功能开发中")
			case "sqlserver":
				color.Info.Println("sqlserver功能开发中")
			default:
				datas.InitMysqlTables(newCtx,db)
				datas.InitMysqlData(newCtx,db)
			}
		}

		frame, _ := cmd.Flags().GetString("frame")
		if frame == "gf" {
			color.Info.Println("gf功能开发中")
			return
		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gin", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
