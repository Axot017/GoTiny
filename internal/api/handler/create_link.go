package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

// swagger:parameters createLink
type createLinkParams struct {
	// in: body
	Body dto.CreateLinkDto
}

// swagger:response createLinkResponse
type createLinkResponse struct {
	// in: body
	Body dto.LinkDto
}

// swagger:route POST /api/v1/link link createLink
//
// # Create short link
//
// This will create a short link with given settings.
//
// Responses:
//
//	201: createLinkResponse
//	400: errorResponse
//	500: errorResponse
type CreateLinkHandler struct {
	createShortLink *usecase.CreateShortLink
	validate        *validator.Validate
}

func NewCreateLinkHandler(
	createShortLink *usecase.CreateShortLink,
	validate *validator.Validate,
) *CreateLinkHandler {
	return &CreateLinkHandler{createShortLink, validate}
}

func (h *CreateLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	create_link_dto, err := util.DeserializeAndValidateBody[dto.CreateLinkDto](request, h.validate)
	if err != nil {
		util.WriteError(writer, err)
		return
	}

	config := model.LinkToCreate{
		Url:                     create_link_dto.Link,
		MaxHits:                 create_link_dto.MaxHits,
		ValidUntil:              create_link_dto.ValidUntil,
		EnableDetailedAnalytics: create_link_dto.EnableDetailedAnalytics,
	}
	link, err := h.createShortLink.Call(request.Context(), config)
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	dto := dto.LinkDtoFromModel(link)

	util.WriteResponseJson(writer, dto, http.StatusCreated)
}

func (h *CreateLinkHandler) Register(router chi.Router) {
	router.Post("/api/v1/link", h.ServeHTTP)
}
