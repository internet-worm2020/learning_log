package logging

import (
	setting "gindome/config"
	"github.com/internet-worm2020/go-pkg/log"
)

func Adcc() *log.Options {
	opts := &log.Options{
		Level:            setting.Conf.LogConfig.Level,
		Format:           "json",
		EnableColor:      false, // if you need output to local path, with EnableColor must be false.
		DisableCaller:    true,
		OutputPaths:      []string{setting.Conf.LogConfig.LogFilePath + setting.Conf.LogConfig.WebLogName, "stdout"},
		ErrorOutputPaths: []string{setting.Conf.LogConfig.LogFilePath + setting.Conf.LogConfig.WebLogErrorName, "stdout"},
	}
	return opts
}
