package configserver

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
)

var ErrNotSetConfig = errors.New("none config info...")

type ConfigServer interface {
	FromJsonBytes() ([]byte, error)
	Error() error
}

type configSever struct {
	ConfigServer
	configFile string
}

func NewConfigServer(configFile string, s ConfigServer) *configSever {
	return &configSever{
		ConfigServer: s,
		configFile:   configFile,
	}
}

func (s *configSever) MustLoad(v any) error {
	//if s.ConfigServer.Error() != nil {
	//	return s.ConfigServer.Error()
	//}
	if s.configFile == "" && s.ConfigServer == nil {
		return ErrNotSetConfig
	}
	if s.ConfigServer == nil {
		// 使用go-zero默认方式
		conf.MustLoad(s.configFile, v)
		return nil
	}
	data, err := s.ConfigServer.FromJsonBytes()
	if err != nil {
		return err
	}
	return conf.LoadFromJsonBytes(data, v)
}

func (s *configSever) Error() error {
	return s.ConfigServer.Error()
}
