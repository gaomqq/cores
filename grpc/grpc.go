package grpc

import (
	"core/config"
	"core/consul"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	held "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func RegisterRpc(f func(s *grpc.Server)) error {
	cos, _ := config.ServiceNaCos()
	c := cos.App
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", c.Ip, c.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	//支持健康检查
	held.RegisterHealthServer(s, health.NewServer())
	//商品服务注册
	err = consul.ServiceConsul(c.Ip, c.Port, c.Secret)

	if err != nil {
		return err
	}
	f(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil

}
