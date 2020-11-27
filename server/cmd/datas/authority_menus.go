package datas

import (
	"github.com/gookit/color"
	sg "github.com/goxiaoy/go-saas/gorm"
	"gorm.io/gorm"
	"os"
)

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
	TenantId sg.HasTenant `gorm:"index;primary_key"`
}



func InitSysAuthorityMenus(tenant sg.MultiTenancy,db *gorm.DB) {

	var authorityMenus = []SysAuthorityMenus{
		{"888", 1,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 2,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 3,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 4,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 5,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 6,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 7,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 8,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 9,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 10,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 11,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 12,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 13,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 14,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 15,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 16,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 17,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 18,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 19,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 20,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 21,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 22,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 23,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 24,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 25,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 26,sg.NewTenantId(tenant.TenantId.String)},
		{"888", 27,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 1,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 2,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 8,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 17,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 18,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 19,sg.NewTenantId(tenant.TenantId.String)},
		{"8881", 20,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 1,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 2,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 3,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 4,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 5,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 6,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 7,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 8,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 9,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 1,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 11,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 12,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 13,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 14,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 15,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 17,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 18,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 19,sg.NewTenantId(tenant.TenantId.String)},
		{"9528", 20,sg.NewTenantId(tenant.TenantId.String)},
	}

	if err := db.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ?", []string{"888", "8881", "9528"}).Find(&[]SysAuthorityMenus{}).RowsAffected == 53 {
			color.Danger.Println("sys_authority_menus表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> sys_authority_menus 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
