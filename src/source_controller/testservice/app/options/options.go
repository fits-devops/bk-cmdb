package options

import (
	"configcenter/src/common/core/cc/config"
	"configcenter/src/storage/dal/mongo"
	"configcenter/src/storage/dal/redis"

	"github.com/spf13/pflag"
)

//ServerOption define option of server in flags
type ServerOption struct {
	ServConf *config.CCAPIConfig
}

// Config export
type Config struct {
	Mongo mongo.Config
	Redis redis.Config
}

//NewServerOption create a ServerOption object
func NewServerOption() *ServerOption {
	s := ServerOption{
		ServConf: config.NewCCAPIConfig(),
	}

	return &s
}

//AddFlags add flags
func (s *ServerOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.ServConf.AddrPort, "addrport", "127.0.0.1:30001", "The ip address and port for the serve on")
	fs.StringVar(&s.ServConf.RegDiscover, "regdiscv", "", "hosts of register and discover server. e.g: 127.0.0.1:2181")
	fs.StringVar(&s.ServConf.ExConfig, "config", "", "The config path. e.g conf/api.conf")
}
