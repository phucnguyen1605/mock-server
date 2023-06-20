package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// App is the top-level framework instance.
type (
	App struct {
		router        *mux.Router
		env           map[string]string
		middlewareFns []MiddlewareFunc
	}

	Env map[string]string

	// HandlerFunc defines a function to server HTTP requests.
	HandlerFunc func(Context)

	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

// HTTP methods
const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
)

// New creates an instance of App
func New() *App {
	return &App{
		env:    Env{},
		router: mux.NewRouter(),
	}
}

func (a *App) SetEnv(key, defaultValue string) {
	v := defaultValue

	// os.Getenv can not recognize whether the environment variable does not exist or is empty.
	for _, e := range os.Environ() {
		arr := strings.Split(e, "=")
		if arr[0] == key {
			v = arr[1]
		}
	}

	a.env[key] = v
}

func (a *App) GET(path string, h HandlerFunc) {
	a.add(GET, path, h)
}

func (a *App) POST(path string, h HandlerFunc) {
	a.add(POST, path, h)
}

func (a *App) PUT(path string, h HandlerFunc) {
	a.add(PUT, path, h)
}

func (a *App) DELETE(path string, h HandlerFunc) {
	a.add(DELETE, path, h)
}

func (a *App) OPTIONS(path string, h HandlerFunc) {
	a.add(OPTIONS, path, h)
}

func (a *App) UseAppMiddlewareFunc(fns ...MiddlewareFunc) {
	a.middlewareFns = append(a.middlewareFns, fns...)
}

func (a *App) UserHttpMiddlewareFunc(fn mux.MiddlewareFunc) {
	a.router.Use(fn)
}

func (a *App) Handle() {
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", a.env["SERVER_PORT"]),
		Handler: a.router,
	}

	log.Printf("Starting server on port: %s\n", a.env["SERVER_PORT"])
	log.Fatal(srv.ListenAndServe())
}

func (a *App) add(method, path string, handler HandlerFunc) {
	h := func(resp http.ResponseWriter, req *http.Request) {
		c := newContext(resp, req, a.env)
		log.Println(req.Method, req.URL.Path)

		defer func() {
			if rcv := recover(); rcv != nil {
				log.Println(rcv)
				c.RenderErrorJSON(http.StatusInternalServerError, rcv)
			}
		}()
		a.execHandlerFunc(c, handler)
	}

	a.router.HandleFunc(path, h).Methods(method, OPTIONS)
}

func (a *App) execHandlerFunc(c Context, handler HandlerFunc) {
	h := handler

	for i := len(a.middlewareFns) - 1; i >= 0; i-- {
		h = a.middlewareFns[i](h)
	}

	h(c)
}
