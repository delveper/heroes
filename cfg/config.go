package cfg

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Helpful conversion yaml => go
// https://zhwt.github.io/yaml-to-go/

type Options struct {
	Repo Postgres `yaml:"postgres"`
	HTTP HTTP     `yaml:"http"`
	API  API      `yaml:"api"`
}

type Postgres struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	UserName   string `yaml:"user_name"`
	Password   string `yaml:"password"`
	DbName     string `yaml:"db_name"`
	DriverName string `yaml:"driver_name"`
}

type HTTP struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
}

type API struct {
	User User `yaml:"user"`
}

type User struct {
	Endpoint        string   `yaml:"endpoint"`
	MandatoryParams []string `yaml:"mandatory_params"`
}

const defaultPath = "./cfg/config.yml"

func NewOptions() (*Options, error) { return load(&Options{}) }

// Load gives a config from file
// we can pass here any struct if we wish retrieve partial config
func load[C *Options](dst C) (C, error) {
	file, err := os.Open(defaultPath)
	if err != nil {
		return nil, err
	}

	defer func(*os.File) {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}(file)

	if err := yaml.NewDecoder(file).Decode(&dst); err != nil {
		return nil, err
	}

	return dst, nil
}
