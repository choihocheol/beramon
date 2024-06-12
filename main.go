package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"bharvest.io/beramon/app"
	"bharvest.io/beramon/server"
	"bharvest.io/beramon/utils"
	"github.com/pelletier/go-toml/v2"
)

func main() {
	ctx := context.Background()

	f, err := os.ReadFile("config.toml")
	if err != nil {
		utils.Error(err, true)
		return
	}
	cfg := app.Config{}
	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		utils.Error(err, true)
		return
	}

	tgTitle := fmt.Sprintf("🤖 beramon 🤖")
	utils.SetTg(cfg.Tg.Enable, tgTitle, cfg.Tg.Token, cfg.Tg.ChatID)

	go server.Run(cfg.General.APIListenPort)

	baseapp := app.NewBaseApp(&cfg)
	for {
		baseapp.Run(ctx)
		time.Sleep(time.Duration(cfg.General.Period) * time.Minute)
	}
}
