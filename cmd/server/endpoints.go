package main
import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/spf13/viper"
	
	"github.com/seizadi/aws-cost/pkg/pb"
)

func RegisterGatewayEndpoints() gateway.Option {
	return gateway.WithEndpointRegistration(viper.GetString("gateway.endpoint"),
		pb.RegisterAwsCostHandlerFromEndpoint,
		pb.RegisterCostInsightsApiHandlerFromEndpoint,
	)
}
