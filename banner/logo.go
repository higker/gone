package banner

import "fmt"

/**
 *
 * Create BY YooDing
 *
 * Des: gone application main function
 *
 * Time: 2019/7/5 1:50 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github<a>
 */
const (
	version     = "0.0.5"
	github      = "https://github.com/YooDing/gone"
	description = "Gone, a multi-functional efficient operation program under Centos written in Go language."
	bannerStr   = "\033[34m ___   __   __ _  ____ \n" +
		" / __) /  \\ (  ( \\(  __)\n" +
		"( (_ \\(  O )/    / ) _) \n" +
		" \\___/ \\__/ \\_)__)(____)\033[0m \n "
)

func Logo() {
	fmt.Print(bannerStr)
	fmt.Printf("\033[42;34m 「welcome using Gone」 \033[0m \033[32m version:%s  \033[0m \n", version)
}
