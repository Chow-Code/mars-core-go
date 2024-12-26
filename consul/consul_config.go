package consul

type Config struct {
	Address       string `yaml:"address"`
	Scheme        string `yaml:"scheme"`
	Port          int    `yaml:"port"`
	CheckTimeOut  string `yaml:"check-timeout"`
	CheckInterval string `yaml:"check-interval"`
	Deregister    string `yaml:"deregister"`
	DefaultStatus string `yaml:"default-status"`
}
