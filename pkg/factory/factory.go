package factory

import (
	"github.com/MirableOne/word-of-wisdom/pkg/handler"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
	"github.com/MirableOne/word-of-wisdom/pkg/server"
)

type serviceContainer struct {
	log *logger.Logger
}

var container serviceContainer

func MakeLogger() *logger.Logger {
	if container.log == nil {
		container.log = logger.NewLogger()
	}

	return container.log
}

func MakeHandler() server.Handler {
	return handler.NewHandler(&handler.Config{
		Log: MakeLogger(),
	})
}
