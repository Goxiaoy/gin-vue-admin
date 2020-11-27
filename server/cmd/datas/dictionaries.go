package datas

import (
	"gin-vue-admin/global"
	"github.com/gookit/color"
	sg "github.com/goxiaoy/go-saas/gorm"
	"os"
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)

func InitSysDictionary(tenant sg.MultiTenancy, db *gorm.DB) {
	var status = new(bool)
	*status = true
	dictionaries := []model.SysDictionary{
		{GVA_MODEL: global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{GVA_MODEL: global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{GVA_MODEL: global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{GVA_MODEL: global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{GVA_MODEL: global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{GVA_MODEL: global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("name IN ?", []string{"sex", "bool"}).Find(&[]model.SysDictionary{}).RowsAffected == 2 {
			color.Danger.Println("sys_dictionaries表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_dictionaries 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}