package global

import (
	"gitee.com/newtengroup/srv-id/internal/constants/enums"
	"github.com/eden-framework/courier/transport_grpc"
	"github.com/eden-framework/courier/transport_http"
	"github.com/sirupsen/logrus"
)

type SnowFlakeConfig struct {
	Epoch    int64
	NodeID   int64
	NodeBits uint8
	StepBits uint8
}

var Config = struct {
	LogLevel logrus.Level

	// administrator
	GRPCServer *transport_grpc.ServeGRPC
	HTTPServer *transport_http.ServeHTTP

	// module
	GenerateAlgorithm enums.GenerateAlgorithm
	SnowFlakeConfig   SnowFlakeConfig
}{
	LogLevel: logrus.DebugLevel,

	GRPCServer: &transport_grpc.ServeGRPC{
		Port: 8900,
	},
	HTTPServer: &transport_http.ServeHTTP{
		Port:     8800,
		WithCORS: true,
	},

	GenerateAlgorithm: enums.GENERATE_ALGORITHM__SNOWFLAKE,
	SnowFlakeConfig: SnowFlakeConfig{
		Epoch:    1002921769000,
		NodeID:   1,
		NodeBits: 10,
		StepBits: 12,
	},
}
