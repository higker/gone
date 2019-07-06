package component

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
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
func Menus() {

	fmt.Println("\n 输入数字选择功能：\n")
	fmt.Println(" 1 - 安装JDK \n")
	fmt.Println(" 2 - 安装Tomcat \n")
	fmt.Println(" 3 - 加速算法 \n")
	fmt.Println(" 0 - 退出程序 \n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("选择功能: ")
	text, _ := reader.ReadString('\n')
	err := unicode.IsDigit([]rune(text)[0])

	if err != true {
		utils.Warning("请输入正确序号！1")
		Menus()
	}
	switch text {
	case "1":
		fmt.Println(text)
		JDK()
	case "2":
		fmt.Println(text)
	default:
		fmt.Println(text)
		utils.Warning("请输入正确序号！")
		Menus()
	}

}
