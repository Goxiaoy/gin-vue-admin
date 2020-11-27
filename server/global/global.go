package global

import (
	"context"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
	sg "github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/go-saas/management/domain"
	"go.uber.org/zap"

	"gin-vue-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
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
	GVA_TENATN_REPO domain.TenantRepo

)

func GVA_DB(ctx context.Context) *gorm.DB {
	//just get default db
	return GVA_DB_PROVIDER.Get(ctx,data.Default)
}