package middelware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/kisstc/microblog/pkg/claim"
	"github.com/kisstc/microblog/pkg/response"
)

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("autorization is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid autorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("invalid autorization format")
	}

	return l[1], nil
}

func Authorizator(next http.Handler) http.Handler {
	signingString := os.Getenv("SIGNING_STRING")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, "heree 1"+err.Error())
			return
		}

		c, err := claim.GetFromToken(tokenString, signingString)
		if err != nil {
			response.HTTPError(w, r, http.StatusUnauthorized, "heree 2"+err.Error())
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", c.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
