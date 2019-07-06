package component

import (
	"fmt"
	"gone/utils"
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
var (
	input string
)

func Menus() {

	fmt.Println("\n 输入数字选择功能：\n")
	fmt.Println(" 1 - 安装JDK \n")
	fmt.Println(" 2 - 安装Tomcat \n")
	fmt.Println(" 3 - 加速算法 \n")
	fmt.Println(" 0 - 退出程序 \n")
	fmt.Print("选择功能: ")

	fmt.Scanln(&input)
	switch input {
	case "1":
		JDK()
	case "2":
		fmt.Println("tomcat")
	default:
		fmt.Println(input)
		utils.Warning("请输入正确序号！")
		Menus()
	}

}
