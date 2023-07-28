package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
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
type GetLinkDetailsHandler struct {
	linkTokenValidator *middleware.LinkTokenValidator
}

func NewGetLinkDetailsHandler(
	linkTokenValidator *middleware.LinkTokenValidator,
) *GetLinkDetailsHandler {
	return &GetLinkDetailsHandler{
		linkTokenValidator: linkTokenValidator,
	}
}

func (h *GetLinkDetailsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	dto := dto.LinkDtoFromModel(*link)

	util.WriteResponseJson(writer, dto)
}

func (h *GetLinkDetailsHandler) Register(router chi.Router) {
	router.With(h.linkTokenValidator.Handle).Get("/api/v1/link/{linkId}", h.ServeHTTP)
}
