package main

import (
	"gitee.com/newtengroup/srv-id/internal/global"
	"gitee.com/newtengroup/srv-id/internal/routers"
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/sirupsen/logrus"
)

func main() {
	app := application.NewApplication(runner,
		application.WithConfig(&global.Config),
		application.WithApollo(&global.ApolloConfig))

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	logrus.SetLevel(global.Config.LogLevel)
	go global.Config.GRPCServer.Serve(ctx, routers.Router)
	return global.Config.HTTPServer.Serve(ctx, routers.Router)
}
