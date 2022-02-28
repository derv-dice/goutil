package main

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Log() *zerolog.Logger {
	return &logger
}

const defaultLogsDir = "./logs"

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

// InitLogs - Настройка логирования
// pretty - Вывод лога в stdout в формате '8:37PM INF Message'
// colored - работает только в связке с pretty и делает лог цветным или монохромным
// file - Писать ли логи в файл. Если true, то создается папка ./logs и в нее пишется тот же лог, что и в stdout,
//	но в формате json для более удобного парсинга логов, если понадобится
func InitLogs(pretty, colored, file bool, logsDir string) (err error) {
	multiWR := zerolog.MultiLevelWriter(os.Stdout)
	var f io.Writer

	if file {
		if logsDir == "" {
			logsDir = defaultLogsDir
		}

		if _, err = os.Stat(logsDir); os.IsNotExist(err) {
			if err = os.MkdirAll(logsDir, 0777); err != nil {
				return
			}
		}

		if f, err = os.Create(path.Join(logsDir, time.Now().Format("2006-01-02_15:04:05.log"))); err != nil {
			return
		}

		multiWR = zerolog.MultiLevelWriter(os.Stdout, f)
		logger = log.Output(multiWR)
	}

	if pretty {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

		if !colored {
			consoleWriter.NoColor = true
		}

		if f == nil {
			multiWR = zerolog.MultiLevelWriter(consoleWriter)
		} else {
			multiWR = zerolog.MultiLevelWriter(consoleWriter, f)
		}

		logger = zerolog.New(multiWR).With().Timestamp().Logger()
	}

	return
}
