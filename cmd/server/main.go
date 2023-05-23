package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	swag "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/crnvl96/go-api/configs"
	_ "github.com/crnvl96/go-api/docs"
	"github.com/crnvl96/go-api/internal/entity"
	"github.com/crnvl96/go-api/internal/infra/database"
	"github.com/crnvl96/go-api/internal/infra/webserver/handlers"
)

// @title           Golang API
// @version         1.0
// @description     Golang API example
// @termsOfService  http://swagger.io/terms/

// @contact.name   crnvl96
// @contact.url    http://github.com/crnvl96
// @contact.email  adran.carnavale@gmail.com

// @license.name  MIT
// @license.url   https://mit-license.org/

// @host                        localhost:8000
// @BasePath                    /
// @securityDefinitions.apiKey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiration)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", productHandler.GetProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/token", userHandler.GetJWT)

	r.Get("/docs/*", swag.Handler(swag.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
