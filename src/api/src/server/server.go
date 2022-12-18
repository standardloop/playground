package server

import "api/src/config"

func Init() {
	r := NewRouter()
	r.Run(config.Env.AppPort)
}
