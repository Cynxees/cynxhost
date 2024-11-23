package helper

import (
	"cynxhost/internal/model/response/responsecode"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidateRequest(r *http.Request, dst interface{}, v *validator.Validate) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return errors.New("invalid request payload")
	}

	if err := v.Struct(dst); err != nil {
		return errors.New("validation failed: " + err.Error())
	}

	return nil
}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to decode request: "+err.Error(), http.StatusBadRequest)
	}
}

func GetResponseCodeName(code responsecode.ResponseCode) string {
	if name, exists := responsecode.ResponseCodeNames[code]; exists {
		return name
	}
	return "Unknown Code"
}
