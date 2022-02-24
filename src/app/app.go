package app

import (
	"database/sql"
	"net/http"

	config "anvarisy.tech/cleangolang/src/app/config"
	dbconfig "anvarisy.tech/cleangolang/src/app/database"
	routerconfig "anvarisy.tech/cleangolang/src/app/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	router *gin.Engine
	db     *sql.DB
}

func Init() *App {
	config.InitConfig()
	var (
		router = routerconfig.GinRouter()
		db     = dbconfig.InitPostgres()
		app    = &App{
			router: router,
			db:     db,
		}
	)
	return app
}

func (app *App) GetEnvironment() string {
	return viper.GetString("APP_ENV")
}

func (app *App) GetHttpRouter() *gin.Engine {
	return app.router
}

func (app *App) GetDb() *sql.DB {
	return app.db
}

func (app *App) SetDb(db *sql.DB) {
	app.db = db
}

func (app *App) Run() error {
	var port = viper.GetString("APP_PORT")
	var host = "0.0.0.0:" + port
	return http.ListenAndServe(host, app.router)
}
