package middleware

import (
	"log"
	"net/http"

	"github.com/pa-vuhn/api-mock/src/app"
	"github.com/pa-vuhn/api-mock/src/config"
	"github.com/pa-vuhn/api-mock/src/utils/jwt"
)

func AuthMiddleWareFunc() app.MiddlewareFunc {
	return func(next app.HandlerFunc) app.HandlerFunc {
		return func(c app.Context) {
			log.Println("Passed to Auth middleware func")
			if c.Request().Method == http.MethodOptions {
				c.RenderNoContent()
				return
			}

			if isSkipAuthMiddleWare(c.Request()) {
				next(c)
				return
			}

			token := c.Header("Authorization")
			id, ok := jwt.ValidUserToken(token, c.GetEnv("JWT_KEY"))
			if !ok {
				c.RenderEmptyBody(http.StatusUnauthorized)
				return
			}
			c.WithContextValue(config.UserAuthCtxKey, id)

			next(c)
		}
	}
}

func isSkipAuthMiddleWare(r *http.Request) bool {
	skipList := map[string]bool{
		"/login":  true,
		"/signup": true,
	}
	if _, ok := skipList[r.URL.Path]; ok {
		return true
	}
	return false
}
