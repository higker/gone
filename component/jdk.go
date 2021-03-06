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
	template = "\n#this using gone configuration\n#https://github.com/YooDing/gone\nexport JAVA_HOME=/usr/local/jdk-12.0.1\nexport PATH=.:$JAVA_HOME/bin:$PATH\nexport CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar\n"
)

func JDK() {
	utils.Info("开始从远程服务器拉取资源文件:"+model.Server.Jdk)
	utils.BarDownload(model.Server.Jdk)
	unErr:=utils.Unzip("./temp/jdk-12.0.1.zip","/usr/local/")
	if unErr != nil {
		utils.Error("解压JDK出错~请稍后重试!")
		os.Exit(1)
	}
	error := utils.AppendStrToFile(profile,template)
	if error != nil {
		utils.Error("配置JDK出错~请稍后重试!")
		os.Exit(1)
	}
	cmd := utils.ExeShell("source /etc/profile")
	// := exec.Command("source","/etc/profile")
	if cmd != nil {
		utils.Error("刷新/etc/profile文件失败!")
		//os.Exit(1)
	}
	utils.Done("/nJDK配置完成,请使用你手动刷新:source /etc/profile,请使用java -version查看相关信息.")
}
