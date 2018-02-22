package mylogger // import "gitlab.csos95.com/csos95/mylogger"

import "log"

type logger struct {
}

func New() *logger {
	return &logger{}
}

func (l *logger) Println(values ...interface{}) {
	log.Println(values...)
}

func (l *logger) Printf(template string, values ...interface{}) {
	log.Printf(template, values...)
}
