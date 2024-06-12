package app

import (
	"context"
	"sync"

	"bharvest.io/beramon/utils"
)

func NewBaseApp(cfg *Config) *BaseApp {
	return &BaseApp{
		cfg: cfg,
	}
}

func (app *BaseApp) Run(ctx context.Context) {
	appCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Check CL
	go func() {
		defer wg.Done()

		utils.Info("Checking CL")

		err := app.checkCL(appCtx)
		if err != nil {
			utils.Error(err, true)
			return
		}
	}()

	// Check EL
	go func() {
		defer wg.Done()

		utils.Info("Checking EL")

		err := app.checkEL(appCtx)
		if err != nil {
			utils.Error(err, true)
			return
		}
	}()

	wg.Wait()

	return
}
