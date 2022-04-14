package handleFunc

import (
	"encoding/json"
	"log"
	"net/http"
	"product/pkg/application/services"
	"product/pkg/domain"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var request domain.Product

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Can't decode into json")); err != nil {
			log.Println("handleFunc/AddProduct/Decode/write")
			return
		}
		log.Println("handleFunc/AddProduct/Decode")
		return
	}
	response := services.AddProduct(request)
	response.Send(w, "AddProduct")
}
