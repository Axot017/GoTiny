package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

// swagger:parameters getLinkDetails
type getLinkDetailsParams struct {
	// in: path
	LinkId string `json:"linkId"`
	// in: query
	Token string `json:"token"`
}

// swagger:response getLinkDetailsResponse
type getLinkDetailsResponse struct {
	// in: body
	Body dto.LinkDto
}

// swagger:route GET /api/v1/link/{linkId} link getLinkDetails
//
// # Get link details
//
// Get details of link with given id.
//
// Responses:
//
//	200: getLinkDetailsResponse
//	400: errorResponse
//	401: errorResponse
//	404: errorResponse
//	500: errorResponse
type GetLinkDetails struct {
	createShortLink *usecase.GetLinkDetails
}

func NewGetLinkDetails(createShortLink *usecase.GetLinkDetails) *GetLinkDetails {
	return &GetLinkDetails{
		createShortLink: createShortLink,
	}
}

func (h *GetLinkDetails) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "linkId")
	token := request.URL.Query().Get("token")
	link, err := h.createShortLink.Call(request.Context(), id, token)
	if err != nil {
		util.WriteError(writer, err)
		return
	}

	if link == nil {
		util.WriteError(writer, model.NewNotFoundError())
		return
	}
	dto := dto.LinkDtoFromModel(*link)

	util.WriteResponseJson(writer, dto)
}

func (h *GetLinkDetails) Path() string {
	return "/v1/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *GetLinkDetails) Method() string {
	return http.MethodGet
}
