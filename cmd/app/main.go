package main

import (
    "log"
    "net/http"
    "testAPI/internal/controllers"
    "testAPI/internal/database"
    "github.com/gorilla/mux"
)

func main() {
    database.InitDB()

    r := mux.NewRouter()

    r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
    r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
    r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
    r.HandleFunc("/products", controllers.CreateProductHandler).Methods("POST")
    r.HandleFunc("/products", controllers.GetProductsHandler).Methods("GET")
    r.HandleFunc("/products/{id}", controllers.GetProductHandler).Methods("GET")
    r.HandleFunc("/products/{id}", controllers.UpdateProductHandler).Methods("PUT")
    r.HandleFunc("/products/{id}", controllers.DeleteProductHandler).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r))
}
