package config

import (
	"errors"
	"log"
	"os"

	"github.com/LuccChagas/clean-scaffold/validation"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type BaseConfig struct {
	log  *zap.Logger
	test string
}

const (
	LOCAL = "/.env"
	PROD  = "/.env.prod"
)

type EnvTypes struct {
	Environment string `json:"local" validate:"required,oneof=local prod"`
}

func SetupEnv() (string, error) {
	env := EnvTypes{
		Environment: os.Args[1],
	}

	err := validation.Validate(env)
	if err != nil {
		err = errors.New("wrong environment") // tratar melhor resposta
		return "Invalid: ", err
	}

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	switch env.Environment {
	case "local":
		if err := godotenv.Load(curDir + LOCAL); err != nil {
			log.Fatalln("can't load env file from current directory: " + curDir)
		}

	case "prod":
		if err := godotenv.Load(curDir + PROD); err != nil {
			log.Fatalln("can't load env file from current directory: " + curDir)
		}

	default:
		err = errors.New("wrong environment") // tratar melhor resposta
		return "Invalid: ", err
	}

	return env.Environment, nil
}
