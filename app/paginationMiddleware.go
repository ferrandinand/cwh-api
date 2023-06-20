package app

import (
	"context"
	"net/http"
	"strconv"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
)

const (
	// PageIDKey refers to the context key that stores the next page id
	PageIDKey string = "page_id"
)

type PaginationMiddleware struct{}

// Pagination middleware is used to extract the next page id from the url query
func (a PaginationMiddleware) paginationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			PageID := r.URL.Query().Get(string(PageIDKey))
			intPageID := 0
			var err error
			if PageID != "" {
				intPageID, err = strconv.Atoi(PageID)
				if err != nil {
					appError := errs.AppError{http.StatusBadRequest, err.Error()}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			}
			ctx := context.WithValue(r.Context(), PageIDKey, intPageID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
