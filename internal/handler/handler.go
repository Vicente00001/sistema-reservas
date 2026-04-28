package handler

import (
	"net/http"

	"proyecto-monolito/internal/auth"
	"proyecto-monolito/internal/model"
	"proyecto-monolito/internal/store"
	"proyecto-monolito/internal/template"
)

type Handler struct {
	Store     store.Store
	Auth      *auth.Auth
	Templates *template.Renderer
}

func NewHandler(store store.Store, authService *auth.Auth, templates *template.Renderer) *Handler {
	return &Handler{Store: store, Auth: authService, Templates: templates}
}

// view construye un ViewData con el UserID siempre relleno desde el contexto de la request.
func (h *Handler) view(r *http.Request, d model.ViewData) model.ViewData {
	d.UserID = auth.GetCurrentUserID(r)
	return d
}

func (h *Handler) RedirectToSpaces(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/espacios", http.StatusSeeOther)
}
