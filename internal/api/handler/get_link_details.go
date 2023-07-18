package handler

import (
	"net/http"

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
type GetLinkDetails struct {
	linkTokenValidator *middleware.LinkTokenValidator
}

func NewGetLinkDetailsHandler(
	linkTokenValidator *middleware.LinkTokenValidator,
) *GetLinkDetails {
	return &GetLinkDetails{
		linkTokenValidator: linkTokenValidator,
	}
}

func (h *GetLinkDetails) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	dto := dto.LinkDtoFromModel(*link)

	util.WriteResponseJson(writer, dto)
}

func (h *GetLinkDetails) Path() string {
	return "/api/v1/link/{linkId:[a-zA-Z0-9]{1,}}"
}

func (h *GetLinkDetails) Method() string {
	return http.MethodGet
}

func (h *GetLinkDetails) Middlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		h.linkTokenValidator.Handle,
	}
}
