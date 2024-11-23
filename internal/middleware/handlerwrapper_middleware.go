package middleware

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/response"
	"net/http"
)

type HandlerFuncWithHelper func(w http.ResponseWriter, r *http.Request) response.APIResponse

func WrapHandler(handler HandlerFuncWithHelper, debug bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Execute the handler and capture response or error
		apiResponse := handler(w, r)
		
		if !debug && apiResponse.Error != "" {
			apiResponse.Error = "Hidden Error, Debug Mode is Enabled"
		}

		if apiResponse.Code != "" {
			apiResponse.CodeName = helper.GetResponseCodeName(apiResponse.Code)
		}

		// Write the successful response
		helper.WriteJSONResponse(w, http.StatusOK, apiResponse)
	}
}
