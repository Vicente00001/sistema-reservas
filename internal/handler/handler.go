package handler

import (
	"net/http"

	"proyecto-monolito/internal/auth"
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

func (h *Handler) RedirectToSpaces(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/espacios", http.StatusSeeOther)
}
