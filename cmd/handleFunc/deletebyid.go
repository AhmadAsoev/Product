package handleFunc

import (
	"github.com/gorilla/mux"
	"net/http"
	"product/pkg/application/services"
)

func DeleteById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := services.DeleteById(id)
	response.Send(w, "DeleteById")
}
