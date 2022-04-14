package handleFunc

import (
	"github.com/gorilla/mux"
	"net/http"
	"product/pkg/application/services"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := services.GetById(id)
	response.Send(w, "GetById")
}
