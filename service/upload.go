package service

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"market_monitor/extend/utils"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UploadService struct{}

// GetImgPath 获取图片相对目录
func (us *UploadService) GetImgPath() string {
	// todo conf
	return "upload/img/"
}

// GetImgFullPath 获取图片完整目录
func (us *UploadService) GetImgFullPath() string {
	// todo conf
	return "public/" + "upload/img/"
}

// GetImgName 获取图片名称
func (us *UploadService) GetImgName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.MakeSha1(fileName)
	return fileName + ext
}

// GetImgFullUrl 获取图片完整URL
func (us *UploadService) GetImgFullUrl(name string) string {
	// todo conf
	return "http://localhost:8000/" + "upload/img/" + name
}

// CheckImgExt 检查图片后缀是否满足要求
func (us *UploadService) CheckImgExt(fileName string) bool {
	ext := path.Ext(fileName)
	// todo conf
	for _, allowExt := range []string{".jpg", ".jpeg", ".png"} {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImgSize 检查图片大小是否超出
func (us *UploadService) CheckImgSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}
	// todo conf
	// 单位转换 bytes to Megabyte
	var converRatio float64 = 1024 * 1024
	fileSize := float64(len(content)) / converRatio
	// 文件大小不得超出5M
	return fileSize <= float64(5)
}

// CheckImgPath 检测图片路径是否创建及权限是否满足
func (us *UploadService) CheckImgPath(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	isExist, err := utils.IsExist(dir + "/" + path)
	if err != nil {
		return fmt.Errorf("utils.IsExist err: %v", err)
	}
	if isExist == false {
		// 若路径不存在，则创建
		err := os.MkdirAll(dir+"/"+path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("os.MkdirAll err: %v", err)
		}
	}
	isPerm := utils.IsPerm(path)
	if isPerm {
		return fmt.Errorf("utils.IsPerm Permission denied src: %s", path)
	}
	return nil
}
