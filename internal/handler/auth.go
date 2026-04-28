package handler

import (
	"net/http"
	"proyecto-monolito/internal/model"
)

func (h *Handler) RegisterForm(w http.ResponseWriter, r *http.Request) {
	h.Templates.Render(w, "registro.gohtml", model.ViewData{Title: "Registro"})
}

func (h *Handler) RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.Templates.Render(w, "registro.gohtml", model.ViewData{Title: "Registro", Error: "No se pudo procesar el formulario"})
		return
	}

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	if email == "" || password == "" {
		h.Templates.Render(w, "registro.gohtml", model.ViewData{Title: "Registro", Error: "Email y contraseña son obligatorios"})
		return
	}

	if _, err := h.Auth.Register(r.Context(), email, password); err != nil {
		h.Templates.Render(w, "registro.gohtml", model.ViewData{Title: "Registro", Error: "No se pudo crear el usuario"})
		return
	}

	h.Auth.SetFlash(w, r, "Cuenta creada con éxito, por favor ingresa")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handler) LoginForm(w http.ResponseWriter, r *http.Request) {
	flash, _ := h.Auth.GetFlash(w, r)
	h.Templates.Render(w, "login.gohtml", model.ViewData{Title: "Iniciar sesión", Flash: flash})
}

func (h *Handler) LoginSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.Templates.Render(w, "login.gohtml", model.ViewData{Title: "Iniciar sesión", Error: "No se pudo procesar el formulario"})
		return
	}

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	if email == "" || password == "" {
		h.Templates.Render(w, "login.gohtml", model.ViewData{Title: "Iniciar sesión", Error: "Email y contraseña son obligatorios"})
		return
	}

	userID, err := h.Auth.Login(r.Context(), email, password)
	if err != nil {
		h.Templates.Render(w, "login.gohtml", model.ViewData{Title: "Iniciar sesión", Error: "Credenciales inválidas"})
		return
	}

	if err := h.Auth.SetSessionUser(w, r, userID); err != nil {
		h.Templates.Render(w, "login.gohtml", model.ViewData{Title: "Iniciar sesión", Error: "No se pudo iniciar sesión"})
		return
	}

	http.Redirect(w, r, "/espacios", http.StatusSeeOther)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if err := h.Auth.Logout(w, r); err != nil {
		h.Templates.Render(w, "login.gohtml", model.ViewData{Title: "Iniciar sesión", Error: "No se pudo cerrar sesión"})
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
