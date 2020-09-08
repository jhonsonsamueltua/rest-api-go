package middleware

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"

	"github.com/rest-api-go/pkg/models"
)

const SESSION_ID = "id"

type GoMiddleware struct {
	cookieStore *sessions.CookieStore
}

// InitMiddleware intialize the middleware
func InitMiddleware(c *sessions.CookieStore) *GoMiddleware {
	return &GoMiddleware{
		cookieStore: c,
	}
}

func InitCookieStore() *sessions.CookieStore {
	authKey := securecookie.GenerateRandomKey(64)
	encryptionKey := securecookie.GenerateRandomKey(32)

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.MaxAge = 86400 * 1
	store.Options.HttpOnly = true

	return store
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-Custom-Header, Upgrade-Insecure-Requests")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		// Access-Control-Allow-Credentials
		if c.Request().Method == "OPTIONS" {
			return c.JSON(http.StatusOK, "ok")
		}

		return next(c)
	}
}

func (m *GoMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		notAuth := []string{"/api/user/login", "/api/user/register"}
		requestPath := c.Request().URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next(c)
				return nil
			}
		}

		var resp models.Responses
		session, _ := m.cookieStore.Get(c.Request(), "SessionID")

		if len(session.Values) == 0 || !session.Values["Authenticated"].(bool) {
			resp.Message = "Not Register"
			resp.Status = models.StatusFailed
			return c.JSON(http.StatusForbidden, resp)
		}

		next(c)

		return nil
	})
}
