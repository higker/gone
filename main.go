package main

import (
	"gone/component"
	"gone/utils"
	"fmt"
	"bytes"
	"github.com/bitly/go-simplejson"
)

/**
 *
 * Create BY YooDing
 *
 * Des: application main function
 *
 * Time: 2019年07月05日13:28:19
 *
 * <a href="https://github.com/YooDing/gone">Github<a>
 */
func main() {
	component.Menus()

	//utils.Done("完毕")
	//utils.Info("信息")
	utils.Error("错误")
	str := `{
  		"version":"0.0.6",
		"zipurl":[
			{"zipname":"jdk1.8","zippath":"www.qq.com"},
    		{"zipname":"tomcat","zippath":"www.baidu.com"}
  		]
	}`
	buf := bytes.NewBuffer([]byte(str))
	js, _ := simplejson.NewFromReader(buf)
	fmt.Println(js.Get("version").String())

	res, _ := simplejson.NewJson([]byte(str))
	fmt.Println(res.Get("tomcat").String())
}
