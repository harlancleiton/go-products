package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/harlancleiton/go-products/configs"
	"github.com/harlancleiton/go-products/internal/entity"
	"github.com/harlancleiton/go-products/internal/infra/database"
	"github.com/harlancleiton/go-products/internal/infra/webserver/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

	http.ListenAndServe(":"+config.Server.Port, r)
}
