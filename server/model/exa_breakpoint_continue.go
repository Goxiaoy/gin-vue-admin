package model

import (
	"gin-vue-admin/global"
	sg "github.com/goxiaoy/go-saas/gorm"
)

// file struct, 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
	sg.MultiTenancy
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
	sg.MultiTenancy
}
