package router

import (
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type route struct {
	method  string
	path    string
	handler handlerFunc
}

type Router struct {
	routes []route
	prefix string
}

func (router *Router) Prefix(prefix string) {
	router.prefix = prefix
}

func (router *Router) Get(path string, handler handlerFunc) {
	router.routes = append(router.routes, route{
		method:  "GET",
		path:    router.prefix + path,
		handler: handler,
	})
}

func (router *Router) Post(path string, handler handlerFunc) {
	router.routes = append(router.routes, route{
		method:  "POST",
		path:    router.prefix + path,
		handler: handler,
	})
}

func (router *Router) Patch(path string, handler handlerFunc) {
	router.routes = append(router.routes, route{
		method:  "PATCH",
		path:    router.prefix + path,
		handler: handler,
	})
}

func (router *Router) Delete(path string, handler handlerFunc) {
	router.routes = append(router.routes, route{
		method:  "DELETE",
		path:    router.prefix + path,
		handler: handler,
	})
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if route.method == r.Method && route.path == r.URL.Path {
			route.handler(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("endpoint not found."))
}
