package handler

import (
	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"

	"gotiny/internal/api/dto"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

type AjaxCreateLinkHandler struct {
	createShortLink *usecase.CreateShortLink
	validate        *validator.Validate
	template        *template.Template
}

func NewAjaxCreateLinkHandler(
	createShortLink *usecase.CreateShortLink,
	validate *validator.Validate,
	template *template.Template,
) *AjaxCreateLinkHandler {
	return &AjaxCreateLinkHandler{
		createShortLink,
		validate,
		template,
	}
}

func (h *AjaxCreateLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	create_link_dto, err := util.DeserializeAndValidateBody[dto.AjaxCreateLinkDto](
		request,
		h.validate,
	)
	if err != nil {
		println("error", err.Error())
		h.template.ExecuteTemplate(writer, "error.html", nil)
		return
	}
	config := model.LinkConfig{
		// MaxHits:    create_link_dto.MaxHits,
		// ValidUntil: create_link_dto.ValidUntil,
		// TrackUntil: create_link_dto.TrackUntil,
	}
	link, err := h.createShortLink.Call(request.Context(), create_link_dto.Link, config)
	if err != nil {
		println("error - use case", err.Error())
		h.template.ExecuteTemplate(writer, "error.html", nil)
		return
	}

	h.template.ExecuteTemplate(writer, "link_details.html", link)
}

func (h *AjaxCreateLinkHandler) Path() string {
	return "/ajax/link"
}

func (h *AjaxCreateLinkHandler) Method() string {
	return http.MethodPost
}
