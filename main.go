package main

import (
	"api-golang/configs"
	"api-golang/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/api/todos", handlers.Create)
	router.Get("/api/todos", handlers.GetAll)
	router.Get("/api/todos/{id}", handlers.Get)
	router.Patch("/api/todos/{id}", handlers.Update)
	router.Delete("/api/todos/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
	fmt.Print("Server is running")
}
