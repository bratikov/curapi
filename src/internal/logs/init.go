package logs

import (
	"currency/internal/config"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger zerolog.Logger
)

type Config struct {
	File  bool
	Path  string
	Level string
}

func Init(config *config.Log) {
	output := &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    50,   // Size in MB before file gets rotated
		MaxBackups: 3,    // Max number of files kept before being
		MaxAge:     28,   // Max number of days to keep the files
		Compress:   true, // Whether to compress log files using gzip
	}

	switch strings.ToLower(config.Level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	Logger = zerolog.New(output).With().Timestamp().Logger()
}
