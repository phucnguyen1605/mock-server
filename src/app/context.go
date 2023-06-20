package app

import (
	"encoding/json"
	"net/http"

	"context"
)

type Context interface {
	Context() context.Context
	Request() *http.Request
	Header(key string) string
	ResponseWriter() http.ResponseWriter
	GetEnv(name string) string
	FormValue(name string) string
	BindJSON(dest interface{}) error

	// renderer
	RenderJSON(i interface{})
	RenderErrorJSON(httpCode int, i interface{})
	RenderEmptyBody(httpCode int)
	RenderNoContent()

	// Storage
	DB() DB

	WithContextValue(key interface{}, v interface{})
	ContextValue(key interface{}) interface{}
}

type defaultContext struct {
	ctx            context.Context
	env            Env
	request        *http.Request
	responseWriter http.ResponseWriter
	db             DB
	vars           map[string]string
}

func NewTestContext() Context {
	c := &defaultContext{}
	return c
}

func newContext(w http.ResponseWriter, r *http.Request, env map[string]string) Context {
	ctx := r.Context()
	c := &defaultContext{
		ctx:            ctx,
		env:            env,
		request:        r,
		responseWriter: w,
		db:             newDB(env),
		vars:           map[string]string{},
	}
	return c
}

func (c *defaultContext) Context() context.Context {
	return c.ctx
}

func (c *defaultContext) Request() *http.Request {
	return c.request
}

func (c *defaultContext) Header(key string) string {
	return c.request.Header.Get(key)
}

func (c *defaultContext) ResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

func (c *defaultContext) GetEnv(name string) string {
	v, _ := c.env[name]
	return v
}

func (c *defaultContext) FormValue(name string) string {
	return c.request.FormValue(name)
}

func (c *defaultContext) BindJSON(dst interface{}) error {
	err := json.NewDecoder(c.request.Body).Decode(dst)
	defer c.request.Body.Close()

	return err
}

func (c *defaultContext) DB() DB {
	return c.db
}

func (c *defaultContext) WithContextValue(key, value interface{}) {
	ctx := context.WithValue(c.request.Context(), key, value)
	c.request = c.Request().WithContext(ctx)
}

func (c *defaultContext) ContextValue(key interface{}) interface{} {
	return c.request.Context().Value(key)
}

// FIXME: move rendering JSON to a renderer
func (c *defaultContext) RenderJSON(i interface{}) {
	setDefaultHeaders(c.ResponseWriter())
	c.ResponseWriter().WriteHeader(http.StatusOK)

	json.NewEncoder(c.ResponseWriter()).Encode(i)
}

// FIXME: move rendering JSON to a renderer
func (c *defaultContext) RenderErrorJSON(httpCode int, i interface{}) {
	setDefaultHeaders(c.ResponseWriter())
	c.ResponseWriter().WriteHeader(httpCode)

	json.NewEncoder(c.ResponseWriter()).Encode(map[string]interface{}{
		"error": i,
	})
}

func (c *defaultContext) RenderEmptyBody(httpCode int) {
	setDefaultHeaders(c.ResponseWriter())
	c.ResponseWriter().WriteHeader(httpCode)
}
func (c *defaultContext) RenderNoContent() {
	setDefaultHeaders(c.ResponseWriter())
	c.ResponseWriter().WriteHeader(http.StatusNoContent)
}

// FIXME: move rendering JSON to a renderer
func setDefaultHeaders(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")
}
