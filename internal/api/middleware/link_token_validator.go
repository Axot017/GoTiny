package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

type LinkTokenValidator struct {
	getLinkDetails *usecase.GetLinkDetails
}

func NewLinkTokenValidator(getLinkDetails *usecase.GetLinkDetails) *LinkTokenValidator {
	return &LinkTokenValidator{getLinkDetails: getLinkDetails}
}

func (m *LinkTokenValidator) Handle(next http.Handler) http.Handler {
	fn := func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "linkId")
		token := request.URL.Query().Get("token")
		link, err := m.getLinkDetails.Call(request.Context(), id, token)
		if err != nil {
			util.WriteError(writer, err)
			return
		}

		if link == nil {
			util.WriteError(writer, model.NewNotFoundError())
			return
		}
		newCtx := context.WithValue(request.Context(), "link", link)
		newRequest := request.WithContext(newCtx)

		next.ServeHTTP(writer, newRequest)
	}

	return http.HandlerFunc(fn)
}
