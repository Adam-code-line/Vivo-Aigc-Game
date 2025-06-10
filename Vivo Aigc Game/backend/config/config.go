package config

import "github.com/dingdinglz/vivo"

type Config struct {
	App     *vivo.Vivo
	Session string
}

func Load() (con *Config) {
	c := Config{
		vivo.NewVivoAIGC(vivo.Config{
			AppID:  "2025366283",
			AppKey: "vrVDQPSuAwNHreqM",
		}),
		vivo.GenerateSessionID(),
	}
	return &c
}
