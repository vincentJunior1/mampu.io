package logger

type LoggerInterface interface {
	Println(args ...interface{})
	Errorln(args ...interface{})
}
