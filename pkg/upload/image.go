package upload

import (
	"fmt"
	"github.com/lawtech0902/go_gin_example/pkg/file"
	"github.com/lawtech0902/go_gin_example/pkg/logging"
	"github.com/lawtech0902/go_gin_example/pkg/setting"
	"github.com/lawtech0902/go_gin_example/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// 获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	
	return fileName + ext
}

// 获取图片保存路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// 获取图片完整保存路径
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	
	return false
}

// 检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	
	return size <= setting.AppSetting.ImageMaxSize
}

// 检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	
	return nil
}
