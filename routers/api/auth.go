package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_example/pkg/app"
	"github.com/lawtech0902/go_gin_example/pkg/e"
	"github.com/lawtech0902/go_gin_example/pkg/util"
	"github.com/lawtech0902/go_gin_example/service/auth_service"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary JWT身份认证
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	
	username := c.Query("username")
	password := c.Query("password")
	
	a := auth{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)
	
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	
	authService := auth_service.Auth{Username: username, Password: password}
	exists, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	
	if !exists {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	
	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
