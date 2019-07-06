package component

import (
	"gone/utils"
	"gone/model"
	"os"
)

/**
 *
 * Create BY YooDing
 *
 * Des: Java & jdk
 *
 * Time: 2019/7/6 3:03 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github</a>
 */

const (
	profile = "/etc/profile"
	template = "\r\n export JAVA_HOME=/usr/local/jdk-12.0.1 \r\n export PATH=.:$JAVA_HOME/bin:$PATH \r\n export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar"
)

func Jdk() {
	err := utils.DownloadFile(model.Server.Jdk)
	if err != nil {
		utils.Error("配置JDK出错~请稍后重试!")
		os.Exit(1)
	}
	error := utils.AppendStrToFile(profile,template)
	if error != nil {
		utils.Error("配置JDK出错~请稍后重试!")
		os.Exit(1)
	}
}
