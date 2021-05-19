package logger

import (
	"go-scrumble/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	Scrumble logger object
	based from: "go.uber.org/zap"
*/
type Logger struct {
	config Config
	zap    *zap.Logger
}

/*
	Initialize scrumble logger object or return error
	use Config for configure logger or nil for default
*/
func Create(config *Config) (*Logger, error) {
	loggerConfig := prepareConfig(config)

	// chose place for writing logs: stdout / logfile
	var syncer zapcore.WriteSyncer
	if loggerConfig.FileDirectory != "" {
		// check logs directory
		if ok, err := isWritable(loggerConfig.FileDirectory); !ok {
			return nil, errors.Wrap(err, "check log filepath from configs")
		}

		syncer = syncToFile(loggerConfig)
	} else {
		syncer = syncToStdout()
	}
	syncZapLumber := zapcore.AddSync(syncer)

	// initialize zap.Logger: https://github.com/uber-go/zap
	cfg := zap.NewProductionEncoderConfig()

	if loggerConfig.ColorLevel == false {
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if loggerConfig.TimeFormat != "" {
		cfg.EncodeTime = zapcore.TimeEncoderOfLayout(loggerConfig.TimeFormat)
	} else {
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// initialize zap.Core
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		syncZapLumber,
		zap.InfoLevel,
	)

	return &Logger{
		config: loggerConfig,
		zap:    zap.New(core),
	}, nil
}
