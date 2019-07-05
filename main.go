package main

import (
	"fmt"
	"gone/utils"
	"gone/banner"
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
	banner.Logo()
	fmt.Println("Hello gone! 还没有开始动工写哟~~")
	utils.Done("完毕")
	utils.Info("信息")
	utils.Error("错误")
}
