package configuration

type Config struct {
	ListenAddr string `yaml:"listenAddr"`
	WWWPath    string `yaml:"wwwPath"`
}
