package routes

import (
    "infoto-web-2/helpers"
	"infoto-web-2/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func Setup(data db.Table) *echo.Echo {
    e := echo.New()

    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.Recover())
    e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
        rate.Limit(20),
    )))

    helpers.NewTemplateRenderer(e)
    e.GET("/", getIndex(data))

    return e
}

func getIndex(data db.Table) echo.HandlerFunc{
    return func (c echo.Context) error {
        item, err := data.Get(0)
        if err != nil {
            c.Render(http.StatusBadRequest, "error", nil)
        }
        return c.Render(http.StatusOK, "index", item)
    }
}
