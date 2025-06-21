package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type V1Group struct {
	e     *echo.Echo
	route *echo.Group
}

func NewV1Group(
	e *echo.Echo,

) *V1Group {
	route := e.Group("/api/v1")

	route.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					return respond(c, he.Code, nil, err)
				}

				if be, ok := err.(*echo.BindingError); ok {
					return respond(c, http.StatusBadRequest, map[string]interface{}{
						be.Field: be.Message,
					}, nil)
				}

				c.Logger().Error(err)
				return respond(c, http.StatusInternalServerError, nil, nil)
			}

			return nil
		}
	})

	V1Group := &V1Group{
		e:     e,
		route: route,
	}

	return V1Group
}

func respond(c echo.Context, code int, i interface{}, e error) error {
	response := map[string]interface{}{
		"status": "success",
	}

	if i != nil {
		response["data"] = i
	}

	if e != nil {
		response["message"] = e.Error()
	}

	if code >= 400 && code <= 499 {
		response["status"] = "fail"
	} else if code >= 500 && code <= 599 {
		response["status"] = "error"
	}

	return c.JSON(code, response)
}
