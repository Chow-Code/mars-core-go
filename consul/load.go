package consul

import (
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"mars-go/node"
	"mars-go/util"
)

const CONFIG_SECTION = "consul"

func Load(nodeConfig node.Config, configPath string, otherInfo string) error {
	var config Config
	err := util.LoadYAMLSection(configPath, CONFIG_SECTION, &config)
	if err != nil {
		log.Error("[consul] 无法读取配置", err)
		return err
	}
	client, err := api.NewClient(&api.Config{
		Address: config.Address,
		Scheme:  config.Scheme,
	})
	if err != nil {
		log.Error("[consul] 无法启动客户端", err)
		return err
	}
	// 注册服务
	registration := &api.AgentServiceRegistration{
		ID:      nodeConfig.Name,
		Name:    nodeConfig.NodeType,
		Port:    config.Port,
		Tags:    []string{otherInfo},
		Address: nodeConfig.TcpAddr.Host,
	}
	//健康检查

	err = client.Agent().ServiceRegister(registration)
	return err
}
