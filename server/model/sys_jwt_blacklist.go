package model

import (
	"gin-vue-admin/global"
	sg "github.com/goxiaoy/go-saas/gorm"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
	sg.HasTenant
}
