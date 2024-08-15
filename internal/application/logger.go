package application

type Logger interface {
	Info(fields ...interface{})
	Error(fields ...interface{})
	Fatal(fields ...interface{})
}
