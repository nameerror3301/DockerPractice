// Базовый шаблон.
// Для быстрого развертывания пустого проекта.
// Сервис в своей работе использует переменные глобального окружения
//
//   - DRIVER_CONFIG={json,yaml,toml}	формат кофигурационного файла.
//     При отсутствии параметров по умолчанию установлен JSON.
package main

import (
	"context"
	"os"
	"os/signal"
	"store/internal/service"
	"store/pkg/loggo"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	var quit = make(chan os.Signal, 1)
	osKillSignals := []os.Signal{
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGINT,
	}
	signal.Notify(quit, osKillSignals...)

	ctx := context.Background()

	go runApp(ctx)

	<-quit
	ctx.Done()
}

func runApp(ctx context.Context) {
	loggo.InitCastomLogger(&logrus.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logrus.TraceLevel, false, true)

	if err := service.Service(ctx); err != nil {
		logrus.Error(err.Error())
	}
}
