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
	"net/http"
	"io/ioutil"
)

//
//type ZipUrl struct {
//	ZipName string
//	ZipPath string
//}
type ConfigData struct {
	Version string
	Jdk     string
	Tomcat  string
	Tcpa_package string
	Tcpa_kernel_bash string
}

var (
	res    *simplejson.Json
	Server ConfigData
)

func init() {
	//init model
	Server = parsing(getJson())
}
func parsing(jsonStr string) ConfigData {
	res, _ = simplejson.NewJson([]byte(jsonStr))
	version, err := res.Get("version").String()
	jdk, _ := res.Get("jdk").String()
	tomcat, _ := res.Get("tomcat").String()
	tcpa_package, _ := res.Get("tcpa_package").String()
	tcpa_kernel_bash, _ := res.Get("tcpa_kernel_bash").String()
	if err != nil {
		utils.Error("程序解析远程仓库配置json文件失败！！！！请你稍后重试！")
		//发生错误程序退出
		os.Exit(1)
	}
	return ConfigData{version, jdk, tomcat,tcpa_package,tcpa_kernel_bash}
}
func getJson() string {
	//send get request
	resp, err := http.Get("https://raw.githubusercontent.com/YooDing/gone/master/config.json")
	if err != nil {
		utils.Error("连接远程服务器失败!请你稍后重试！")
		//发生错误程序退出
		os.Exit(1)
	}
	//The last execution
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Error("解析最新资源失败!请你稍后重试！")
		//发生错误程序退出
		os.Exit(1)
	}
	return string(body)
}
