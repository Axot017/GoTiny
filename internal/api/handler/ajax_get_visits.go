package handler

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gotiny/internal/api/middleware"
	"gotiny/internal/api/util"
	"gotiny/internal/core/model"
	"gotiny/internal/core/usecase"
)

type AjaxGetVisitsHandler struct {
	getVisits          *usecase.GetLinkVisits
	linkTokenValidator *middleware.LinkTokenValidator
	template           *template.Template
}

func NewAjaxGetVisitsHandler(
	getVisits *usecase.GetLinkVisits,
	linkTokenValidator *middleware.LinkTokenValidator,
	template *template.Template,
) *AjaxGetVisitsHandler {
	return &AjaxGetVisitsHandler{
		getVisits:          getVisits,
		linkTokenValidator: linkTokenValidator,
		template:           template,
	}
}

func (h *AjaxGetVisitsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	link := request.Context().Value("link").(*model.Link)
	pageToken := request.URL.Query().Get("pageToken")
	var pagePtr *string
	if pageToken != "" {
		pagePtr = &pageToken
	}

	visits, err := h.getVisits.Call(request.Context(), link.Id, pagePtr)
	if err != nil {
		util.WriteAjaxError(writer, err)
		return
	}

	util.WriteTemplate(request, writer, h.template, "link_visits_list.html", map[string]interface{}{
		"Page": visits,
		"Link": link,
	})
}

func (h *AjaxGetVisitsHandler) Register(router chi.Router) {
	router.With(h.linkTokenValidator.Handle).Get("/ajax/link/{linkId}/visits", h.ServeHTTP)
}
