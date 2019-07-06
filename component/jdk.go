package component

import (
	"gone/utils"
	"gone/model"
	"os"
	"os/exec"
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
	template = "export JAVA_HOME=/usr/local/jdk-12.0.1 export PATH=.:$JAVA_HOME/bin:$PATH export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar"
)

func JDK() {
	utils.Info("开始从远程服务器拉取资源文件:"+model.Server.Jdk)
	err := utils.DownloadFile(model.Server.Jdk)
	if err != nil {
		utils.Error("下载JDK安装包出错~请稍后重试!")
		os.Exit(1)
	}
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
	cmdErr := exec.Command("source","/etc/profile").Start()
	if cmdErr != nil {
		utils.Error("刷新/etc/profile文件失败!请使用你手动刷新:source /etc/profile")
		os.Exit(1)
	}
	utils.Done("JDK配置完成,请使用java -version查看相关信息.")
}
