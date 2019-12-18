module github.com/lawtech0902/go_gin_example

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.0
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.6 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3
	github.com/tealeg/xlsx v1.0.5
	github.com/unknwon/com v1.0.1
	golang.org/x/image v0.0.0-20191214001246-9130b4cfad52 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.0.0-20191210023423-ac6580df4449 // indirect
	golang.org/x/tools v0.0.0-20191212051200-825cb0626375 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.2 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	github.com/lawtech0902/go_gin_example/conf => ./go_gin_example/conf
	github.com/lawtech0902/go_gin_example/docs => ./go_gin_example/docs
	github.com/lawtech0902/go_gin_example/middleware => ./go_gin_example/middleware
	github.com/lawtech0902/go_gin_example/models => ./go_gin_example/models
	github.com/lawtech0902/go_gin_example/pkg/app => ./go_gin_example/pkg/app
	github.com/lawtech0902/go_gin_example/pkg/e => ./go_gin_example/pkg/e
	github.com/lawtech0902/go_gin_example/pkg/export => ./go_gin_example/pkg/export
	github.com/lawtech0902/go_gin_example/pkg/file => ./go_gin_example/pkg/file
	github.com/lawtech0902/go_gin_example/pkg/gredis => ./go_gin_example/pkg/gredis
	github.com/lawtech0902/go_gin_example/pkg/logging => ./go_gin_example/pkg/logging
	github.com/lawtech0902/go_gin_example/pkg/qrcode => ./go_gin_example/pkg/qrcode
	github.com/lawtech0902/go_gin_example/pkg/setting => ./go_gin_example/pkg/setting
	github.com/lawtech0902/go_gin_example/pkg/upload => ./go_gin_example/pkg/upload
	github.com/lawtech0902/go_gin_example/pkg/util => ./go_gin_example/pkg/util
	github.com/lawtech0902/go_gin_example/routers => ./go_gin_example/routers
	github.com/lawtech0902/go_gin_example/service => ./go_gin_example/service
)
