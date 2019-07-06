package component

import (
	"gone/banner"
	"fmt"
)

/**
 *
 * Create BY YooDing
 *
 * Des:  application console menus
 *
 * Time: 2019/7/5 8:11 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github<a>
 */
func Menus() {
	banner.Logo()
	fmt.Println("\n 输入数字选择功能：\n")
	fmt.Println(" 1 - 软件列表 \n")
	fmt.Println(" 2 - 常用工具 \n")
	fmt.Println(" 3 - 加速算法 \n")
	fmt.Println(" 0 - 退出程序 \n")
	fmt.Println(" 输入数字选择功能：")
}
