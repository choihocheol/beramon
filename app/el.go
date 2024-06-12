package app

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"bharvest.io/beramon/client/el"
	"bharvest.io/beramon/store"
	"bharvest.io/beramon/utils"
)

func (app *BaseApp) checkEL(ctx context.Context) error {
	// Init EL status
	store.GlobalState.EL.Status = true

	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(4)

	client, err := el.New(app.cfg.EL.JsonRPC)
	if err != nil {
		return err
	}

	// Check sync status
	go func() {
		defer wg.Done()

		isSyncing, err := client.GetSyncStatus(appCtx)
		if err != nil {
			utils.Error(err, true)
			return
		}
		store.GlobalState.EL.Sync = isSyncing

		if isSyncing {
			store.GlobalState.EL.Status = false

			msg := "EL Node is syncing"
			utils.SendTg(msg)
			utils.Error(errors.New(msg), false)

			return
		}
	}()

	// Check Latest Block
	go func() {
		defer wg.Done()

		height, err := client.GetLatestBlock(appCtx)
		if err != nil {
			utils.Error(err, true)
			return
		}
		store.GlobalState.EL.CurrentHeight = height
	}()

	// Check Peer Count
	go func() {
		defer wg.Done()

		peers, err := client.GetPeerCnt(appCtx)
		if err != nil {
			utils.Error(err, true)
			return
		}
		store.GlobalState.EL.Peers = peers

		if peers < app.cfg.EL.PeerThreshold {
			store.GlobalState.EL.Status = false

			msg := fmt.Sprintf("EL Node has low peers: %d", peers)
			utils.SendTg(msg)
			utils.Error(errors.New(msg), false)

			return
		}
	}()

	// Check Txpool Queued
	go func() {
		defer wg.Done()

		cnt, err := client.GetTxQueuedCnt(appCtx)
		if err != nil {
			utils.Error(err, true)
			return
		}
		store.GlobalState.EL.TxpoolQueued = cnt

		if cnt >= int(app.cfg.EL.TxpoolQueuedThreshold) {
			store.GlobalState.EL.Status = false

			msg := fmt.Sprintf("Txpool Queued is too high: %d", cnt)
			utils.SendTg(msg)
			utils.Error(errors.New(msg), false)

			return
		}
	}()

	wg.Wait()

	return nil
}
