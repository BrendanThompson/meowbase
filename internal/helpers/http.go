package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func RespondError(w http.ResponseWriter, data map[string]interface{}) {
	resp := Message(http.StatusBadRequest, "error")
	resp["data"] = data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RespondErrorMessage(w http.ResponseWriter, message string) {
	resp := Message(http.StatusBadRequest, "error")
	resp["message"] = message
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func RespondSuccess(w http.ResponseWriter, data interface{}) {
	resp := Message(http.StatusOK, "success")
	resp["data"] = data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetIDFromRequest(r *http.Request) int {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("unable to retrieve id from query parameters: %v", err)
	}

	return id
}

type Response struct {
	Status  int                 `json:"status"`            // HTTP status code
	Message string              `json:"message,omitempty"` // Message for the response
	Data    interface{}         `json:"data,omitempty"`    // Data object for successful response
	Error   string              `json:"error,omitempty"`   // Error message
	Writer  http.ResponseWriter `json:"-"`                 // Response Writer
}

func (r *Response) Respond() {
	r.Writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(r.Writer).Encode(r)
}
