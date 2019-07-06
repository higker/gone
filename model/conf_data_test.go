package model

import (
	"fmt"
	"encoding/json"
)

/**
 *
 * Create BY YooDing
 *
 * Des: test01
 *
 * Time: 2019/7/6 11:31 AM.
 *
 * <a href="https://github.com/YooDing/gone">Github</a>
 */
func main() {
	var c ConfigData
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
