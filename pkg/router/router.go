package router

import (
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type route struct {
	method  string
	path    string
	handler http.Handler
}

type Router struct {
	routes      []route
	middlewares []func(http.Handler) http.Handler
	prefix      string
}

func (router *Router) Prefix(prefix string) {
	router.prefix = prefix
}

func (router *Router) Use(middleware func(http.Handler) http.Handler) {
	router.middlewares = append(router.middlewares, middleware)
}

func (router *Router) Get(path string, handlerFunc handlerFunc) {
	router.appendRoute("GET", path, handlerFunc)
}

func (router *Router) Post(path string, handlerFunc handlerFunc) {
	router.appendRoute("POST", path, handlerFunc)
}

func (router *Router) Patch(path string, handlerFunc handlerFunc) {
	router.appendRoute("PATCH", path, handlerFunc)
}

func (router *Router) Delete(path string, handlerFunc handlerFunc) {
	router.appendRoute("DELETE", path, handlerFunc)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if route.method == r.Method && route.path == r.URL.Path {
			route.handler.ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("endpoint not found."))
}

// Добавление цепочки для обхода middleware в обратном прорядке
func (router *Router) applyMiddlewares(handler http.Handler) http.Handler {
	for i := len(router.middlewares) - 1; i >= 0; i-- {
		handler = router.middlewares[i](handler)
	}
	return handler
}

func (router *Router) appendRoute(method string, path string, handlerFunc handlerFunc) {
	var handler http.Handler
	handler = http.HandlerFunc(handlerFunc)
	handler = router.applyMiddlewares(handler)

	router.routes = append(router.routes, route{
		method:  method,
		path:    router.prefix + path,
		handler: handler,
	})
}
