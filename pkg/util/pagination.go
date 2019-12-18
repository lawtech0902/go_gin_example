package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_example/pkg/setting"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	/*
		获取分页页码
	*/
	res := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		res = (page - 1) * setting.AppSetting.PageSize
	}
	
	return res
}
