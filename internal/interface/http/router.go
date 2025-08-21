// File: internal/interface/http/router.go
// Sets up Fiber routes, middleware, and starts the server

package router

import (
	"log"
	"os"
	"training-portal/configs"
	"training-portal/internal/interface/http/handler"
	"training-portal/internal/interface/http/middleware"
	"training-portal/internal/interface/repository/postgres"
	courseusecase "training-portal/internal/usecase/course"
	userusecase "training-portal/internal/usecase/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func SetupAndRun() {
	// Load .env and config
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	configs.LoadConfig()

	// Set env vars from config if not already set
	setEnvIfEmpty("PORT", "3000")
	setEnvIfEmpty("JWT_SECRET", viperGetString("jwt.secret"))
	setEnvIfEmpty("DB_HOST", viperGetString("database.host"))
	setEnvIfEmpty("DB_PORT", viperGetString("database.port"))
	setEnvIfEmpty("DB_USER", viperGetString("database.user"))
	setEnvIfEmpty("DB_PASSWORD", viperGetString("database.password"))
	setEnvIfEmpty("DB_NAME", viperGetString("database.dbname"))

	// Connect to DB
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Init repositories
	userRepo := postgres.NewUserRepository(db)
	courseRepo := postgres.NewCourseRepository(db)
	moduleRepo := postgres.NewModuleRepository(db)

	// Init services
	userService := &userusecase.UserService{Repo: userRepo}
	courseService := &courseusecase.CourseService{Repo: courseRepo}
	moduleService := &courseusecase.ModuleService{Repo: moduleRepo}

	// Init handlers
	userHandler := &handler.UserHandler{Service: userService}
	courseHandler := &handler.CourseHandler{Service: courseService}
	moduleHandler := &handler.ModuleHandler{Service: moduleService}

	app := fiber.New()

	// Enable CORS for all origins (adjust as needed for production)
	app.Use(cors.New())

	// Public routes
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	app.Get("/user/:id", userHandler.GetUser)
	app.Get("/users", userHandler.ListUsers)
	app.Get("/course/:id", courseHandler.GetCourse)
	app.Get("/courses", courseHandler.ListCourses)
	app.Get("/course/:course_id/modules", moduleHandler.ListModulesByCourse)

	// Protected API routes
	api := app.Group("/api", middleware.JWTMiddleware())

	// User management
	api.Put("/user/:id", userHandler.UpdateUser)
	api.Put("/user/:id/password", userHandler.UpdatePassword)
	api.Delete("/user/:id", userHandler.DeleteUser)

	// Course management
	api.Post("/course", courseHandler.CreateCourse)
	api.Put("/course/:id", courseHandler.UpdateCourse)
	api.Delete("/course/:id", courseHandler.DeleteCourse)

	// Module management
	api.Post("/module", moduleHandler.CreateModule)
	api.Get("/module/:id", moduleHandler.GetModule)
	api.Put("/module/:id", moduleHandler.UpdateModule)
	api.Delete("/module/:id", moduleHandler.DeleteModule)

	api.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the protected dashboard!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}

// setEnvIfEmpty sets an environment variable if it is not already set.
func setEnvIfEmpty(key, value string) {
	if os.Getenv(key) == "" && value != "" {
		os.Setenv(key, value)
	}
}

// viperGetString is a helper to get a string from viper config.
func viperGetString(key string) string {
	return configsGetString(key)
}

// configsGetString is a wrapper for viper.GetString.
func configsGetString(key string) string {
	return viper.GetString(key)
}
