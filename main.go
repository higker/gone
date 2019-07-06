package main

import (
	"gone/utils"
	"runtime"
	"gone/model"
	"gone/component"
	"gone/banner"
	"os"
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

	//utils.Done("完毕")
	//utils.Info("信息")
	//utils.Error("错误")
	utils.Info(model.Server.Jdk)



}

func init() {
	os.Create("./logs/")
	os.Create("./temp")
	//str := "\r\n export JAVA_HOME=/usr/local/jdk1.8.0_144 \r\n export PATH=.:$JAVA_HOME/bin:$PATH \r\n export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar"
	banner.Logo()
	component.Menus()
	//utils.AppendStrToFile("/Users/jdode/Desktop/login.jsp",str)
	//utils.DownloadFile("https://m7.music.126.net/20190706150200/8154e950dde9990b0a1c53e524d968cc/ymusic/0f5e/070b/0e5e/84f91eaa3f0a04eba0a7944662ed4af5.mp3")
	//checkSys()
}

func checkSys() {
	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		// LINUX系统
		utils.Done("OK")
	} else {
		utils.Error("NO")
	}
}
