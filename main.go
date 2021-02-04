package main

import (
	"go-gin-start/app/router"
	"go-gin-start/app/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	////////////////////////
	// App Bootstrap
	////////////////////////
	util.CfgBootstrap()
	// util.DbBootstrap()

	////////////////////////
	// App first run check
	////////////////////////
	util.FirstRunChecking()

	////////////////////////
	// Set Gin mode
	////////////////////////
	switch util.Cfg.Section("").Key("app_mode").String() {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	}

	////////////////////////
	// Timed tasks
	////////////////////////
	// if true {
	// 	// 2min
	// 	go func() {
	// 		timeTickerChan := time.Tick(time.Minute * 2)
	// 		for {
	// 			// 定时任务
	// 			// ...
	// 			<-timeTickerChan
	// 		}
	// 	}()
	// }

	////////////////////////
	// Add cors && Run
	////////////////////////
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("authorization, x-requested-with")
	corsConfig.AddExposeHeaders("X-Total-Count")
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	r = router.APIRouter(r)
	r.Run(":" + util.Cfg.Section("server").Key("http_port").String())

	////////////////////////
	// Defers
	////////////////////////
	defer util.DBC.Close()

}
