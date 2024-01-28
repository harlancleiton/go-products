package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/harlancleiton/go-products/configs"
	_ "github.com/harlancleiton/go-products/docs"
	"github.com/harlancleiton/go-products/internal/entity"
	"github.com/harlancleiton/go-products/internal/infra/database"
	"github.com/harlancleiton/go-products/internal/infra/webserver/handler"
	"github.com/swaggo/http-swagger/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Products API
// @version 1.0
// @description API para gerenciamento de produtos
// @termsOfService http://swagger.io/terms/

// @contact.name Harlan Cleiton
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDb := database.NewProduct(db)
	productHandler := handler.NewProductHandler(productDb)

	userDb := database.NewUser(db)
	userHandler := handler.NewUserHandler(userDb)

	authHandler := handler.NewAuthHandler(config.Server.TokenAuth, config.Server.JwtExpiresIn, userDb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", authHandler.Login)

	r.Post("/users", userHandler.CreateUser)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.Server.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":"+config.Server.Port, r)
}
