package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/NupalHariz/DD/src/business/usecase"
	"github.com/NupalHariz/DD/src/utils/config"
	"github.com/go-co-op/gocron"
	"github.com/reyhanmichiels/go-pkg/v2/log"
)

var (
	once = &sync.Once{}
)

type Interface interface {
	Run()
}

type SchedulerTaskConf struct {
	Name             string
	Enabled          bool
	TimeType         string
	Interval         time.Duration
	Weekday          time.Weekday
	ScheduledTime    string
	MultipleSchedule []string
}

type scheduler struct {
	cron     *gocron.Scheduler
	metaConf config.ApplicationMeta
	log      log.Interface
	uc       *usecase.Usecases
}

type InitParam struct {
	MetaConf config.ApplicationMeta
	Log      log.Interface
	Uc       *usecase.Usecases
}

func Init(param InitParam) Interface {
	s := &scheduler{}

	once.Do(func() {
		cron := gocron.NewScheduler(time.Local)
		cron.TagsUnique()

		s = &scheduler{
			cron:     cron,
			metaConf: param.MetaConf,
			log:      param.Log,
			uc:       param.Uc,
		}

		s.assignScheduledTask()
	})

	return s
}

func (s *scheduler) assignScheduledTask() {
	s.assignTask(
		SchedulerTaskConf{
			Name:          "ResetWeeklyBudget",
			Enabled:       true,
			TimeType:      schedulerTypeWeekly,
			ScheduledTime: "23:59",
		},
		s.uc.Budget.WeeklyResetScheduler,
	)

	s.assignTask(
		SchedulerTaskConf{
			Name:          "ResetMonthlyBudget",
			Enabled:       true,
			TimeType:      schedulerTypeMonthly,
			ScheduledTime: "23:59",
		},
		s.uc.Budget.WeeklyResetScheduler,
	)
}

func (s *scheduler) Run() {
	s.cron.StartAsync()
	s.log.Info(context.Background(), "scheduler is running on background")
}
