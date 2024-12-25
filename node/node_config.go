package node

type Config struct {
	//节点类型
	NodeType string `yaml:"node-type"`
	//节点名称
	Name string `yaml:"name"`
	//节点TCP链接地址
	TcpAddr NetAddr `yaml:"tcp-addr"`
	//节点的HTTP接口地址
	HttpAddr NetAddr `yaml:"http-addr"`
	//该节点能处理哪些微服务消息类型
	MessageTypes []int `yaml:"message-types"`
	//该节点能处理哪些游戏类型
	GameTypes []int `yaml:"game-types"`
	//节点权重
	Weight int `yaml:"weight"`
	//节点的其他信息
	OtherInfo string `yaml:"other-info"`
}

// NetAddr 网络数据结构
type NetAddr struct {
	//ip地址或者主机名称
	Host string `yaml:"host"`
	//端口号
	Port uint16 `yaml:"port"`
}
