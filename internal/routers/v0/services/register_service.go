package services

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(RegisterService{}))
}

// 注册服务节点
type RegisterService struct {
	httpx.MethodPost
}

func (req RegisterService) Path() string {
	return ""
}

func (req RegisterService) Output(ctx context.Context) (result interface{}, err error) {
	return
}
