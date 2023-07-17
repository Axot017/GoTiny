package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"gotiny/internal/api/middleware"
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
		return
	}

	fmt.Println(visits)

	h.template.ExecuteTemplate(writer, "link_visits_list.html", map[string]interface{}{
		"Page": visits,
		"Link": link,
	})
}

func (h *AjaxGetVisitsHandler) Path() string {
	return "/link/{linkId}/visits"
}

func (h *AjaxGetVisitsHandler) Method() string {
	return http.MethodGet
}

func (h *AjaxGetVisitsHandler) Middlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		h.linkTokenValidator.Handle,
	}
}
