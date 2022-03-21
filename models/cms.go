package models

import (
	"fmt"
	"gin-fast/conf"
	"gin-fast/pkg/util"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type FileModel struct {
	BaseModel
	Name string `gorm:"comment:'文件名';column:name;not null;" json:"name"`
	Path string `gorm:"comment:'路径';column:path;not null;" json:"path"`
}

func (file *FileModel) TableName() string {
	return "cms_files"
}

// DatePath 获取年月日文件夹
func (file *FileModel) DatePath() string {
	now := time.Now()
	year := now.Year()
	month := now.Month().String()
	day := now.Day()
	return fmt.Sprintf("%s/%s/%s/%s",
		conf.ProjectCfg.MediaFilePath,
		strconv.Itoa(year),
		month,
		strconv.Itoa(day))
}

// MkMediaDir 创建 年/月/日 文件夹
func (file *FileModel) MkMediaDir() (string, error) {
	dir := conf.ProjectCfg.MediaFilePath + time.Now().Format("2021/12/08")
	if !util.FileOrDirExists(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return dir, err
		} else {
			return dir, nil
		}
	} else {
		return dir, nil
	}
}

// BuildAbsoluteUri 构建全路径url
func (file *FileModel) BuildAbsoluteUri(ctx *gin.Context) {
	file.Path = util.BuildAbsoluteUri(ctx, file.Path)
}
