package main

import (
	"fmt"
	"http_server/config"
	"http_server/server"
)


func main() {
	config := config.LoadConfig()
	fmt.Println(config.Port)
	fmt.Println(config.JWTSecret)
	fmt.Println(config.DatabaseURL)

	app := server.NewApp()
	app.RunServer(config.Port)
}
