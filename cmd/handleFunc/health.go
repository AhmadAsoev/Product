package handleFunc

import (
	"net/http"
	"product/pkg/domain"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Server is Ok!",
	}
	response.Send(w, "Health")
}
