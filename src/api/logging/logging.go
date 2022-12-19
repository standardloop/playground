package logging

import (
	"api/config"

	"github.com/rs/zerolog"
)

func Init() {
	switch config.Env.LogLevel {
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
		break
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		break
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		break
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
