package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var xraytime = "2022-02-09"

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
func Cmd(args []byte) <-chan struct{} {
	//case "windows":
	args = bytes.Trim(args, " ")
	argss := strings.Split("/C "+string(args), " ")
	cmd := exec.Command("c:\\windows\\system32\\cmd.exe", argss...)
	closed := make(chan struct{})
	defer close(closed)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() { // 命令在执行的过程中, 实时地获取其输出
			cmdRe := ConvertByte2String(scanner.Bytes(), UTF8) //编码
			fmt.Printf("%s\n", string(cmdRe))
		}
	}()

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return closed
}

func init() {
	log := `
██╗  ██╗██████╗  █████╗ ██╗   ██╗      ██████╗██╗  ██╗███████╗ ██████╗██╗  ██╗
╚██╗██╔╝██╔══██╗██╔══██╗╚██╗ ██╔╝     ██╔════╝██║  ██║██╔════╝██╔════╝██║ ██╔╝
 ╚███╔╝ ██████╔╝███████║ ╚████╔╝      ██║     ███████║█████╗  ██║     █████╔╝ 
 ██╔██╗ ██╔══██╗██╔══██║  ╚██╔╝       ██║     ██╔══██║██╔══╝  ██║     ██╔═██╗ 
██╔╝ ██╗██║  ██║██║  ██║   ██║███████╗╚██████╗██║  ██║███████╗╚██████╗██║  ██╗
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝
                                                                              
																
																				--by:GOGOGO(Need UAC)
------------------------------------------------------------------------------------
`
	fmt.Print(log)
}

func SetTime(settime string) {
	time.Sleep(time.Duration(1) * time.Second)

	Cmd([]byte("date "+settime))
}
func main() {
	args := os.Args
	var xrayargs string
	var xray string
	for x := 1; x < len(args); x++ {
		if len(args) == 1 {
			xrayargs = ""
			break
		}
		xrayargs = xrayargs + " " + args[x]
	}
	settime := time.Now().Format("2006-01-02")
	//fmt.Println(settime)

	Cmd([]byte("date " + xraytime))
	switch runtime.GOOS {
	case "windows":
		xray = "xray.exe"
	//case "darwin":
	//	xray = "./xray"
	//case "linux":
	//	xray = "./xray"
	}

	go SetTime(settime)
	<-Cmd([]byte(xray + xrayargs))
}
