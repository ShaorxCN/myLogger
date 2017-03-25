**定义logger接口，简单实现logger即sLogger**

可以根据接口自己实现logger,sLogger是这边提供的一个实现示例,,sLogger通过读取配置文件实现控制台以及文件的日志输出,未设置文件路径的则只会在控制台输出，sLogger在并发下需要用户自己保证其正确性

使用示例
```go
   package main

	import (
		"github.com/ShaorxCN/myLogger/logger"
		_ "github.com/ShaorxCN/myLogger/sLogger"
	)

	func main() {
		/*设置配置文件路径，配置文件指定了各类文件日志的路径
		  暂时提供如下文件标识warnFile, debugFile, infoFile, errorFile, panicFile, fatalFile
		*/
		logger.SetConfig("config")

		//获取logger
		l := logger.GetLogger("test")

		l.Debugln("debug")
		l.Errorln("error")
		l.Infoln("info")
		l.Warnln("warn")

	}
```

config 文件如下,这里仅配置了errorFile,infoFile
```
errorFile=errfile
infoFile=infoFile
```

代码运行后会在控制台输出
```
<Debug><test>2017/03/25 13:56:50 [debug]
<Error><test>2017/03/25 13:56:50 [error]
<Info><test>2017/03/25 13:56:50 [info]
<Warn><test>2017/03/25 13:56:50 [warn]
```
同时会在指定目录下创建errFile以及infoFile(这里就是在当前目录),写入对应的内容

errfile内容如下
`<Error><test>2017/03/25 13:56:50 [error]`
infoFile内容如下
`<Info><test>2017/03/25 13:56:50 [info]`