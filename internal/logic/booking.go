package logic

import (
	"context"
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"proyecto-monolito/internal/db"
)

type ReservationProvider interface {
	GetReservasBySpaceAndDate(ctx context.Context, espacioID int64, fecha string) ([]db.Reserva, error)
}

func ValidarDisponibilidad(ctx context.Context, espacio db.Space, fecha, horaInicio, horaFin string, store ReservationProvider) (bool, error) {
	inicio, err := parseTime(horaInicio)
	if err != nil {
		return false, err
	}
	fin, err := parseTime(horaFin)
	if err != nil {
		return false, err
	}

	horaApertura, err := parseTime(espacio.HoraApertura)
	if err != nil {
		return false, err
	}
	horaCierre, err := parseTime(espacio.HoraCierre)
	if err != nil {
		return false, err
	}

	if inicio < horaApertura || fin > horaCierre || fin <= inicio {
		return false, errors.New("el horario debe estar dentro del rango de apertura y cierre")
	}

	duracion := fin - inicio
	if duracion < int(espacio.DuracionMinMinutos) {
		return false, errors.New("la duración es menor al mínimo permitido")
	}

	reservas, err := store.GetReservasBySpaceAndDate(ctx, espacio.ID, fecha)
	if err != nil {
		return false, err
	}

	for _, reserva := range reservas {
		resInicio, err := parseTime(reserva.HoraInicio)
		if err != nil {
			return false, err
		}
		resFin, err := parseTime(reserva.HoraFin)
		if err != nil {
			return false, err
		}
		if resInicio < fin && resFin > inicio {
			return false, errors.New("la hora seleccionada se solapa con otra reserva")
		}
	}

	return true, nil
}

func CalcularPrecio(espacio db.Space, fecha, horaInicio, horaFin string) float64 {
	inicio, err := parseTime(horaInicio)
	if err != nil {
		return 0
	}
	fin, err := parseTime(horaFin)
	if err != nil {
		return 0
	}

	duracionHoras := float64(fin-inicio) / 60.0
	subtotal := duracionHoras * espacio.PrecioHora

	if esFinDeSemana(fecha) {
		subtotal += subtotal * espacio.RecargoFinSemana
	}

	if duracionHoras >= float64(espacio.HorasParaDescuento) && espacio.HorasParaDescuento > 0 {
		subtotal -= subtotal * espacio.DescuentoVolumen
	}

	return math.Round(subtotal*100) / 100
}

func parseTime(value string) (int, error) {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return 0, errors.New("formato de hora inválido")
	}
	horas, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	minutos, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	if horas < 0 || horas > 23 || minutos < 0 || minutos > 59 {
		return 0, errors.New("hora inválida")
	}
	return horas*60 + minutos, nil
}

func esFinDeSemana(fecha string) bool {
	fechaTime, err := time.Parse("2006-01-02", fecha)
	if err != nil {
		return false
	}
	weekday := fechaTime.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
