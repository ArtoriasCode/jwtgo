package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func (s *Logger) ExtraFields(fields map[string]interface{}) *Logger {
	return &Logger{s.WithFields(fields)}
}

var instance Logger
var once sync.Once

func GetLogger(level string) Logger {
	once.Do(func() {
		parsedLevel, err := logrus.ParseLevel(level)
		if err != nil {
			log.Fatalln(err)
		}

		l := logrus.New()
		l.SetReportCaller(true)

		l.Formatter = &logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf(" %s", f.Function)
			},
			DisableQuote:    true,
			DisableColors:   false,
			ForceColors:     true,
			FullTimestamp:   true,
			PadLevelText:    true,
			TimestampFormat: "2006-01-02 15:04:05",
		}

		l.SetOutput(os.Stdout)
		l.SetLevel(parsedLevel)

		instance = Logger{logrus.NewEntry(l)}
	})

	return instance
}
