package config

var (
	Cfg Config
)

type Config struct {
	KubeConfig string
	Mode       string
}
