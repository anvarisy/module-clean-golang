package modules

import (
	"anvarisy.tech/cleangolang/src/app"
	"anvarisy.tech/cleangolang/src/app/module"
	"anvarisy.tech/cleangolang/src/modules/auth"
)

type Modules struct {
	Auth module.ModuleInterface
}

func Init(app *app.App) *Modules {
	return &Modules{
		Auth: auth.Init(app),
	}
}
