package handler

import (
	"html/template"
	"net/http"

	"github.com/go-playground/form/v4"
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
	formDecoder     *form.Decoder
}

func NewAjaxCreateLinkHandler(
	createShortLink *usecase.CreateShortLink,
	validate *validator.Validate,
	template *template.Template,
	formDecoder *form.Decoder,
) *AjaxCreateLinkHandler {
	return &AjaxCreateLinkHandler{
		createShortLink,
		validate,
		template,
		formDecoder,
	}
}

func (h *AjaxCreateLinkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	create_link_dto, err := util.DecodeAndValidateFrom[dto.CreateLinkDto](
		request,
		h.validate,
		h.formDecoder,
	)
	if err != nil {
		h.template.ExecuteTemplate(writer, "timed_error.html", nil)
		return
	}
	config := model.LinkToCreate{
		Url:        create_link_dto.Link,
		MaxHits:    create_link_dto.MaxHits,
		ValidUntil: create_link_dto.ValidUntil,
		TrackUntil: create_link_dto.TrackUntil,
	}
	link, err := h.createShortLink.Call(request.Context(), config)
	if err != nil {
		h.template.ExecuteTemplate(writer, "timed_error.html", err.Error())
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
