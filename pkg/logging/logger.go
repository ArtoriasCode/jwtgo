package logging

import (
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func (l *Logger) ExtraFields(fields map[string]interface{}) *Logger {
	zapFields := make([]interface{}, 0, len(fields)*2)

	for k, v := range fields {
		zapFields = append(zapFields, k, v)
	}

	return &Logger{l.With(zapFields...)}
}

var (
	instance Logger
	once     sync.Once
)

func GetLogger(level string) Logger {
	once.Do(func() {
		zapLevel := parseLevel(level)

		encoderCfg := zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			NameKey:      "logger",
			CallerKey:    "caller",
			MessageKey:   "msg",
			FunctionKey:  zapcore.OmitKey,
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeCaller: zapcore.ShortCallerEncoder,
		}

		consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

		core := zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			zapLevel,
		)

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		instance = Logger{logger.Sugar()}
	})

	return instance
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}
