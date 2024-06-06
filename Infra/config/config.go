package config

import (
	"path/filepath"

	"github.com/linchengzhi/go-clean-backend/domain/dto"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func NewConfig(path string) (*dto.Config, error) {
	v := viper.New()

	dir, file := filepath.Split(path)
	fileName := file[:len(file)-len(filepath.Ext(file))] // 移除文件扩展名

	v.SetConfigName(fileName) // 设置配置文件的名称
	v.AddConfigPath(dir)      // 添加配置文件所在的目录
	v.SetConfigType("yaml")   // 设置配置文件的类型

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &dto.Config{}
	op := viper.DecoderConfigOption(func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
	})
	err = v.Unmarshal(c, op)
	if err != nil {
		return nil, err
	}
	return c, nil
}
