package main

import (
	"arfifa/belajar-golang-restful-api/app"
	"arfifa/belajar-golang-restful-api/controller"
	"arfifa/belajar-golang-restful-api/helper"
	"arfifa/belajar-golang-restful-api/repository"
	"arfifa/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3009",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}