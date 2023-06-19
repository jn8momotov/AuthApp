package main

import "AuthApp/internal/app"

const configDir = "configs"

func main() {
	app.Run(configDir)
}
