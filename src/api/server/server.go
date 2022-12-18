package server

import "api/config"

func Init() {
	r := NewRouter()
	r.Run(config.Env.AppPort)
}
