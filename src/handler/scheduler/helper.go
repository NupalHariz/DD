package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/reyhanmichiels/go-pkg/v2/appcontext"
)

const (
	schedulerUserAgent = "cron scheduler: %s"

	schedulerRunning      string = "scheduler %s is running"
	schedulerDoneSuccess  string = "scheduler %s is success"
	schedulerDoneError    string = "scheduler %s is error: %v"
	scheduleTimeExecution string = "scheduler %s done in %v"

	schedulerTypeWeekly  string = "weekly"
	schedulerTypeMonthly string = "monthly"
)

type handlerFunc func(ctx context.Context) error

func (s *scheduler) createContext(conf SchedulerTaskConf) context.Context {
	ctx := context.Background()

	ctx = appcontext.SetUserAgent(ctx, fmt.Sprintf(schedulerUserAgent, conf.Name))
	ctx = appcontext.SetRequestId(ctx, uuid.NewString())
	ctx = appcontext.SetRequestStartTime(ctx, time.Now())
	ctx = appcontext.SetServiceVersion(ctx, s.metaConf.Version)

	return ctx
}

func (s *scheduler) assignTask(conf SchedulerTaskConf, task handlerFunc) {
	if conf.Enabled {
		var err error
		ctx := context.Background()
		schedulerFunc := s.taskWrapper(conf, task)

		switch conf.TimeType {
		case schedulerTypeWeekly:
			_, err = s.cron.Every(1).Week().Sunday().Tag(conf.Name).At(conf.ScheduledTime).Do(schedulerFunc)
		case schedulerTypeMonthly:
			_, err = s.cron.Every(1).Month(-1).Tag(conf.Name).At(conf.ScheduledTime).Do(schedulerFunc)
		}

		if err != nil {
			s.log.Fatal(ctx, fmt.Sprintf(schedulerDoneError, conf.Name, err))
		}
	}
}

func (s *scheduler) taskWrapper(conf SchedulerTaskConf, task handlerFunc) func() {
	return func() {
		ctx := s.createContext(conf)
		s.log.Info(ctx, fmt.Sprintf(schedulerRunning, conf.Name))

		if err := task(ctx); err != nil {
			s.log.Error(ctx, fmt.Sprintf(schedulerDoneError, conf.Name, err))
		} else {
			s.log.Info(ctx, fmt.Sprintf(schedulerDoneSuccess, conf.Name))
		}

		startTime := appcontext.GetRequestStartTime(ctx)
		s.log.Info(ctx, fmt.Sprintf(scheduleTimeExecution, conf.Name, startTime))
	}
}
