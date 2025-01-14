package middleware

import (
	"context"
	"cynxhost/internal/constant/types"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/helper"
	contextmodel "cynxhost/internal/model/context"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"net/http"
	"strconv"
)

func AuthMiddleware(JWTManager *dependencies.JWTManager, next http.HandlerFunc, debug bool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("AuthToken")
		if err != nil {
			if err == http.ErrNoCookie {
				apiResponse := response.APIResponse{
					Code:  responsecode.CodeAuthenticationError,
					Error: "AuthToken cookie missing",
				}
				helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
				return
			}
			apiResponse := response.APIResponse{
				Code:  responsecode.CodeAuthenticationError,
				Error: "Error retrieving cookie",
			}
			helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
			return
		}

		token := cookie.Value
		claims, err := JWTManager.VerifyToken(token)

		if err != nil { // Replace with your token verification logic
			apiResponse := response.APIResponse{
				Code:  responsecode.CodeAuthenticationError,
				Error: "Invalid or expired access token",
			}
			helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
			return
		}

		// Extract user information from claims
		userId := claims.UserId // Adjust according to your claims structure

		// Convert userId to int
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			apiResponse := response.APIResponse{
				Code:  responsecode.CodeAuthenticationError,
				Error: "Invalid user ID format: " + err.Error(),
			}
			helper.WriteJSONResponse(w, http.StatusUnauthorized, apiResponse)
			return
		}

		// Inject user data into the request context
		ctx := context.WithValue(r.Context(), types.ContextKeyUser, contextmodel.User{
			Id: userIdInt,
		})

		next(w, r.WithContext(ctx))
	}

}
