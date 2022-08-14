package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"github.com/phuslu/log"
	"github.com/shuyangzhang/kyouka-light/configs"
	"github.com/shuyangzhang/kyouka-light/handlers"
)

func main() {
	logger := log.Logger{
		Level:  log.InfoLevel,
		Writer: &log.ConsoleWriter{},
	}

	app := khl.New(configs.EnvConfigs.Token, plog.NewLogger(&logger))

	handlers.RegisterHandlers(app)

	app.Open()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc
	app.Close()
}
