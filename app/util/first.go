package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

// FirstIni
var FirstIni *ini.File

// 启动检测
func FirstRunChecking() {

	// 读取文件
	_, err := ioutil.ReadFile("first.lock")

	// 创建文件
	if err != nil {

		// 创建
		f, err := os.Create("first.lock")
		if err != nil {
			log.Fatal("first.lock create fail")
		}

		// 写入
		_, err = io.WriteString(f, "created_at = "+time.Now().String())
		if err != nil {
			log.Fatal("first.lock write fail")
		}
	}

	// read
	firstIni, err := ini.Load("first.lock")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	FirstIni = firstIni

}
