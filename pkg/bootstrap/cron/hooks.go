package cron

import (
	"context"

	"github.com/jasonlvhit/gocron"
	"go.uber.org/fx"
)

// ModuleLifecycleCronHooks ...
var ModuleLifecycleCronHooks = fx.Invoke(RegisterCronHooks)

// RegisterCronHooks ...
func RegisterCronHooks(lifecycle fx.Lifecycle, cron *gocron.Scheduler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go cron.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cron.Clear()
			return nil
		},
	})
}
