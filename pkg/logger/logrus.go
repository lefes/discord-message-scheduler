package logger

func Debug(msg ...interface{}) {
	Log.Debug(msg...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Info(msg ...interface{}) {
	Log.Info(msg...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Warn(msg ...interface{}) {
	Log.Warn(msg...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Error(msg ...interface{}) {
	Log.Error(msg...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}
