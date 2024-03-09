package utils

import (
	"log"

	"github.com/robfig/cron/v3"
)

var CronList []cron.EntryID
var CronInstance *cron.Cron
var Crontab = map[string]func(){
	"@every 24h": func() {
		Logger.Info("软删除数据库中过期的用户token")
		CleanExpiredDatabaseUserToken()
	},
}

func SetupCron() {
	c := cron.New()

	//添加定时任务
	for spec, cmd := range Crontab {
		// 添加定时任务，每天凌晨 2 点执行一次清理操作
		id, err := c.AddFunc(spec, cmd)
		if err != nil {
			log.Panicln("添加定时任务时出错:", err)
		} else {
			CronList = append(CronList, id)
		}
	}

	// 启动 cron
	c.Start()
	CronInstance = c
}
