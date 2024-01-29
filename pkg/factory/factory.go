package factory

import (
	"github.com/MirableOne/word-of-wisdom/pkg/handler"
	"github.com/MirableOne/word-of-wisdom/pkg/hashcash"
	"github.com/MirableOne/word-of-wisdom/pkg/logger"
	"github.com/MirableOne/word-of-wisdom/pkg/server"
	"github.com/MirableOne/word-of-wisdom/pkg/storage"
)

type serviceContainer struct {
	log       *logger.Logger
	hashStore *storage.InMemoryMapStorage
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
		Log:     MakeLogger(),
		Storage: MakeStorage(),
	})
}

func MakeStorage() hashcash.Storage {
	if container.hashStore == nil {
		container.hashStore = storage.NewInMemoryMapStorage()
	}

	return container.hashStore
}
