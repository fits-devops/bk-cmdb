package options

import (
	"github.com/spf13/pflag"

	"configcenter/src/auth/authcenter"
	"configcenter/src/common/core/cc/config"
	"configcenter/src/storage/dal/mongo"
)

type ServerOption struct {
	ServConf *config.CCAPIConfig
}

type Config struct {
	BusinessTopoLevelMax int `json:"level.businessTopoMax"`
	Mongo                mongo.Config
	ConfigMap            map[string]string
	Auth                 authcenter.AuthConfig
}

func NewServerOption() *ServerOption {
	s := ServerOption{
		ServConf: config.NewCCAPIConfig(),
	}

	return &s
}

func (s *ServerOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.ServConf.AddrPort, "addrport", "127.0.0.1:40001", "The ip address and port for the serve on")
	fs.StringVar(&s.ServConf.RegDiscover, "regdiscv", "", "hosts of register and discover server. e.g: 127.0.0.1:2181")
	fs.StringVar(&s.ServConf.ExConfig, "config", "", "The config path. e.g conf/api.conf")
}
