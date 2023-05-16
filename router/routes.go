package router

import (
	"echo-api/controller"
	"echo-api/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello My Friends")
	})

	// Inisialisasi middleware
	authMiddleware := middleware.AuthMiddleware

	panitiaGroup := e.Group("/panitia")
	panitiaGroup.Use(authMiddleware)
	panitiaGroup.GET("", controller.GetAllUsers)
	panitiaGroup.POST("", controller.CreateUsers)
	panitiaGroup.PUT("/:id", controller.UpdateUsers)
	panitiaGroup.DELETE("/:id", controller.DeleteUsersById)

	return e
}
