package util

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// Cfg 全局cfg 配置
var Cfg *ini.File

// CfgBootstrap 初始化
func CfgBootstrap() {
	// cfg
	iniCfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	Cfg = iniCfg
}
