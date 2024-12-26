package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"mars-go/node"
	"mars-go/util"
	"sync/atomic"
	"time"
)

const CONFIG_SECTION = "consul"

var client *api.Client

var initialized *atomic.Bool

func Init(nodeConfig node.Config, configPath string, otherInfo string) error {
	if initialized.Load() {
		log.Error("[consul] 无法初始化, 因为已经初始化过了")
		return fmt.Errorf("无法初始化conul, consul已初始化过了")
	}
	initialized.Store(true)
	config, err := LoadConfig(configPath)
	if err != nil {
		return err
	}
	err = Register(nodeConfig, config, otherInfo)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(configPath string) (Config, error) {
	var config Config
	err := util.LoadYAMLSection(configPath, CONFIG_SECTION, &config)
	if err != nil {
		log.Error("[consul] 无法读取配置", err)
		return config, err
	}
	return config, err
}

func Register(nodeConfig node.Config, config Config, otherInfo string) error {
	// 注册服务
	client, err := api.NewClient(&api.Config{
		Address: config.Address,
		Scheme:  config.Scheme,
	})
	if err != nil {
		log.Error("[consul] 无法启动客户端", err)
		return err
	}
	registration := &api.AgentServiceRegistration{
		ID:      nodeConfig.Name,
		Name:    nodeConfig.NodeType,
		Port:    config.Port,
		Tags:    []string{otherInfo},
		Address: nodeConfig.TcpAddr.Host,
		//健康检查
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d%s", nodeConfig.TcpAddr.Host, config.Port, "/health-check"),
			Timeout:                        config.CheckTimeOut,
			Interval:                       config.CheckInterval,
			DeregisterCriticalServiceAfter: config.Deregister,
			Status:                         config.DefaultStatus,
		},
	}
	err = client.Agent().ServiceRegister(registration)
	return err
}

func listen() {
	for {
		// 监听服务
		services, _, err := client.Catalog().Services(nil)
		if err != nil {
			log.Printf("Error retrieving services: %v", err)
			continue
		}

		for name, tags := range services {
			log.Printf("Service: %s, Tags: %v", name, tags)
		}

		time.Sleep(5 * time.Second) // 每5秒查询一次
	}
}

func Deregister() {
	if client == nil {
		log.Error("[consul] consul未注册, 不可注销")
		return
	}
}

func FindService() {
	if client == nil {
		log.Error("[consul] consul未注册, 不可发现服务")
	}
	services, err := client.Agent().Services()
	if err != nil {
		log.Error(err)
		return
	}
	for _, service := range services {
		log.Info("[consul] 收到节点信息:", service.Address, service.Port, service.ID)
	}
}
