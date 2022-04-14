package handleFunc

import (
	"net/http"
	"product/pkg/application/services"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	response := services.GetAllProducts()
	response.Send(w, "GetALlProducts")
}
