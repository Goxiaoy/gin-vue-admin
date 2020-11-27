package service

import (
	"context"
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils/upload"
	"mime/multipart"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func Upload(ctx context.Context,file model.ExaFileUploadAndDownload) error {
	return global.GVA_DB(ctx).Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.ExaFileUploadAndDownload

func FindFile(ctx context.Context,id uint) (error, model.ExaFileUploadAndDownload) {
	var file model.ExaFileUploadAndDownload
	err := global.GVA_DB(ctx).Where("id = ?", id).First(&file).Error
	return err, file
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func DeleteFile(ctx context.Context,file model.ExaFileUploadAndDownload) (err error) {
	var fileFromDb model.ExaFileUploadAndDownload
	err, fileFromDb = FindFile(ctx,file.ID)
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil{
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB(ctx).Where("id = ?", file.ID).Unscoped().Delete(file).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetFileRecordInfoList(ctx context.Context,info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB(ctx)
	var fileLists []model.ExaFileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func UploadFile(ctx context.Context,header *multipart.FileHeader, noSave string) (err error, file model.ExaFileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.ExaFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return Upload(ctx,f), f
	}
	return
}
