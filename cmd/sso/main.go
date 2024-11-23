package main

import (
	"fmt"

	"github.com/cyrillus31/go_auth_server/internal/config"
)

func main() {
	// TODO: инициализировать объект конфига
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложоения-сервис (app)

	// TODO: запустить gRPC сервер
}
