package main

import (
	"fmt"
	"os"

	mainapp "anvarisy.tech/cleangolang/src/app"
	"anvarisy.tech/cleangolang/src/modules"
)

func main() {
	var (
		app    = mainapp.Init()
		module = modules.Init(app)
		router = app.GetHttpRouter()
	)
	router.Mount("/auth", module.Auth.GetHttpRouter())
	var err = app.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
