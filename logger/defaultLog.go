package logger

import (
	"fmt"
	"log"
	"os"
)

//DefaultLoggerFactory 默认os.Stdout输出
type DefaultLoggerFactory struct {
}

//DefaultLogger struct
type DefaultLogger struct {
	l    *log.Logger
	name string
}

//GetLogger get the logger from DefaultLoggerFactory
func (Df *DefaultLoggerFactory) GetLogger(name string) Logger {
	return &DefaultLogger{l: log.New(os.Stdout, fmt.Sprintf("<%s> ", name), log.Ldate|log.Ltime|log.Lshortfile), name: name}
}

//Name get the name from DefaultLoggerFactory
func (d *DefaultLogger) Name() string {
	return d.name
}

//Info logger.info
func (d *DefaultLogger) Info(msg string) {
	d.l.Print(msg)
}

//Infof logger.Infof
func (d *DefaultLogger) Infof(format string, v ...interface{}) {
	d.l.Printf(format, v)
}

//Infoln logger.Infoln
func (d *DefaultLogger) Infoln(v ...interface{}) {
	d.l.Println(v)
}

//Debug logger.Debug
func (d *DefaultLogger) Debug(msg string) {
	d.l.Print(msg)
}

//Debugf logger.Debugf
func (d *DefaultLogger) Debugf(format string, v ...interface{}) {
	d.l.Printf(format, v)
}

//Debugln logger.Debugln
func (d *DefaultLogger) Debugln(v ...interface{}) {
	d.l.Println(v)
}

//Warn logger.Warn
func (d *DefaultLogger) Warn(msg string) {
	d.l.Print(msg)
}

//Warnf logger.Warnf
func (d *DefaultLogger) Warnf(format string, v ...interface{}) {
	d.l.Printf(format, v)
}

//Warnln logger.Warnln
func (d *DefaultLogger) Warnln(v ...interface{}) {
	d.l.Println(v)
}

//Error logger.Error
func (d *DefaultLogger) Error(msg string) {
	d.l.Print(msg)
}

//Errorf logger.Errorf
func (d *DefaultLogger) Errorf(format string, v ...interface{}) {
	d.l.Printf(format, v)
}

//Errorln logger.Errorln
func (d *DefaultLogger) Errorln(v ...interface{}) {
	d.l.Println(v)
}

//Panic logger.Panic
func (d *DefaultLogger) Panic(msg string) {
	d.l.Panic(msg)
}

//Panicf logger.Panicf
func (d *DefaultLogger) Panicf(format string, v ...interface{}) {
	d.l.Panicf(format, v)
}

//Panicln logger.Panicln
func (d *DefaultLogger) Panicln(v ...interface{}) {
	d.l.Panicln(v)
}

//Fatal logger.Fatal
func (d *DefaultLogger) Fatal(msg string) {
	d.l.Fatal(msg)
}

//Fatalf logger.Fatalf
func (d *DefaultLogger) Fatalf(format string, v ...interface{}) {
	d.l.Panicf(format, v)
}

//Fatalln logger.Fatalln
func (d *DefaultLogger) Fatalln(v ...interface{}) {
	d.l.Panicln(v)
}
