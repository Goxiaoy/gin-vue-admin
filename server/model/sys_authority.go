package model

import (
	sg "github.com/goxiaoy/go-saas/gorm"
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"not null;primary_key;comment:角色ID;size:90"`
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	TenantId sg.HasTenant `gorm:"index;primary_key"`
}
