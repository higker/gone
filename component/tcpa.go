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
func TCPA()  {
	utils.Info("开始从远程服务器拉取资源文件~")
	utils.BarDownload(model.Server.Tcpa_package)
	unErr:=utils.Unzip("./temp/tcpa_packets.zip","/usr/local/")
	if unErr != nil {
		utils.Error("安装TCPA出错~请稍后重试!")
		os.Exit(1)
	}
	//cmd := exec.Command("source","/etc/profile")
	//if cmd.Run() != nil {
	//	utils.Error("刷新/etc/profile文件失败!")
	//	os.Exit(1)
	//}
	utils.Done("TCPA配置完成系统需要重启.")
	utils.SleepInfo(10,"秒后重启.")
}