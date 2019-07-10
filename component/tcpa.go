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
 * Des: Tencent TCPA Data transmission acceleration algorithm
 *
 * Time: 2019/7/10 1:36 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github</a>
 */
func TCPA() {
	utils.Info("开始从远程服务器拉取资源文件~")
	utils.BarDownload(model.Server.Tcpa_package)
	unErr := utils.Unzip("./temp/tcpa_packets.zip", "./temp/")
	if unErr != nil {
		utils.Error("安装TCPA出错~请稍后重试!")
		os.Exit(1)
	}
	cmd1 := utils.ExeShell("sh ./temp/tcpa_packets/install.sh")
	start()
	cmd2 := utils.ExeShell(model.Server.Tcpa_kernel_bash)
	//cmd := exec.Command("source","/etc/profile")
	if cmd1 != nil || cmd2 != nil {
		utils.Error("安装TCPA失败!")
		os.Exit(1)
	}
	utils.Done("TCPA配置完成系统需要重启.可以使用命令:lsmod|grep tcpa 查看是否开启成功.")
	utils.SleepInfo(10, "秒后重启.")
}
func start(){
	utils.ExeShell("sh /usr/local/storage/tcpav2/start.sh")
}