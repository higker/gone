package model

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	"fmt"

)

/**
 *
 * Create BY YooDing
 *
 * Des: 远程配置文件json数据
 *
 * Time: 2019/7/6 11:19 AM.
 *
 * <a href="https://github.com/YooDing/gone">Github</a>
 */
//
//type ZipUrl struct {
//	ZipName string
//	ZipPath string
//}

type ConfigData struct {
	Version string
	Jdk string
	tomcat string
}

var(
	res *simplejson.Json
)

func init(){
	str := `{
		"version": "0.0.6",
		"jdk1.8": "www.qq.com",
        "tomcat": "www.baidu.com"
	}`
	buf := bytes.NewBuffer([]byte(str))
	js, _ := simplejson.NewFromReader(buf)
	fmt.Println(js.Get("version").String())
	res, _ = simplejson.NewJson([]byte(str))
}
