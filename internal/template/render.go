package template

import (
	"embed"
	"html/template"

	"proyecto-monolito/internal/model"
)

//go:embed *.gohtml
var templatesFS embed.FS

type Renderer struct {
	templates *template.Template
}

func NewRenderer() (*Renderer, error) {
	tpl, err := template.New("layout").Funcs(template.FuncMap{
		"mulFloat": func(a, b float64) float64 { return a * b },
	}).ParseFS(templatesFS, "*.gohtml")
	if err != nil {
		return nil, err
	}
	return &Renderer{templates: tpl}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data model.ViewData) error {
	if data.Title == "" {
		data.Title = "Reserva de espacios"
	}
	return r.templates.ExecuteTemplate(w, name, data)
}
