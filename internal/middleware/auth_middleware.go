package middleware

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"net/http"
	"strings"
)

func AuthMiddleware(JWTManager *dependencies.JWTManager, next http.HandlerFunc, debug bool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Check for the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			apiResponse := response.APIResponse{
				Code: responsecode.CodeAuthenticationError,
				Error: "Authorization header missing",
			}
			helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
			return
		}

		// Check if the token starts with "Bearer"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			apiResponse := response.APIResponse{
				Code: responsecode.CodeAuthenticationError,
				Error: "Invalid authorization token format",
			}
			helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
			return
		}

		// You could verify the token here if needed (e.g., check JWT signature)
		token := parts[1]
		if _, err := JWTManager.VerifyToken(token); err != nil { // Replace with your token verification logic
			apiResponse := response.APIResponse{
				Code: responsecode.CodeAuthenticationError,
				Error: "Invalid or expired access token",
			}
			helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
			return
		}

		next(w, r)
	}

}
