/*
	简单的解析conf文件，基本格式如下
	[se1]
	keya1=valuea
	keyb2=valueb
	[se2]
	keya2=valuea2
	keyb2=valueb2
	...

*/
package confutil

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Config struct {
	Secs map[string]map[string]string
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//针对的是结尾是否有空格
func (c *Config) Anly(path string) {
	c.Secs = make(map[string]map[string]string)
	var sec map[string]string
	f, err := os.Open(path)
	CheckErr(err)

	defer f.Close()
	r := bufio.NewReader(f)

	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		str := strings.TrimSpace(string(b))
		//忽视注释
		if strings.Index(str, "#") == 0 {
			continue
		}

		n1 := strings.Index(str, "[")
		n2 := strings.Index(str, "]")
		//这边默认没有空的sec.没有sec则指定为default
		if n1 != -1 && n2 > n1+1 {
			secName := str[n1+1 : n2]
			sec = make(map[string]string)
			c.Secs[secName] = sec
			continue
		}

		index := strings.Index(str, "=")
		if index != -1 {
			if len(c.Secs) == 0 {
				sec = make(map[string]string)
				c.Secs["default"] = sec
			}
			key := strings.TrimSpace(str[:index])
			value := strings.TrimSpace(str[index+1:])
			sec[key] = value
			continue
		}
	}
}
