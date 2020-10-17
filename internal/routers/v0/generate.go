package v0

import (
	"context"
	"gitee.com/newtengroup/srv-id/internal/constants/errors"
	"gitee.com/newtengroup/srv-id/internal/global"
	"gitee.com/newtengroup/srv-id/internal/modules/algorithm"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(GenerateID{}))
}

// 生成ID
type GenerateID struct {
	httpx.MethodGet
}

func (req GenerateID) Path() string {
	return "/id"
}

type GenerateIDResp struct {
	ID uint64 `json:"id,string"`
}

func (req GenerateID) Output(ctx context.Context) (result interface{}, err error) {
	generator := algorithm.GeneratorContainerInstance.GetAlgorithm(global.Config.GenerateAlgorithm)
	if generator == nil {
		return nil, errors.AlgorithmNotFound
	}
	id, err := generator.GenerateUniqueID()
	if err != nil {
		return nil, errors.InternalError.StatusError().WithDesc(err.Error())
	}

	return GenerateIDResp{ID: id}, nil
}
