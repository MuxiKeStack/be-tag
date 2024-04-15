//go:build wireinject

package main

import (
	"github.com/MuxiKeStack/be-tag/grpc"
	"github.com/MuxiKeStack/be-tag/ioc"
	"github.com/MuxiKeStack/be-tag/pkg/grpcx"
	"github.com/MuxiKeStack/be-tag/repository"
	"github.com/MuxiKeStack/be-tag/repository/dao"
	"github.com/MuxiKeStack/be-tag/service"
	"github.com/google/wire"
)

func InitGRPCServer() grpcx.Server {
	wire.Build(
		ioc.InitGRPCxKratosServer,
		grpc.NewTagServiceServer,
		service.NewTagService,
		repository.NewTagRepository,
		dao.NewGORMTagDAO,
		ioc.InitEtcdClient,
		ioc.InitLogger,
		ioc.InitDB,
	)
	return grpcx.Server(nil)
}
