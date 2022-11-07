package Launcher

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Checker(pwd string) []string {
	var Versions []string
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		fmt.Print(err)
	}
	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			if Exists(".minecraft\\versions\\" + fileInfoList[i].Name() + "/" + fileInfoList[i].Name() + ".json") {
				Versions = append(Versions, fileInfoList[i].Name())
			}
		}
	}
	return Versions
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if os.IsNotExist(err) {
		return false
	}
	return true
}
