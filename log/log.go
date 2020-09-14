package log

var (
	infoLog  logger = &infoLogger{}
	errorLog logger = &errorLogger{}
	debugLog logger = &debugLogger{}
)

type logger interface {
	logOut(format *string, v ...interface{})
	init()
}

func Info(args ...interface{})  {
	infoLog.logOut(nil, args)
}

func FInfo(format string, args ...interface{})  {
	infoLog.logOut(&format, args...)
}

func Error(args ...interface{})  {
	errorLog.logOut(nil, args)
}

func Debug(args ...interface{})  {
	debugLog.logOut(nil, args)
}
