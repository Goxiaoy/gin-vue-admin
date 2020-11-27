package datas

import (
	"github.com/gookit/color"
	sg "github.com/goxiaoy/go-saas/gorm"
	"os"
	"time"

	"gin-vue-admin/model"
	"gorm.io/gorm"
)



func InitSysAuthority(tenant sg.MultiTenancy, db *gorm.DB) {
	var authorities = []model.SysAuthority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: "888"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: "测试角色", ParentId: "0"},
	}
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("sys_authorities表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_authorities 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
