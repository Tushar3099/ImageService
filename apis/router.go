package apis

import (
	"net/http"

	"github.com/RetailPulse/modules/controller"
	"github.com/RetailPulse/service"
	"github.com/gorilla/mux"
)

type API struct {
	service *service.Service
	router  *mux.Router
}

func New(c *controller.Controller) *API {
	var api API
	api.service = service.New(c)
	api.Register()
	return &api
}

func (a *API) Router() *mux.Router {
	return a.router
}

func (a *API) Register() {
	a.router = mux.NewRouter().StrictSlash(true)

	a.router.Methods("Post").Path("/api/submit").Handler(http.HandlerFunc(a.service.PostJob))
	a.router.Methods("GET").Path("/api/status").Handler(http.HandlerFunc(a.service.GetJob))
	a.router.Methods("GET").Path("/").Handler(http.HandlerFunc((a.service.GetIndex)))
}
