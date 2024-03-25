package consul

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func ServiceConsul(address, port, name string) error {

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}

	ip, _ := strconv.Atoi(port)

	return client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      "",
		Name:    name,
		Tags:    []string{"GRPC"},
		Port:    ip,
		Address: address,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%v:%v", address, port),
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "5s",
		},
	})

}

// 服务发现
func CheckConsul(name string) (string, error) {

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return "", err
	}

	_, i, err := client.Agent().AgentHealthServiceByName(name)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v:%v", i[0].Service.Address, i[0].Service.Port), nil

}

func Client2(ctx context.Context, toService string, port string) (*grpc.ClientConn, error) {
	//consul, err := consul.GetConfig(toService, port)

	sdn := fmt.Sprintf("consul://%v:%v/%v?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{\"LoadBalancingPolicy\": \"round_robin\"}`), "127.0.0.1", "8500", "user")

	conn, err := grpc.Dial(sdn)

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn, nil

}
