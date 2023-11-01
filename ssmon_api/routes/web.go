package routes

import (
	"ssmon-api/config"
	"ssmon-api/controllers/ssmon"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	var app *gin.Engine
	if config.IS_TEST == false {
		app = gin.New()
		gin.SetMode(gin.ReleaseMode)
	} else {
		app = gin.Default()
	}
	// app.MaxMultipartMemory = 64 << 20 // 64 MB
	app.Use(cors.Default())
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	// SSMON
	app_ssmon := app.Group("/ssmon")
	app_ssmon.POST("/request", ssmon.Request)	// SP 요청 처리 (1건)
	app_ssmon.POST("/request_array", ssmon.Request_array) // SP 요청 처리 (여러건)
	app_ssmon.POST("/login", ssmon.Login)  		// 로그인 처리
	app_ssmon.POST("/pass", ssmon.Pass)			// SP 요청 처리: PASS : 지정된 SP에 대해서는 토큰 체크하지 않고 통과시킨다.
	app_ssmon.POST("/send", ssmon.Send)			// SP 요청 처리: SEND : 지정된 SP만 실행 (토큰체크는 한다)

	return app
}
