/*
	logger定义
*/
package logger

import "log"

type LoggerFactory interface {
	GetLogger(name string) Logger
	SetConfig(path string)
}

type Logger interface {
	//获取logger的主体信息
	Name() string

	//debug信息输出
	Debug(msg string)
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})

	//Info信息输出
	Info(msg string)
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})

	//warn信息输出
	Warn(msg string)
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})

	//错误信息输出
	Error(msg string)
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})

	//崩溃信息输出
	Panic(msg string)
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})

	//程序退出信息输出
	Fatal(msg string)
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}

var f LoggerFactory = nil

//注册logger工厂,暂定工厂类唯一,nil则注册defaultFactory
func RegisterLogger(factory LoggerFactory) {
	if f != nil {
		panic(" registed twice!")
	}
	f = factory
}

//注销loggerFactory
func UnRegisterLogger() {
	f = nil
}

//为loggerFactory 设置配置
func SetConfig(path string) {
	if f == nil {
		panic("no factory registed,can't set config ")

	}

	f.SetConfig(path)
}

//获取logger实例,默认使用os.stdout输出
func GetLogger(name string) Logger {
	if f == nil {
		log.Println("no factory registed,using default LoggerFacotry")
		return new(DefaultLoggerFactory).GetLogger(name)
	}

	return f.GetLogger(name)

}
