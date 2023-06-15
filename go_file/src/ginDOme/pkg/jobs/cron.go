package jobs

import "github.com/robfig/cron"

var mainCron *cron.Cron

func init() {
	mainCron = cron.New()
	mainCron.Start()
}

func InitJobs() {
	// 每5s钟调度一次，并传参
	//mainCron.AddJob(
	//	"*/50 * * * * ?",
	//	TestJob{Id: 1, Name: "zhangsan"},
	//)
}
