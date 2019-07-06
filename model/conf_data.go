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
package model

import (
	"github.com/bitly/go-simplejson"
	"gone/utils"
	"os"

)
//
//type ZipUrl struct {
//	ZipName string
//	ZipPath string
//}
type ConfigData struct {
	Version string
	Jdk     string
	tomcat  string
}

var (
	res    *simplejson.Json
	Server ConfigData
)

func init() {
	str := `{
		"version": "0.0.6",
		"jdk": "www.qq.com",
        "tomcat": "www.baidu.com"
	}`
	//buf := bytes.NewBuffer([]byte(str))
	//js, _ := simplejson.NewFromReader(buf)
	//fmt.Println(js.Get("version").String())
	res, _ = simplejson.NewJson([]byte(str))
	version, _ := res.Get("version").String()
	jdk, _ := res.Get("jdk").String()
	tomcat, _ := res.Get("tomcat").String()
	Server = ConfigData{version, jdk, tomcat}
}

func parsing(jsonStr string) ConfigData {
	res, _ = simplejson.NewJson([]byte(jsonStr))
	version, err := res.Get("version").String()
	jdk, _ := res.Get("jdk").String()
	tomcat, _ := res.Get("tomcat").String()
	if err != nil {
		utils.Error("程序解析远程仓库配置json文件失败！！！！请你稍后重试！")
		//发生错误程序退出
		os.Exit(1)
	}
	return ConfigData{version, jdk, tomcat}
}
