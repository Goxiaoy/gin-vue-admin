package datas

import (
	"gin-vue-admin/global"
	"github.com/gookit/color"
	"os"
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
	sg "github.com/goxiaoy/go-saas/gorm"
)

func InitSysDictionaryDetail(tenant sg.MultiTenancy,db *gorm.DB) {
	status := new(bool)
	*status = true
	DictionaryDetail := []model.SysDictionaryDetail{
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "smallint", 1, status, 1, 2,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumint", 2, status, 2, 2,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "int", 3, status, 3, 2,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "bigint", 4, status, 4, 2,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "date", 0, status, 0, 3,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "time", 1, status, 1, 3,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "year", 2, status, 2, 3,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "datetime", 3, status, 3, 3,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "timestamp", 5, status, 5, 3,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "float", 0, status, 0, 4,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "double", 1, status, 1, 4,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "decimal", 2, status, 2, 4,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "char", 0, status, 0, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "varchar", 1, status, 1, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyblob", 2, status, 2, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinytext", 3, status, 3, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "text", 4, status, 4, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "blob", 5, status, 5, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumblob", 6, status, 6, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumtext", 7, status, 7, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longblob", 8, status, 8, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longtext", 9, status, 9, 5,tenant},
		{global.GVA_MODEL{ CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyint", 0, status, 0, 6,tenant},
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("label IN ?", []string{"smallint", "tinyint"}).Find(&[]model.SysDictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("sys_dictionary_details表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&DictionaryDetail).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_dictionary_details 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
