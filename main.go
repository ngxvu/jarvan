package main

import (
	"context"
	"gitlab.com/merakilab9/j4/conf"
	"gitlab.com/merakilab9/j4/pkg/route"
	"gitlab.com/merakilab9/j4/pkg/utils"
	"gitlab.com/merakilab9/meracore/logger"
	"os"
)

const (
	APPNAME = "J4"
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
