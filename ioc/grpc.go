package ioc

import (
	"github.com/MuxiKeStack/be-tag/grpc"
	"github.com/MuxiKeStack/be-tag/pkg/grpcx"
	"github.com/MuxiKeStack/be-tag/pkg/logger"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	grpc2 "github.com/seata/seata-go/pkg/integration/grpc"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitGRPCxKratosServer(tagServer *grpc.TagServiceServer, ecli *clientv3.Client, l logger.Logger) grpcx.Server {
	type Config struct {
		Name    string `yaml:"name"`
		Weight  int    `yaml:"weight"`
		Addr    string `yaml:"addr"`
		EtcdTTL int64  `yaml:"etcdTTL"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	server := kgrpc.NewServer(
		kgrpc.Address(cfg.Addr),
		kgrpc.Middleware(recovery.Recovery()),
		kgrpc.UnaryInterceptor(grpc2.ServerTransactionInterceptor),
	)
	tagServer.Register(server)
	return &grpcx.KratosServer{
		Server:     server,
		Name:       cfg.Name,
		Weight:     cfg.Weight,
		EtcdTTL:    time.Second * time.Duration(cfg.EtcdTTL),
		EtcdClient: ecli,
		L:          l,
	}
}
