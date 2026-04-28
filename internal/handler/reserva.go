package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"proyecto-monolito/internal/auth"
	"proyecto-monolito/internal/logic"
	"proyecto-monolito/internal/model"
)

type availabilityResponse struct {
	Available bool    `json:"available"`
	Message   string  `json:"message"`
	Price     float64 `json:"price,omitempty"`
}

type priceResponse struct {
	Price float64 `json:"price"`
}

func (h *Handler) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	espacioID, err := strconv.ParseInt(r.URL.Query().Get("espacio_id"), 10, 64)
	if err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: "espacio_id inválido"}, http.StatusBadRequest)
		return
	}
	fecha := r.URL.Query().Get("fecha")
	horaInicio := r.URL.Query().Get("hora_inicio")
	horaFin := r.URL.Query().Get("hora_fin")

	space, err := h.Store.GetSpaceByID(r.Context(), espacioID)
	if err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: "Espacio no encontrado"}, http.StatusNotFound)
		return
	}

	ok, err := logic.ValidarDisponibilidad(r.Context(), space, fecha, horaInicio, horaFin, h.Store)
	if err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: err.Error()}, http.StatusOK)
		return
	}
	price := logic.CalcularPrecio(space, fecha, horaInicio, horaFin)
	writeJSON(w, availabilityResponse{Available: ok, Message: "Disponible", Price: price}, http.StatusOK)
}

func (h *Handler) PriceJSON(w http.ResponseWriter, r *http.Request) {
	espacioID, err := strconv.ParseInt(r.URL.Query().Get("espacio_id"), 10, 64)
	if err != nil {
		writeJSON(w, priceResponse{Price: 0}, http.StatusBadRequest)
		return
	}
	fecha := r.URL.Query().Get("fecha")
	horaInicio := r.URL.Query().Get("hora_inicio")
	horaFin := r.URL.Query().Get("hora_fin")

	space, err := h.Store.GetSpaceByID(r.Context(), espacioID)
	if err != nil {
		writeJSON(w, priceResponse{Price: 0}, http.StatusNotFound)
		return
	}

	price := logic.CalcularPrecio(space, fecha, horaInicio, horaFin)
	writeJSON(w, priceResponse{Price: price}, http.StatusOK)
}

func (h *Handler) CreateReserva(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		h.Templates.Render(w, "lista_espacios.gohtml", model.ViewData{Title: "Mis espacios", Error: "ID de espacio inválido"})
		return
	}

	if err := r.ParseForm(); err != nil {
		h.Templates.Render(w, "detalle_espacio.gohtml", model.ViewData{Title: "Detalle de espacio", Error: "No se pudo procesar el formulario"})
		return
	}

	espacio, err := h.Store.GetSpaceByID(r.Context(), id)
	if err != nil {
		h.Templates.Render(w, "detalle_espacio.gohtml", model.ViewData{Title: "Detalle de espacio", Error: "Espacio no encontrado"})
		return
	}

	fecha := r.PostForm.Get("fecha")
	horaInicio := r.PostForm.Get("hora_inicio")
	horaFin := r.PostForm.Get("hora_fin")

	ok, err := logic.ValidarDisponibilidad(r.Context(), espacio, fecha, horaInicio, horaFin, h.Store)
	if err != nil || !ok {
		h.Templates.Render(w, "detalle_espacio.gohtml", model.ViewData{Title: "Detalle de espacio", Error: err.Error(), Payload: espacio})
		return
	}

	price := logic.CalcularPrecio(espacio, fecha, horaInicio, horaFin)
	if _, err := h.Store.CreateReserva(r.Context(), id, fecha, horaInicio, horaFin, price); err != nil {
		h.Templates.Render(w, "detalle_espacio.gohtml", model.ViewData{Title: "Detalle de espacio", Error: "No se pudo crear la reserva", Payload: espacio})
		return
	}

	h.Auth.SetFlash(w, r, "Reserva confirmada")
	http.Redirect(w, r, "/espacios/"+strconv.FormatInt(id, 10), http.StatusSeeOther)
}

func (h *Handler) CancelReserva(w http.ResponseWriter, r *http.Request) {
	reservaID, err := strconv.ParseInt(chi.URLParam(r, "reservaID"), 10, 64)
	if err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: "ID de reserva inválido"}, http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: "ID de espacio inválido"}, http.StatusBadRequest)
		return
	}

	espacio, err := h.Store.GetSpaceByID(r.Context(), id)
	if err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: "Espacio no encontrado"}, http.StatusNotFound)
		return
	}

	if espacio.UsuarioID != auth.GetCurrentUserID(r) {
		writeJSON(w, availabilityResponse{Available: false, Message: "No autorizado"}, http.StatusForbidden)
		return
	}

	if err := h.Store.CancelReserva(r.Context(), reservaID); err != nil {
		writeJSON(w, availabilityResponse{Available: false, Message: "No se pudo cancelar la reserva"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", fmt.Sprintf("/espacios/%d", id))
	w.WriteHeader(http.StatusOK)
}

func writeJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
