package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/util"
	"gotiny/internal/core/usecase"
)

// swagger:parameters getLinkDetails
type getVisitsParams struct {
	// in: path
	LinkId string `json:"linkId"`
	// in: query
	Token string `json:"token"`
	// in: query
	Page *string `json:"page"`
}

// swagger:response getLinkDetailsResponse
type getVisitsResponse struct {
	// in: body
	Body []dto.VisitDto
}

type GetVisitsHandler struct {
	getVisits *usecase.GetLinkVisits
}

// swagger:route GET /api/v1/link/{linkId}/visits link getVisits
//
// # Get link visits
//
// # Get paginated list of visits for a link
//
// Responses:
//
//	200: getLinkDetailsResponse
//	400: errorResponse
//	401: errorResponse
//	404: errorResponse
//	500: errorResponse
func NewGetVisitsHandler(getVisits *usecase.GetLinkVisits) *GetVisitsHandler {
	return &GetVisitsHandler{getVisits: getVisits}
}

func (h *GetVisitsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "linkId")
	pageToken := request.URL.Query().Get("pageToken")
	var pagePtr *string
	if pageToken != "" {
		pagePtr = &pageToken
	}

	visits, err := h.getVisits.Call(request.Context(), id, pagePtr)
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	dto := dto.PagedResponseDtoFromModel(visits, dto.VisitDtoFromModel)

	util.WriteResponseJson(writer, dto)
}

func (h *GetVisitsHandler) Path() string {
	return "/v1/link/{linkId}/visits"
}

func (h *GetVisitsHandler) Method() string {
	return http.MethodGet
}
