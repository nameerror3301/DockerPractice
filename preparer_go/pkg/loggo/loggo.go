package loggo

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	PKG_APP_LOGGER_WR_FILE_PATH = "default.logger.log"
	FINAL_WRITER_LOG            io.Writer
)

var (
	COLORReset  = "\033[0m"
	COLORRed    = "\033[31m"
	COLORYellow = "\033[33m"
	COLORPurple = "\033[35m"
	COLORCyan   = "\033[36m"
	COLORWhite  = "\033[37m"
	BOLDText    = "\033[1m"
)

type JSONLogFormat struct {
	Level string `json:"level"`
	Mgs   string `json:"msg"`
	Time  string `json:"time"`
}

func InitCastomLogger(format logrus.Formatter, level logrus.Level, file bool, std bool) {
	logrus.SetFormatter(format)
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)

	if file {
		file, err := os.OpenFile(PKG_APP_LOGGER_WR_FILE_PATH, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logrus.Fatal(err.Error())
			panic(err)
		}
		FINAL_WRITER_LOG = io.MultiWriter(file)
	}
	if std {
		if FINAL_WRITER_LOG != nil {
			FINAL_WRITER_LOG = io.MultiWriter(FINAL_WRITER_LOG, &OutConsole{})
		} else {
			FINAL_WRITER_LOG = &OutConsole{}
		}
	}
	if FINAL_WRITER_LOG != nil {
		logrus.SetOutput(FINAL_WRITER_LOG)
	}
}

func AddOut(writer io.Writer) {
	if writer != nil {
		FINAL_WRITER_LOG = io.MultiWriter(FINAL_WRITER_LOG, writer)
		logrus.SetOutput(FINAL_WRITER_LOG)
	}
}

func ChangePathFile(path string) {
	if path == "" {
		return
	}
	PKG_APP_LOGGER_WR_FILE_PATH = path
}
