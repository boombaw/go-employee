package main

import (
	database "go-employee/pkg/v1/database/mysql"
	"go-employee/pkg/v1/routers"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	_ = godotenv.Load(".env")

}

func main() {
	log.Debug().Msg("App")
	database.Open()
	routers.Routes()

}
