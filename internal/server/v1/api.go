package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kisstc/microblog/internal/data"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())
	r.Mount("/posts", pr.Routes())

	return r
}
