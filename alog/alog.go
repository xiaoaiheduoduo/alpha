package alog

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/alphaframework/alpha/autil"
	"github.com/alphaframework/alpha/autil/ahttp/request"
	"github.com/alphaframework/alpha/forked/lumberjack"
)

const (
	defaultLogLevel  = "info"
	defaultLogFormat = "console"

	RequestIdKey = "request_id"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

type Config struct {
	ApplicationName  string
	Directory        string
	Level            string
	Format           string
	MaxSize          int // MB
	MaxAge           int // day
	MaxBackups       int
	Compress         bool
	BackupTimeFormat string
}

// InitLogger init Logger and Sugar
// applicationName
// directory: log directory
// level: log level (debug/info/warn/error/panic/fatal)
// format: log format (console/json)
func InitLogger(config *Config) error {
	formatList := []string{"", "console", "json"}
	if !autil.In(config.Format, formatList) {
		return fmt.Errorf("log format: %s does not validate as in %#v", config.Format, formatList)
	}

	if config.Directory == "" {
		return fmt.Errorf("directory is required")
	}

	if config.Level == "" {
		config.Level = defaultLogLevel
	}
	if config.Format == "" {
		config.Format = defaultLogFormat
	}

	// todo check backup time format
	// if config.BackupTimeFormat != "" {
	// 	backupTimeFormat := time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC).Format(config.BackupTimeFormat)
	// 	if config.BackupTimeFormat != backupTimeFormat {
	// 		return fmt.Errorf("log/backup_time_format %q not valid", config.BackupTimeFormat)
	// 	}
	// }

	var l zapcore.Level
	if err := l.Set(config.Level); err != nil {
		return err
	}

	getWriter := func(level string) *lumberjack.Logger {
		var logPath string
		if config.ApplicationName == "" {
			logPath = config.Directory + "/" + level + ".log"
		} else {
			logPath = config.Directory + "/" + config.ApplicationName + "." + level + ".log"

		}
		logger := &lumberjack.Logger{
			Filename:         logPath,
			MaxSize:          config.MaxSize, // MB
			MaxAge:           config.MaxAge,  // day
			MaxBackups:       config.MaxBackups,
			Compress:         config.Compress,
			BackupTimeFormat: config.BackupTimeFormat,
		}
		logger.Complete()
		return logger
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if config.Format == "json" {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var cores []zapcore.Core
	for l <= zapcore.FatalLevel {
		var writers = []zapcore.WriteSyncer{zapcore.AddSync(getWriter(l.String()))}
		cores = append(cores, zapcore.NewCore(
			encoder,
			zapcore.NewMultiWriteSyncer(writers...),
			zap.NewAtomicLevelAt(l),
		))
		l++
	}
	core := zapcore.NewTee(cores...)

	caller := zap.AddCaller()
	development := zap.Development()

	Logger = zap.New(core, caller, development)

	if config.ApplicationName != "" {
		field := zap.Fields(zap.String("application_name", config.ApplicationName))
		Logger.WithOptions(field)
	}

	Logger.Info(`log.InitLogger successfully`)

	Sugar = Logger.Sugar()

	return nil
}

func CtxSugar(ctx context.Context) *zap.SugaredLogger {
	return Ctx(ctx).Sugar()
}

func Ctx(ctx context.Context) *zap.Logger {
	field := zap.Fields(zap.String(RequestIdKey, request.RequestIdValue(ctx)))
	return Logger.WithOptions(field)
}
