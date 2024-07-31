package controllers

import (
	"encoding/json"
	"net/http"
	"testAPI/internal/database"
	"testAPI/internal/models"

	"github.com/gorilla/mux"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    json.NewDecoder(r.Body).Decode(&product)

    result := database.DB.Create(&product)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
    var products []models.Product
    result := database.DB.Find(&products)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(products)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var product models.Product
    result := database.DB.First(&product, id)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var product models.Product
    result := database.DB.First(&product, id)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
    }

    var updatedProduct models.Product
    json.NewDecoder(r.Body).Decode(&updatedProduct)

    product.Name = updatedProduct.Name
    product.Description = updatedProduct.Description
    product.Price = updatedProduct.Price

    database.DB.Save(&product)
    json.NewEncoder(w).Encode(product)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var product models.Product
    result := database.DB.First(&product, id)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
    }

    database.DB.Delete(&product)
    w.WriteHeader(http.StatusNoContent)
}
