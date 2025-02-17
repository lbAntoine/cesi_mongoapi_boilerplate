package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lbAntoine/mongoapi_boilerplate/internal/db"
	"github.com/lbAntoine/mongoapi_boilerplate/internal/repositories"
)

type App struct {
	router   *gin.Engine
	mongodb  *db.MongoDB
	userRepo repositories.UserRepository
	config   *Config
}

func NewApp(config *Config) (*App, error) {
	mongodb, err := db.NewMongoDB(config.MongoDB.URI, config.MongoDB.Database, config.MongoDB.Timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(mongodb.GetCollection("users"))

	// Initialize Gin
	gin.SetMode(config.Server.Mode)
	router := gin.Default()

	app := &App{
		router:   router,
		mongodb:  mongodb,
		userRepo: userRepo,
		config:   config,
	}

	app.setupRoutes()
	return app, nil
}

func (a *App) setupRoutes() {
	api := a.router.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("/", a.createUser)
			users.GET("/:id", a.getUser)
			users.GET("/", a.listUsers)
			users.PUT("/:id", a.updateUser)
			users.DELETE("/:id", a.deleteUser)
		}
	}
}

func (a *App) Run() error {
	return a.router.Run(":" + a.config.Server.Port)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.mongodb.Close(ctx)
}
