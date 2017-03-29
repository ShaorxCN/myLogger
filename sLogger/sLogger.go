/*
	基于logger接口实现的简单logger,参考部分log4j功能
	暂时支持控制台以及文件作为输出源
	就不写大文件分割以及周期生成文件了
*/
package sLogger

import (
	"fmt"
	"github.com/ShaorxCN/myLogger/confutil"
	"github.com/ShaorxCN/myLogger/logger"
	"io"
	"log"
	"os"
)

var warnFile, debugFile, infoFile, errorFile, panicFile, fatalFile string

type sLoggerFaactory struct {
}

type sLogger struct {
	l    *log.Logger
	name string
}

func init() {
	slf := new(sLoggerFaactory)
	logger.RegisterLogger(slf)
}

func (slf *sLoggerFaactory) SetConfig(path string) {
	config := new(confutil.Config)
	config.Anly(path)
	warnFile = config.Secs["default"]["warnFile"]
	infoFile = config.Secs["default"]["infoFile"]
	errorFile = config.Secs["default"]["errorFile"]
	panicFile = config.Secs["default"]["panicFile"]
	fatalFile = config.Secs["default"]["fatalFile"]
	debugFile = config.Secs["default"]["debugFile"]

}

func (slf *sLoggerFaactory) GetLogger(name string) logger.Logger {
	return &sLogger{l: log.New(nil, "", log.Ldate|log.Ltime), name: name}
}

func (d *sLogger) Name() string {

	return d.name
}

func (d *sLogger) Info(msg string) {
	d.l.SetPrefix(fmt.Sprintf("<Info><%s>", d.name))

	var output io.Writer = os.Stdout
	if infoFile != "" {
		file, err := os.OpenFile(infoFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}

	d.l.SetOutput(output)
	d.l.Print(msg)
}

func (d *sLogger) Infof(format string, v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Info><%s>", d.name))

	var output io.Writer = os.Stdout
	if infoFile != "" {
		file, err := os.OpenFile(infoFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}

	d.l.SetOutput(output)

	d.l.Printf(format, v)
}

func (d *sLogger) Infoln(v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Info><%s>", d.name))

	var output io.Writer = os.Stdout
	if infoFile != "" {
		file, err := os.OpenFile(infoFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}

	d.l.SetOutput(output)
	d.l.Println(v)
}

func (d *sLogger) Debug(msg string) {
	d.l.SetPrefix(fmt.Sprintf("<Debug><%s>", d.name))

	var output io.Writer = os.Stdout
	if debugFile != "" {
		file, err := os.OpenFile(debugFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}

	d.l.SetOutput(output)
	d.l.Print(msg)
}

func (d *sLogger) Debugf(format string, v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Debug><%s>", d.name))

	var output io.Writer = os.Stdout
	if debugFile != "" {
		file, err := os.OpenFile(debugFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Printf(format, v)
}

func (d *sLogger) Debugln(v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Debug><%s>", d.name))

	var output io.Writer = os.Stdout
	if debugFile != "" {
		file, err := os.OpenFile(debugFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Println(v)
}

func (d *sLogger) Warn(msg string) {
	d.l.SetPrefix(fmt.Sprintf("<Warn><%s>", d.name))

	var output io.Writer = os.Stdout
	if warnFile != "" {
		file, err := os.OpenFile(warnFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Print(msg)
}

func (d *sLogger) Warnf(format string, v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Warn><%s>", d.name))

	var output io.Writer = os.Stdout
	if warnFile != "" {
		file, err := os.OpenFile(warnFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Printf(format, v)
}

func (d *sLogger) Warnln(v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Warn><%s>", d.name))

	var output io.Writer = os.Stdout
	if warnFile != "" {
		file, err := os.OpenFile(warnFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Println(v)
}

func (d *sLogger) Error(msg string) {
	d.l.SetPrefix(fmt.Sprintf("<Error><%s>", d.name))

	var output io.Writer = os.Stdout
	if errorFile != "" {
		file, err := os.OpenFile(errorFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Print(msg)
}

func (d *sLogger) Errorf(format string, v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Error><%s>", d.name))

	var output io.Writer = os.Stdout
	if errorFile != "" {
		file, err := os.OpenFile(errorFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Printf(format, v)
}

func (d *sLogger) Errorln(v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Error><%s>", d.name))

	var output io.Writer = os.Stdout
	if errorFile != "" {
		file, err := os.OpenFile(errorFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}

	d.l.SetOutput(output)
	d.l.Println(v)
}

func (d *sLogger) Panic(msg string) {
	d.l.SetPrefix(fmt.Sprintf("<Panic><%s>", d.name))

	var output io.Writer = os.Stdout
	if panicFile != "" {
		file, err := os.OpenFile(panicFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Panic(msg)
}

func (d *sLogger) Panicf(format string, v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Panic><%s>", d.name))

	var output io.Writer = os.Stdout
	if panicFile != "" {
		file, err := os.OpenFile(panicFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Panicf(format, v)
}

func (d *sLogger) Panicln(v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Panic><%s>", d.name))

	var output io.Writer = os.Stdout
	if panicFile != "" {
		file, err := os.OpenFile(panicFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Panicln(v)
}

func (d *sLogger) Fatal(msg string) {
	d.l.SetPrefix(fmt.Sprintf("<Fatal><%s>", d.name))

	var output io.Writer = os.Stdout
	if fatalFile != "" {
		file, err := os.OpenFile(fatalFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Fatal(msg)
}

func (d *sLogger) Fatalf(format string, v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Fatal><%s>", d.name))

	var output io.Writer = os.Stdout
	if fatalFile != "" {
		file, err := os.OpenFile(fatalFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Panicf(format, v)
}

func (d *sLogger) Fatalln(v ...interface{}) {
	d.l.SetPrefix(fmt.Sprintf("<Fatal><%s>", d.name))

	var output io.Writer = os.Stdout
	if fatalFile != "" {
		file, err := os.OpenFile(fatalFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		confutil.CheckErr(err)

		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}
	d.l.SetOutput(output)
	d.l.Panicln(v)
}
