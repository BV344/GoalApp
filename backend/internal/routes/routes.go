package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/bv344/goalapp/internal/config"
	"github.com/bv344/goalapp/internal/handlers"
	"github.com/bv344/goalapp/internal/middleware"
)

func Setup(cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	goalHandler := handlers.NewGoalHandler()
	authHandler := handlers.NewAuthHandler()
	scheduleHandler := handlers.NewScheduleHandler()

	r.GET("/api/health", handlers.Health)

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.Auth(cfg.JWTSecret))
	{
		goals := api.Group("/goals")
		{
			goals.GET("", goalHandler.List)
			goals.POST("", goalHandler.Create)
			goals.GET("/:id", goalHandler.Get)
			goals.PUT("/:id", goalHandler.Update)
			goals.DELETE("/:id", goalHandler.Delete)
		}
		schedule := api.Group("/schedule")
		{
			schedule.GET("", scheduleHandler.List)
			schedule.POST("", scheduleHandler.Create)
			schedule.DELETE("/:id", scheduleHandler.Delete)
		}
	}

	return r
}
