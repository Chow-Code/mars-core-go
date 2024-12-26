package app

import (
	log "github.com/sirupsen/logrus"
	"mars-go/consul"
	"mars-go/node"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func Run(configPath string) {
	//加载Node配置
	log.Info("开始加载配置")
	nodeConfig, err := node.Load(configPath)
	if err != nil {
		log.Info(err.Error())
		log.Info("======== 节点配置加载失败, 进程结束 ========")
		return
	}
	log.Info("配置加载成功 nodeConfig:", nodeConfig)
	err = consul.Init(nodeConfig, configPath, "")
	if err != nil {
		log.Info(err.Error())
		log.Info("======== consul启动失败 ========")
	}
	log.Info("======== consul注册成功 ========")
}
