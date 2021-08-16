package main

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGRPCServer(logger *logrus.Logger) (*grpc.Server, error) {
	interceptors := []grpc.UnaryServerInterceptor{
			// logging middleware
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
			
			// Request-Id interceptor
			requestid.UnaryServerInterceptor(),
			
			// Metrics middleware
			grpc_prometheus.UnaryServerInterceptor,
			
			// validation middleware
			grpc_validator.UnaryServerInterceptor(),
			
			// collection operators middleware
			gateway.UnaryServerInterceptor(),
	}
	
	return CreateServer(logger, interceptors)
}
