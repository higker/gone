/**
 *
 * Create BY YooDing
 *
 * Des: record log file
 *
 * Time: 2019/7/5 7:22 PM.
 *
 * <a href="https://github.com/YooDing/gone">Github<a>
 */
package utils

import (
	"log"
	"os"
	"io"
)


var (
	//LogInfo *log.Logger
	//LogWarning *log.Logger
	logError * log.Logger
)

func init(){
	errFile,err:=os.OpenFile("./logs/errors_"+NowTimeStr()+".log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("Error writing log file :",err)
	}
	//LogInfo = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	//LogWarning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	//logError = log.New(io.MultiWriter(os.Stderr,errFile),"Error : ",log.Ldate | log.Ltime | log.Lshortfile)
	logError = log.New(io.MultiWriter(errFile),"Error : ",log.Ldate | log.Ltime | log.Lshortfile)
}
//
//func getErrInput() *log.Logger {
//	return  logError
//}