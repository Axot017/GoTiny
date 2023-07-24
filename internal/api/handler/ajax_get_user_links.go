package handler

import (
	"html/template"
	"net/http"

	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
	"gotiny/internal/core/usecase"
)

type AjaxGetUserLinks struct {
	getUserLinks       *usecase.GetUserLinks
	idCookieMiddleware *middleware.IdCookieMiddleware
	template           *template.Template
}

func NewAjaxGetUserLinks(
	getUserLinks *usecase.GetUserLinks,
	idCookieMiddleware *middleware.IdCookieMiddleware,
	template *template.Template,
) *AjaxGetUserLinks {
	return &AjaxGetUserLinks{getUserLinks, idCookieMiddleware, template}
}

func (h *AjaxGetUserLinks) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value("user_id").(*string)
	pageToken := request.URL.Query().Get("pageToken")
	var pagePtr *string
	if pageToken != "" {
		pagePtr = &pageToken
	}
	links, err := h.getUserLinks.Call(request.Context(), *userId, pagePtr)
	if err != nil {
		util.WriteAjaxError(writer, err)
		return
	}
	util.WriteTemplate(request, writer, h.template, "user_links_list.html", links)
}

func (h *AjaxGetUserLinks) Path() string {
	return "/ajax/links"
}

func (h *AjaxGetUserLinks) Method() string {
	return http.MethodGet
}

func (h *AjaxGetUserLinks) Middlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		h.idCookieMiddleware.Handle,
	}
}
