package main

import (
	"context"
	"gitlab.com/merakilab9/meracrawler/fortune/conf"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/route"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/utils"
	"os"

	"gitlab.com/merakilab9/meracore/logger"
)

const (
	APPNAME = "4tune"
)

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()

	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
