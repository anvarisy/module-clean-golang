package auth

import (
	"anvarisy.tech/cleangolang/src/app"
	module "anvarisy.tech/cleangolang/src/app/module"
	"anvarisy.tech/cleangolang/src/modules/auth/domain/usecase"
	"anvarisy.tech/cleangolang/src/modules/auth/endpoint"
	"anvarisy.tech/cleangolang/src/modules/auth/repository"
	httptransport "anvarisy.tech/cleangolang/src/modules/auth/transport/http"
	"github.com/go-chi/chi"
)

type Module struct {
	repository usecase.RepositoryInterface
	business   usecase.BusinessInterface
	endpoint   endpoint.Endpoint
	router     *chi.Mux
}

func Init(app *app.App) module.ModuleInterface {
	var (
		repo    = repository.Init(app)
		usecase = usecase.Init()
		ep      = endpoint.Init(usecase, repo)
		handler = httptransport.Init(ep)
	)

	return &Module{
		repository: repo,
		business:   usecase,
		endpoint:   ep,
		router:     handler,
	}
}

func (module *Module) GetHttpRouter() *chi.Mux {
	return module.router
}
