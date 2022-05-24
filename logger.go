package goutil

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
)

func Log() *zerolog.Logger {
	return &logger
}

const defaultLogsDir = "./logs"

const (
	LevelTrace = "trace"
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelFatal = "fatal"
	LevelPanic = "panic"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

// InitLoggerConf - Конфигуратор логов
// - Level: Уровень логов <trace | debug | info | warn | error | fatal | panic>. По умолчанию - info
// - Pretty: В stdout будет выводиться лог формате '1:01PM LEVEL MESSAGE'
// - Colored: В stdout будет выводиться цветной лог (Работает только если Pretty=true)
// - ToFile: Логировать ли в файл
// - Dir: В какую директорию положить файл с логами (Работает только если ToFile=true). По умолчанию - ./logs/
// - CodeLine: Логирование строки кода
type InitLoggerConf struct {
	Level    string
	Pretty   bool
	Colored  bool
	ToFile   bool
	Dir      string
	CodeLine bool
}

func InitLogger(conf InitLoggerConf) (err error) {
	var writers []io.Writer

	if conf.ToFile {
		if conf.Dir == "" {
			conf.Dir = defaultLogsDir
		}

		if _, err = os.Stat(conf.Dir); os.IsNotExist(err) {
			if err = os.MkdirAll(conf.Dir, 0775); err != nil {
				return
			}
		}

		var f io.Writer
		if f, err = os.Create(path.Join(conf.Dir, time.Now().Format("2006-01-02_15:04:05.log"))); err != nil {
			return
		}

		writers = append(writers, f)
	}

	if conf.Pretty {
		cw := zerolog.ConsoleWriter{Out: os.Stdout}
		if !conf.Colored {
			cw.NoColor = true
		}

		writers = append(writers, cw)
	} else {
		writers = append(writers, os.Stdout)
	}

	var lvl zerolog.Level
	if conf.Level == "" {
		lvl = zerolog.InfoLevel
	} else {
		if lvl, err = zerolog.ParseLevel(conf.Level); err != nil {
			return
		}
	}

	logger = zerolog.New(zerolog.MultiLevelWriter(writers...)).Level(lvl).With().Timestamp().Logger()

	if conf.CodeLine {
		logger = logger.With().Caller().Logger()
	}

	return
}
