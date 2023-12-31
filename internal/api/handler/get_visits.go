package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

// swagger:parameters getVisits
type getVisitsParams struct {
	// in: path
	LinkId string `json:"linkId"`
	// in: query
	Token string `json:"token"`
	// in: query
	Page *string `json:"page"`
}

// swagger:response getVisitsResponse
type getVisitsResponse struct {
	// in: body
	Body []dto.VisitDto
}

type GetVisitsHandler struct {
	getVisits          *usecase.GetLinkVisits
	linkTokenValidator *middleware.LinkTokenValidator
}

// swagger:route GET /api/v1/link/{linkId}/visits link getVisits
//
// # Get link visits
//
// Get paginated list of visits for a link containing IP, user agent, and approximate location.
// If detailed diagnostics is disabled it will return an empty list.
//
// Responses:
//
//	200: getVisitsResponse
//	400: errorResponse
//	401: errorResponse
//	404: errorResponse
//	500: errorResponse
func NewGetVisitsHandler(
	getVisits *usecase.GetLinkVisits,
	linkTokenValidator *middleware.LinkTokenValidator,
) *GetVisitsHandler {
	return &GetVisitsHandler{
		getVisits:          getVisits,
		linkTokenValidator: linkTokenValidator,
	}
}

func (h *GetVisitsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	pageToken := request.URL.Query().Get("pageToken")
	var pagePtr *string
	if pageToken != "" {
		pagePtr = &pageToken
	}

	visits, err := h.getVisits.Call(request.Context(), link.Id, pagePtr)
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	dto := dto.PagedResponseDtoFromModel(visits, dto.VisitDtoFromModel)

	util.WriteResponseJson(writer, dto)
}

func (h *GetVisitsHandler) Register(router chi.Router) {
	router.With(h.linkTokenValidator.Handle).Get("/api/v1/link/{linkId}/visits", h.ServeHTTP)
}
