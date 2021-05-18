package logger

import (
	"fmt"
	"github.com/mr-chelyshkin/go-scrumble/pkg/errors"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"syscall"
)

/*
	Default configs values
*/
var fileName    = "service.log"
var fileAgeDays = 30
var fileBackups = 1
var fileSize    = 5

/*
	Define internal errors
*/
const (
	permDirErr    = "user doesn't have permission to write to this directory"
	permFileErr   = "write permission bit is not set on this file for user"
	existErr      = "log file path doesn't exist"
	directoryErr  = "path is not a directory"
	statErr       = "unable to get stat"
)

/*
	Prepare Config for Logger object
	concatenate customs configs with base and return filled object even income config is nil
*/
func prepareConfig(config *Config) Config {
	loggerConfig := Config{
		FileDirectory:  "",
		FileName:       fileName,

		FileBackups:    fileBackups,
		FileSize:       fileSize,
		FileAge:        fileAgeDays,

		Compress:       false,
	}

	if config == nil {
		return loggerConfig
	}

	switch {
	case config.FileName != "":
		loggerConfig.FileName = config.FileName
	case config.FileDirectory != "":
		loggerConfig.FileDirectory = config.FileDirectory
	case config.FileSize != 0:
		loggerConfig.FileSize = config.FileSize
	case config.FileAge != 0:
		loggerConfig.FileAge = config.FileAge
	case config.FileBackups != 0:
		loggerConfig.FileBackups = config.FileBackups
	case config.Compress:
		loggerConfig.Compress = config.Compress
	}

	return loggerConfig
}

/*
	Check directory for writing *.log files
*/
func isWritable(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, errors.Core.NewF(err, existErr)
	}

	if !info.IsDir() {
		return false, errors.Core.New(fmt.Errorf(directoryErr))
	}

	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false, errors.Core.New(fmt.Errorf(permFileErr))
	}

	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false, errors.Core.NewF(err, statErr)
	}

	if uint32(os.Geteuid()) != stat.Uid {
		return false, errors.Core.New(fmt.Errorf(permDirErr))
	}

	return true, nil
}

/*
	Crete sync logger object and *.log files
*/
func syncToFile(config Config) zapcore.WriteSyncer {
	lumberLog := &lumberjack.Logger{
		Filename:   path.Join(config.FileDirectory, config.FileName),
		Compress:   config.Compress,
		MaxSize:    config.FileSize,
		MaxAge:     config.FileAge,
		MaxBackups: config.FileBackups,
	}
	return zapcore.AddSync(lumberLog)
}

/*
	Crete sync logger object and stdout
*/
func syncToStdout() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
