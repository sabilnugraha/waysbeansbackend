package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	productsdto "waysbeans/dto/products"
	dto "waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type productHandler struct {
	ProductRepository repositories.ProductRepository
}

// Create `path_file` Global variable here ...
var path_file = "http://localhost:5000/uploads/"

func HandlerProduct(ProductRepository repositories.ProductRepository) *productHandler {
	return &productHandler{ProductRepository}
}

func convertResponseProduct(u models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
		Image: u.Image,
		Stock: u.Stock,
		Desc:  u.Desc,
	}
}

// get all products
func (h *productHandler) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	for i, p := range products {
		products[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: products}
	json.NewEncoder(w).Encode(response)
}

// get product
func (h *productHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	product.Image = path_file + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: convertResponseProduct(product)}
	json.NewEncoder(w).Encode(response)
}

// create product
func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	request := productsdto.CreateProductRequest{
		Name:  r.FormValue("name"),
		Price: price,
		Stock: stock,
		Desc:  r.FormValue("desc"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db product
	product := models.Product{
		Name:  request.Name,
		Price: request.Price,
		Image: filename,
		Stock: request.Stock,
		Desc:  request.Desc,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
