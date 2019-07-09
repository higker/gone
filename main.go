package main

import (
	"gone/utils"
	"runtime"
	"gone/model"
	"os"
	"fmt"
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

	err, out, _ :=  utils.Shellout("who")
	if err != nil {
		utils.Error("失败!")
		os.Exit(1)
	}else {
		fmt.Println(out)
	}

}

func init() {
	//os.Create("./logs/")
	//os.Create("./temp")
	//banner.Logo()
	//component.Menus()
}

func checkSys() {
	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		// LINUX系统
		utils.Done("OK")
	} else {
		utils.Error("NO")
	}
}
