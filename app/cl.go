package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"bharvest.io/beramon/client/cl"
	"bharvest.io/beramon/store"
	"bharvest.io/beramon/utils"
)

func (app *BaseApp) checkCL(ctx context.Context) error {
	// Init CL status
	store.GlobalState.CL.Status = true
	store.GlobalState.CL.ValidatorAddr = app.cfg.CL.ValidatorAddress

	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := cl.New(app.cfg.CL.RPC)
	if err != nil {
		return err
	}

	resp, err := client.GetValidatorInfo(appCtx, app.cfg.CL.ValidatorAddress)
	if err != nil {
		return err
	}

	var window uint64 = 10000
	missCnt := window - resp.CommitCount
	store.GlobalState.CL.Missed = fmt.Sprintf("%d / %d", missCnt, window)

	if missCnt >= app.cfg.CL.MissThreshold {
		store.GlobalState.CL.Status = false

		msg := fmt.Sprintf("CL Node is missing blocks : %s", store.GlobalState.CL.Missed)
		utils.SendTg(msg)
		utils.Error(errors.New(msg), false)

		return errors.New(msg)
	}

	return nil
}
