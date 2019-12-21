package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/requilence/integram"
)

func main() {
	var cfg Config
	envconfig.MustProcess("GITLAB", &cfg)

	integram.Register(cfg, cfg.BotConfig.Token)

	integram.Run()
}
