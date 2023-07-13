package handler

import (
	"net/http"

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
// This will create a short link.
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

	config := model.LinkConfig{
		MaxHits:    create_link_dto.MaxHits,
		ValidUntil: create_link_dto.ValidUntil,
		TrackUntil: create_link_dto.TrackUntil,
	}
	link, err := h.createShortLink.Call(request.Context(), create_link_dto.Link, config)
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	dto := dto.LinkDtoFromModel(link)

	util.WriteResponseJson(writer, dto, http.StatusCreated)
}

func (h *CreateLinkHandler) Path() string {
	return "/api/v1/link"
}

func (h *CreateLinkHandler) Method() string {
	return http.MethodPost
}
