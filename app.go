package main

import (
	app "github.com/sanjeyac/gokickstart/internal/app"
)

func main() {
	server := new(app.Server)
	server.Init()
	server.RegisterRoutes()
	server.Start()
}
