package main

import (
	"market_monitor/schedule"
	"strconv"

	"market_monitor/config"
	"market_monitor/extend/logger"
	"market_monitor/extend/redis"
	"market_monitor/models"
	"market_monitor/router"
)

// @title Market Monitor API
// @version 1.0
// @description Market Monitor 简易API文档
// @termsOfService http://swagger.io/terms/

// @contact.name ink
// @contact.url http://jiangink.github.com
// @contact.email jiangink.cn@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /api/v1
func main() {
	// 基本配置初始化
	config.Setup()
	// 日志初始化
	logger.Setup()
	// 数据库初始化
	models.Setup()
	// 缓存初始化
	redis.Setup()
	// 调度任务初始化
	schedule.Setup()

	r := router.InitRouter()
	// 服务监听
	r.Run(":" + strconv.Itoa(config.ServerConf.Port))
}
