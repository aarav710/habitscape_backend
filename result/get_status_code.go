package result

import (
	"habitscape/backend/ent"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if ent.IsNotFound(err) {
    return http.StatusNotFound
	}	else {
		return http.StatusBadRequest
	}
}