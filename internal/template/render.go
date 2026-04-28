package template

import (
	"embed"
	"html/template"
	"net/http"

	"proyecto-monolito/internal/model"
)

//go:embed *.gohtml
var templatesFS embed.FS

type Renderer struct {
	templates map[string]*template.Template
}

func NewRenderer() (*Renderer, error) {
	pages := []string{
		"detalle_espacio.gohtml",
		"editar_espacio.gohtml",
		"lista_espacios.gohtml",
		"login.gohtml",
		"nuevo_espacio.gohtml",
		"registro.gohtml",
	}

	templates := make(map[string]*template.Template)
	funcs := template.FuncMap{"mulFloat": func(a, b float64) float64 { return a * b }}

	for _, page := range pages {
		tpl, err := template.New(page).Funcs(funcs).ParseFS(templatesFS, "layout.gohtml", page)
		if err != nil {
			return nil, err
		}
		templates[page] = tpl
	}
	return &Renderer{templates: templates}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data model.ViewData) error {
	if data.Title == "" {
		data.Title = "Reserva de espacios"
	}
	tpl, ok := r.templates[name]
	if !ok {
		return nil // Should handle error ideally
	}
	return tpl.ExecuteTemplate(w, name, data)
}
