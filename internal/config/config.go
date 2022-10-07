package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server         serverConf
	ChromeHeadless chromeHeadlessConf
}

type serverConf struct {
	Port string `env:"HTMLTOIMAGE_SERVER_PORT,required"`
	IP   string `env:"HTMLTOIMAGE_SERVER_IP,required"`
}

type chromeHeadlessConf struct {
	Address string `env:"HTMLTOIMAGE_SERVER_CHROME_HEADLESS_WS_URL,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
