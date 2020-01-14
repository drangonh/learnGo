package utils

import (
	"os"
)

//获取go env文件路径
func EnvPath() string {
	dir, _ := os.UserConfigDir()
	return dir
}
