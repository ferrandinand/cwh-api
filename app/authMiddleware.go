package app

import (
	"net/http"
	"strings"

	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-lib/errs"

	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				user, isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)
				//For each request add user id header
				r.Header.Add("User", user)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{http.StatusForbidden, "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
