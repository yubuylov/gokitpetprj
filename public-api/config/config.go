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
		       CvcKey string `default:":12345"`
		       Qps    int64  `default:"1000"`
	       }
	Relations  struct {
		       Storage       string `default:"localhost:36701"`
		       MicroServiceA string `default:"localhost:36801"`
		       MicroServiceZ string `default:"localhost:36808"`
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
