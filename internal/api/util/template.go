package util

import (
	"html/template"
	"net/http"

	"golang.org/x/exp/slog"
)

var TemplatesFunctions = template.FuncMap{
	"eq": func(x, y interface{}) bool {
		return x == y
	},
	"sub": func(y, x int) int {
		return x - y
	},
}

func WriteTemplate(
	r *http.Request,
	w http.ResponseWriter,
	t *template.Template,
	name string,
	data interface{},
) {
	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		slog.ErrorCtx(r.Context(), "error while executing template", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
}
