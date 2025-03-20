package jobs

import (
	"github.com/jasonlvhit/gocron"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Invoke(InitJobs)

type Params struct {
	fx.In
	// Logger
	Logger zerolog.Logger
	// GoCron
	Cron *gocron.Scheduler
}
