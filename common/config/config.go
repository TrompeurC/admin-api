package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var Config *config

// 总配置文件
type config struct {
	Server        server        `yaml:"server"`
	DB            db            `yaml:"db"`
	Redis         redis         `yaml:"redis"`
	ImageSettings imageSettings `yaml:"imageSettings"`
	Log           log           `yaml:"log"`
}

// 项目端口配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type db struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MinOpen  int    `yaml:"minOpen"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

// imageSettings图片上传配置
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

// 日志配置
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

// 项目的初始化
func init() {
	yamlFile, err := os.ReadFile("./config.yaml")
	// 有错就 panic
	if err != nil {
		panic(err)
	}
	// 绑定值
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
