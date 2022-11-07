package Launcher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func CmdGen(name string, version string, assetsid string, java_path string, work_path string, platform string, arch string, launcherversion string, isolation bool, xmn string, xmx string) {
	java_cmd := Execute_game(version, assetsid, version, name, java_path, work_path, platform, arch, launcherversion, isolation, xmn, xmx)
	cmd := exec.Command("powershell", java_cmd)
	/*
		stdout, err := os.OpenFile("stdout.log", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalln(err)
		}
		defer stdout.Close()
		cmd.Stdout = stdout // 重定向标准输出到文件
		// 执行命令
		if err := cmd.Start(); err != nil {
			log.Println(err)
		}
		fmt.Print("HelloWorld")
	*/

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}
	readerr := bufio.NewReader(stdout)
	go func() {
		GetOutput(readerr)
	}()
	/*
		//读取所有输出
		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			log.Fatalln("ReadAll Stdout:", err.Error())
		}

		fmt.Println(string(bytes))
	*/
}

func GetOutput(reader *bufio.Reader) {
	var sumOutput string //统计屏幕的全部输出内容
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes) //获取屏幕的实时输出(并不是按照回车分割，所以要结合sumOutput)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])
		fmt.Print(output) //输出屏幕内容
		sumOutput += output
	}
	return
}
