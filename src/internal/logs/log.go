package logs

func Debug(msg string) {
	Logger.Debug().Msg(msg)
}

func Info(msg string) {
	Logger.Info().Msg(msg)
}

func Infof(format string, v ...interface{}) {
	Logger.Info().Msgf(format, v...)
}

func Warn(msg string) {
	Logger.Warn().Msg(msg)
}

func Error(msg string, err error) {
	Logger.Err(err).Msg(msg)
}

func CheckError(err error) {
	if err != nil {
		Logger.Err(err)
	}
}

func Fatal(msg string, err error) {
	Logger.Fatal().Err(err).Msg(msg)
}

func Panic(msg string) {
	Logger.Panic().Msg(msg)
}
