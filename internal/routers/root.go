package routers

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/swagger"

	v0 "gitee.com/newtengroup/srv-id/internal/routers/v0"
)

var Router = courier.NewRouter(RootRouter{})

func init() {
	Router.Register(v0.Router)
	Router.Register(swagger.SwaggerRouter)
}

type RootRouter struct {
	courier.EmptyOperator
}

func (RootRouter) Path() string {
	return "/id"
}
