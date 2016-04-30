package config

import (
	"flag"
	"github.com/jinzhu/configor"
	"log"
	"os"
)

type AppConfig struct {
	Server struct {
		       Listen string `default:":8001"`
	       }
	Mysql  struct {
		       Host    string `default:"localhost:3306"`
		       Port    string `default:"3306"`
		       User    string `default:"root"`
		       Pass    string `default:""`
		       DBName  string `default:"test"`
		       TBLName string `default:"test"`
	       }
}

func Load() AppConfig {
	config := flag.String("config", "./config/config.local.yml", "configuration yml file")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		//os.Exit(1)
	}

	if _, err := os.Stat(*config); os.IsNotExist(err) {
		log.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	var appConfig AppConfig
	if err := configor.Load(&appConfig, *config); err != nil {
		log.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("Server conf: %+v", appConfig.Server)

	return appConfig
}
