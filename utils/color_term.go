package utils

import (
	"reflect"
	"fmt"
	"strings"
)

/**
 *
 * Create BY YooDing
 *
 * Des: ColorTerm util
 *
 * Time: 2019/7/9 8:30 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github</a>
 */
//绿色字体，modifier里，第一个控制闪烁，第二个控制下划线
func Green(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 0, modifier...)
}

//淡绿
func LightGreen(str string, modifier ...interface{}) string {
	return cliColorRender(str, 32, 1, modifier...)
}

//青色/蓝绿色
func Cyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 0, modifier...)
}

//淡青色
func LightCyan(str string, modifier ...interface{}) string {
	return cliColorRender(str, 36, 1, modifier...)
}

//红字体
func Red(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 0, modifier...)
}

//淡红色
func LightRed(str string, modifier ...interface{}) string {
	return cliColorRender(str, 31, 1, modifier...)
}

//黄色字体
func Yellow(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0, modifier...)
}

//黑色
func Black(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 0, modifier...)
}

//深灰色
func DarkGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 30, 1, modifier...)
}

//浅灰色
func LightGray(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 0, modifier...)
}

//白色
func White(str string, modifier ...interface{}) string {
	return cliColorRender(str, 37, 1, modifier...)
}

//蓝色
func Blue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 0, modifier...)
}

//淡蓝
func LightBlue(str string, modifier ...interface{}) string {
	return cliColorRender(str, 34, 1, modifier...)
}

//紫色
func Purple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 0, modifier...)
}

//淡紫色
func LightPurple(str string, modifier ...interface{}) string {
	return cliColorRender(str, 35, 1, modifier...)
}

//棕色
func Brown(str string, modifier ...interface{}) string {
	return cliColorRender(str, 33, 0, modifier...)
}

func cliColorRender(str string, color int, weight int, extraArgs ...interface{}) string {
	//闪烁效果
	var isBlink int64 = 0
	if len(extraArgs) > 0 {
		isBlink = reflect.ValueOf(extraArgs[0]).Int()
	}
	//下划线效果
	var isUnderLine int64 = 0
	if len(extraArgs) > 1 {
		isUnderLine = reflect.ValueOf(extraArgs[1]).Int()
	}
	var mo []string
	if isBlink > 0 {
		mo = append(mo, "05")
	}
	if isUnderLine > 0 {
		mo = append(mo, "04")
	}
	if weight > 0 {
		mo = append(mo, fmt.Sprintf("%d", weight))
	}
	if len(mo) <= 0 {
		mo = append(mo, "0")
	}
	return fmt.Sprintf("\033[%s;%dm"+str+"\033[0m", strings.Join(mo, ";"), color)
}