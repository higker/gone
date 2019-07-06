package main

import (
	"gone/component"
	"gone/utils"

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
	utils.Info(model.Server.Jdk)
	utils.Done(model.GetJson())
}
