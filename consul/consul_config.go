package consul

type Config struct {
	Address string `yaml:"address"`
	Scheme  string `yaml:"scheme"`
	Port    int    `yaml:"port"`
}
