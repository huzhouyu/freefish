package main

import (
	"github.com/freefishgo/freeFishGo/middlewares/printTimeMiddleware"
	"{{[.ProjectName]}}/conf"
	_ "{{[.ProjectName]}}/controllers"
	"{{[.ProjectName]}}/fishgo"
	_ "{{[.ProjectName]}}/routers"
)

func main() {
	// 通过注册中间件来打印任务处理时间服务
	conf.Build.UseMiddleware(&printTimeMiddleware.PrintTimeMiddleware{})
	// 利用中间件来实现http到https的转换
	//conf.Build.UseMiddleware(&httpToHttps.HttpToHttps{})
	// 把mvc中间件注册到管道中
	conf.Build.UseMiddleware(fishgo.Mvc)
	conf.Build.Config.Listen.HTTPPort = 8080
	conf.Build.Run()
}
