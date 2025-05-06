package scheduler

import (
	"currency/internal/logs"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	cron *gocron.Scheduler
)

func AddCron(name string, crona func(), timer int) {
	_, err := cron.Every(timer).Seconds().Tag(name).Name(name).Do(crona)
	if err != nil {
		logs.Fatal("Cant create scheduler job: "+name, err)
	}
}

func AddCronDayAt(name string, crona func(), every int, at string) {
	_, err := cron.Every(every).Day().At(at).Tag(name).Name(name).Do(crona)
	if err != nil {
		logs.Fatal("Cant create scheduler job: "+name, err)
	}
}

func HasCron(name string) bool {
	crons, err := cron.FindJobsByTag(name)
	if err != nil {
		logs.Info("Cant find job by tag: " + name)
	}

	if len(crons) == 0 {
		return false
	}

	return true
}

func init() {
	cron = gocron.NewScheduler(time.UTC)
	cron.StartAsync()
}
