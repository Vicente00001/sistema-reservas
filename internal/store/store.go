package store

import (
	"context"
	"database/sql"

	"proyecto-monolito/internal/db"
)

type Store interface {
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	CreateUser(ctx context.Context, email, passwordHash string) (db.User, error)
	ListSpacesByUser(ctx context.Context, usuarioID int64) ([]db.Space, error)
	GetSpaceByID(ctx context.Context, id int64) (db.Space, error)
	CreateSpace(ctx context.Context, usuarioID int64, nombre, tipo, horaApertura, horaCierre string, duracionMinMinutos int32, precioHora, recargoFinSemana, descuentoVolumen float64, horasParaDescuento int32) (db.Space, error)
	UpdateSpace(ctx context.Context, nombre, tipo, horaApertura, horaCierre string, duracionMinMinutos int32, precioHora, recargoFinSemana, descuentoVolumen float64, horasParaDescuento int32, id int64) error
	GetReservasBySpaceAndDate(ctx context.Context, espacioID int64, fecha string) ([]db.Reserva, error)
	ListReservasBySpace(ctx context.Context, espacioID int64) ([]db.Reserva, error)
	GetReservaByID(ctx context.Context, id int64) (db.Reserva, error)
	CreateReserva(ctx context.Context, espacioID int64, fecha, horaInicio, horaFin string, precioTotal float64) (db.Reserva, error)
	CancelReserva(ctx context.Context, id int64) error
}

type SQLStore struct {
	queries *db.Queries
}

func NewStore(dbConn *sql.DB) *SQLStore {
	return &SQLStore{queries: db.New(dbConn)}
}

func (s *SQLStore) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return s.queries.GetUserByEmail(ctx, email)
}

func (s *SQLStore) CreateUser(ctx context.Context, email, passwordHash string) (db.User, error) {
	return s.queries.CreateUser(ctx, email, passwordHash)
}

func (s *SQLStore) ListSpacesByUser(ctx context.Context, usuarioID int64) ([]db.Space, error) {
	return s.queries.ListSpacesByUser(ctx, usuarioID)
}

func (s *SQLStore) GetSpaceByID(ctx context.Context, id int64) (db.Space, error) {
	return s.queries.GetSpaceByID(ctx, id)
}

func (s *SQLStore) CreateSpace(ctx context.Context, usuarioID int64, nombre, tipo, horaApertura, horaCierre string, duracionMinMinutos int32, precioHora, recargoFinSemana, descuentoVolumen float64, horasParaDescuento int32) (db.Space, error) {
	return s.queries.CreateSpace(ctx, usuarioID, nombre, tipo, horaApertura, horaCierre, duracionMinMinutos, precioHora, recargoFinSemana, descuentoVolumen, horasParaDescuento)
}

func (s *SQLStore) UpdateSpace(ctx context.Context, nombre, tipo, horaApertura, horaCierre string, duracionMinMinutos int32, precioHora, recargoFinSemana, descuentoVolumen float64, horasParaDescuento int32, id int64) error {
	return s.queries.UpdateSpace(ctx, nombre, tipo, horaApertura, horaCierre, duracionMinMinutos, precioHora, recargoFinSemana, descuentoVolumen, horasParaDescuento, id)
}

func (s *SQLStore) GetReservasBySpaceAndDate(ctx context.Context, espacioID int64, fecha string) ([]db.Reserva, error) {
	return s.queries.GetReservasBySpaceAndDate(ctx, espacioID, fecha)
}

func (s *SQLStore) ListReservasBySpace(ctx context.Context, espacioID int64) ([]db.Reserva, error) {
	return s.queries.ListReservasBySpace(ctx, espacioID)
}

func (s *SQLStore) GetReservaByID(ctx context.Context, id int64) (db.Reserva, error) {
	return s.queries.GetReservaByID(ctx, id)
}

func (s *SQLStore) CreateReserva(ctx context.Context, espacioID int64, fecha, horaInicio, horaFin string, precioTotal float64) (db.Reserva, error) {
	return s.queries.CreateReserva(ctx, espacioID, fecha, horaInicio, horaFin, precioTotal)
}

func (s *SQLStore) CancelReserva(ctx context.Context, id int64) error {
	return s.queries.CancelReserva(ctx, id)
}
