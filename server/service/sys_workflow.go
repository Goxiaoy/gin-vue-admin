package service

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Create
//@description: 创建工作流
//@param: wk model.SysWorkflow
//@return: error

func Create(ctx context.Context,wk model.SysWorkflow) error {
	err := global.GVA_DB(ctx).Create(&wk).Error
	return err
}
