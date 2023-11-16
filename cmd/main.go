package main

import (
	"stock-services/envs"
	"stock-services/internal/app"
)

func init() {
	envs.LoadEnvs()
}

func main() {
	app.Run()
}
