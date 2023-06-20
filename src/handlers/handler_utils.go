package handlers

import (
	"net/http"

	"github.com/pa-vuhn/api-mock/src/app"
	"github.com/pa-vuhn/api-mock/src/config"
)

func renderError(c app.Context, err error) {
	switch err {
	case config.ErrorUnauthorized:
		c.RenderEmptyBody(http.StatusUnauthorized)
	case config.ErrorTooManyRequest:
		c.RenderEmptyBody(http.StatusTooManyRequests)
	case config.ErrorLoginError:
		c.RenderErrorJSON(http.StatusUnauthorized, config.ErrorLoginError.Error())
	default:
		c.RenderEmptyBody(http.StatusInternalServerError)
	}
}
