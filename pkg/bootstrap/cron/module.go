package cron

import (
	"github.com/jasonlvhit/gocron"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewCron)

type Params struct {
	fx.In

	// Locker
}

// NewCron ...
func NewCron(params Params) *gocron.Scheduler {

	return gocron.NewScheduler()
}
