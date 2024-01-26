package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":"+config.Server.Port, r)
}
