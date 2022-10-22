package cfg

import (
	_ "embed"
	"time"

	"gopkg.in/yaml.v3"
)

// Helpful conversion yaml => go
// https://zhwt.github.io/yaml-to-go/

type Options struct {
	Repo Repo `yaml:"postgres"`
	HTTP HTTP `yaml:"http"`
	API  API  `yaml:"api"`
}

type Repo struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	UserName   string `yaml:"user_name"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db_name"`
	DriverName string `yaml:"driver_name"`
}

type HTTP struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type API struct {
	User User `yaml:"user"`
}

type User struct {
	Endpoint        string   `yaml:"endpoint"`
	MandatoryParams []string `yaml:"mandatory_params"`
}

func NewOptions() (*Options, error) { return load(&Options{}) }

var configYml []byte

// load helps to retrieve any king of config from file
func load[C *Options | *Repo | *HTTP | *API | *User](dst C) (C, error) {
	if err := yaml.Unmarshal(configYml, &dst); err != nil {
		return nil, err
	}

	return dst, nil
}
