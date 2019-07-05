package utils

import (
	"fmt"
	"time"
	"strconv"
)

/**
 *
 * Create BY YooDing
 *
 * Des: system util
 *
 * Time: 2019年07月05日13:32:44
 *
 * <a href="https://github.com/YooDing/gone">Github<a>
 */
func Done(logStr string) {
	fmt.Printf("\033[42;34m 「Done」 \033[0m \033[32m %s --> %s \033[0m \n", NowTimeStr(), logStr)
}
func Info(logStr string) {
	fmt.Printf("\033[44;30m 「Info」 \033[0m \033[34m %s --> %s \033[0m \n", NowTimeStr(), logStr)
}
func Error(logStr string) {
	fmt.Printf("\033[41;30m 「Fail」 \033[0m \033[31m %s --> %s \033[0m \n", NowTimeStr(), logStr)
	logError.Print(logStr)
}


func NowTimeStr() (timeStr string) {
	now := time.Now()
	return strconv.Itoa(now.Year()) + "-" + now.Month().String() + "-" + strconv.Itoa(now.Day()) + " " + strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute()) + ":" + strconv.Itoa(now.Second())
}
