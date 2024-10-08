// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTPmFirmwareVersion = "t_pm_firmware_version"

// TPmFirmwareVersion mapped from table <t_pm_firmware_version>
type TPmFirmwareVersion struct {
	Id          int64          `gorm:"column:id;primaryKey" json:"id"`                // 唯一主键
	FirmwareId  int64          `gorm:"column:firmware_id;not null" json:"firmwareId"` // 固件主键编号
	Version     string         `gorm:"column:version" json:"version"`                 // 固件SDk版本号
	Desc        string         `gorm:"column:desc" json:"desc"`                       // 固件SDK版本描述
	Status      int32          `gorm:"column:status" json:"status"`                   // 状态（1上架 2未上架）
	FileName    string         `gorm:"column:file_name" json:"fileName"`              // 文件名称
	FilePath    string         `gorm:"column:file_path" json:"filePath"`              // 固件版本文件
	FileKey     string         `gorm:"column:file_key" json:"fileKey"`                // 固件版本MD5值
	FileSize    int32          `gorm:"column:file_size" json:"fileSize"`              // 固件版本包大小
	CreatedAt   time.Time      `gorm:"column:created_at" json:"createdAt"`            // 创建时间
	UpdatedBy   int64          `gorm:"column:updated_by" json:"updatedBy"`            // 修改人
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updatedAt"`            // 修改时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`            // 删除标识（0-正常 1-删除）
	ZipFileName string         `gorm:"column:zip_file_name" json:"zipFileName"`       // 固件版本资源文件名称
	ZipFilePath string         `gorm:"column:zip_file_path" json:"zipFilePath"`       // 固件版本资源文件
	ZipFileKey  string         `gorm:"column:zip_file_key" json:"zipFileKey"`         // 固件版本资源MD5值
	ZipFileSize int32          `gorm:"column:zip_file_size" json:"zipFileSize"`       // 固件版本资源包大小
	IsMust      int32          `gorm:"column:is_must" json:"isMust"`                  // 是否必须
}

// TableName TPmFirmwareVersion's table name
func (*TPmFirmwareVersion) TableName() string {
	return TableNameTPmFirmwareVersion
}
