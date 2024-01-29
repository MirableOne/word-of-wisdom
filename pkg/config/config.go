package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/go-playground/validator/v10"
)

func Extract(cfg interface{}) error {
	return env.Parse(cfg)
}

func Validate(cfg interface{}) error {
	validate := validator.New()

	return validate.Struct(cfg)
}

func Configure(cfg interface{}) error {
	err := Extract(cfg)

	if err != nil {
		return err
	}

	return Validate(cfg)
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
