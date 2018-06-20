package api

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/vbogretsov/sgmock"
)

const (
	authorization = "Authorization"
	bearer        = "Bearer "
)

func authentication(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			hdr := req.Header.Get(authorization)

			token := hdr[len(bearer):]
			if key != token {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}

type api struct {
	mock sgmock.Mock
}

type success struct {
	Message string `json:"message"`
}

func (a *api) send(c echo.Context) error {
	msg := sgmock.Message{}

	if err := c.Bind(&msg); err != nil {
		return err
	}

	// TODO(vbogretsov): validate and return SendGrid validation errors.
	if err := a.mock.Send(msg); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, success{Message: "message sent"})
}

func (a *api) list(c echo.Context) error {
	return c.JSON(http.StatusOK, a.mock.List())
}

func (a *api) clear(c echo.Context) error {
	a.mock.Clear()
	return nil
}

// New creates new API router.
func New(key string, mock sgmock.Mock) *echo.Echo {
	e := echo.New()

	v3 := e.Group("/v3/mail", authentication(key))

	a := &api{mock: mock}

	v3.POST("/send", a.send)
	e.GET("/ctl/list", a.list)
	e.POST("/ctl/clear", a.clear)

	return e
}
