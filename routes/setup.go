package routes

import (
	"htmx-test/db"
	"htmx-test/helpers"
	"htmx-test/models"
	"net/http"
	"strconv"

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

    e.Static("/css", "css")

    helpers.NewTemplateRenderer(e)
    e.GET("/", getIndex(data))
    e.POST("/item", getAddItem(data))

    return e
}

func ErrorToast(msg string) models.ToastResponse {
    return models.ToastResponse {
        ClassName: "notification is-danger",
        Message: msg,
    }
}
func SuccessToast(msg string) models.ToastResponse {
    return models.ToastResponse {
        ClassName: "notification is-primary",
        Message: msg,
    }
}

func getIndex(data db.Table) echo.HandlerFunc{
    return func (c echo.Context) error {
        items, err := data.All()
        if err != nil {
            c.Render(http.StatusBadRequest, "error", nil)
        }
        request := models.IndexRequest{
            Items: items,
        }
        return c.Render(http.StatusOK, "index", request)
    }
}

func getAddItem(data db.Table) echo.HandlerFunc {
    return func (c echo.Context) error {
        itemName := c.FormValue("name")
        quantity, err := strconv.Atoi(c.FormValue("amount"))
        if err != nil {
            c.Render(
                http.StatusBadRequest,
                "toast",
                ErrorToast("amount was incorrect"))
        }
        item := &models.Item {
            Name: itemName,
            Quantity: quantity,
        }
        data.Add(item)
        msg := SuccessToast("Added Item!")
        c.Render(http.StatusOK, "toast", msg)
        return c.Render(http.StatusOK, "item_info", item)
    }
}
