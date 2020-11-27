package global

import (
	"github.com/goxiaoy/go-saas/common"
	sg "github.com/goxiaoy/go-saas/gorm"
	"go.uber.org/zap"

	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG    *zap.Logger

	//interface for store
	GVA_TENANT_STORE common.TenantStore
	//interface for provider
	GVA_DB_PROVIDER sg.DbProvider
	GVA_DB_CLEAN sg.DbClean

)
