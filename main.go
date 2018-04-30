package main

import (
	"github.com/astaxie/beego"
	"github.com/lorock/PPGo_Job/jobs"
	"github.com/lorock/PPGo_Job/models"
	_ "github.com/lorock/PPGo_Job/routers"
)

const (
	// VERSION 版本号
	VERSION = "1.0.0"
)

func init() {
	//初始化数据模型
	models.Init()
	jobs.InitJobs()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
