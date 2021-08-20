package main

import (
	"time"
	
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	
	"github.com/seizadi/cost-insight-backend/pkg/pb"
	"github.com/seizadi/cost-insight-backend/pkg/svc"
)

func CreateServer(logger *logrus.Logger, interceptors []grpc.UnaryServerInterceptor) (*grpc.Server, error) {
	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
			},
		), grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))
	
	// register all of our services into the grpcServer
	s, err := svc.NewBasicServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsCostServer(grpcServer, s)
	
	cs, err := svc.NewCostInsightApiServerServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterCostInsightsApiServer(grpcServer, cs)
	
	return grpcServer, nil
}
