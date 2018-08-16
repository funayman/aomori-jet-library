package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r RouteHolder
)

type RouteHolder struct {
	*mux.Router
}

func init() {
	r.Router = mux.NewRouter()
}

// ReadConfig returns the information
func ReadConfig() RouteHolder {
	return r
}

// Instance returns the router
func Instance() *mux.Router {
	return r.Router
}

func Route(path string, fn http.HandlerFunc) *mux.Route {
	return r.HandleFunc(path, fn)
}

func RouteStatic(path, dir string) *mux.Route {
	// r.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
	return r.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

func GetParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}
