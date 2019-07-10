package main

import (
	"gone/utils"
	"runtime"
	"gone/model"
	"os"
	"gone/component"
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

	utils.Info(model.Server.Jdk)


	//utils.SleepInfo(20,"秒后系统自动重启")

	//utils.BarDownload("")


}

func init() {
	banner.Logo()
	component.Menus()
	os.Chdir("./logs")
	os.Chdir("./temp")
}

func checkSys() {
	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		// LINUX系统
		utils.Done("OK")
	} else {
		utils.Error("NO")
	}
}
