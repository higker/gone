package banner

import "fmt"

/**
 *
 * Create BY YooDing
 *
 * Des: gone banner
 *
 * Time: 2019/7/5 1:50 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github<a>
 */
const (
	Version     = "0.0.5 (beta)"
	Github      = "https://github.com/YooDing/gone"
	Description = "Gone, a multi-functional efficient operation program under Centos written in Go language."
	bannerStr   = "\033[34m ___   __   __ _  ____ \n" +
		" / __) /  \\ (  ( \\(  __)\n" +
		"( (_ \\(  O )/    / ) _) \n" +
		" \\___/ \\__/ \\_)__)(____)\033[0m \n \033[32m version:"+Version+" \033[0m  \r\n "
)
/*print logo*/
func Logo() {
	fmt.Print(bannerStr)
	fmt.Println("\033[42;34m 「welcome using Gone」 \033[0m ")
}
