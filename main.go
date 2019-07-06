package main

import (
	"gone/component"
	"gone/utils"
	"encoding/json"
	"fmt"
	"gone/model"
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
	var c model.ConfigData
	str := `{
  "Version":"0.0.6",
  "Zipurl":[
    {"ZipName":"jdk1.8","ZipPath":"www.qq.com"},
    {"ZipName":"tomcat","ZipPath":"www.baidu.com"}
  ]
}`
	json.Unmarshal([]byte(str), &c)
	fmt.Println(c)
	for key, val := range c.ZipUrl {
		println(key, "：")
		fmt.Println("Name：", val.ZipName)
		fmt.Println("Path：", val.ZipPath)
		println()
	}

}
