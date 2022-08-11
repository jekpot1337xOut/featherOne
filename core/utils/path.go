package utils

import (
	"os"
	"path/filepath"
)

var (
	confPath string = "conf/config.yml"
)

func GetConfPath() string {
	return filepath.Join(getRunPath(), confPath)
}

func getRunPath() string {
	pwd, _ := os.Getwd()
	return pwd
}
