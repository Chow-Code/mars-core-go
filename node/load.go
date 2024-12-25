package node

import (
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"mars-go/util"
	"os"
	"strings"
)

func Load(path string) (Config, error) {
	var config Config
	err := util.LoadYAMLSection(path, "node-config", &config)
	if err != nil {
		log.Error("[node] 无法读取配置", err)
		return config, err
	}
	//校验必填项
	if config.NodeType == "" {
		return config, errors.New("node-type为必填项")
	}
	if config.TcpAddr.Port == 0 {
		return config, errors.New("tcp-addr.port为必填项")
	}
	//完善配置
	if config.Name == "" {
		nodeType := strings.ToLower(config.NodeType)
		config.Name, err = os.Hostname()
		if err == nil {
			config.Name = strings.ToLower(config.Name)
			if !strings.Contains(config.Name, nodeType) {
				config.Name = nodeType + "_" + config.Name
			}
			log.Info("[节点管理] 使用HostName命名本节点: %s", config.Name)
		} else {
			config.Name = nodeType + "_" + uuid.New().String()
			log.Info("[节点管理] 使用随机数命名本节点: %s", config.Name)
		}
	}
	log.Info(config.Name)
	return config, nil
}
