package db

import (
	"context"
	"database/sql"
)

type Queries struct {
	db *sql.DB
}

type User struct {
	ID           int64  `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

type Space struct {
	ID                 int64   `db:"id"`
	UsuarioID          int64   `db:"usuario_id"`
	Nombre             string  `db:"nombre"`
	Tipo               string  `db:"tipo"`
	HoraApertura       string  `db:"hora_apertura"`
	HoraCierre         string  `db:"hora_cierre"`
	DuracionMinMinutos int32   `db:"duracion_min_minutos"`
	PrecioHora         float64 `db:"precio_hora"`
	RecargoFinSemana   float64 `db:"recargo_fin_semana"`
	DescuentoVolumen   float64 `db:"descuento_volumen"`
	HorasParaDescuento int32   `db:"horas_para_descuento"`
}

type Reserva struct {
	ID          int64   `db:"id"`
	EspacioID   int64   `db:"espacio_id"`
	Fecha       string  `db:"fecha"`
	HoraInicio  string  `db:"hora_inicio"`
	HoraFin     string  `db:"hora_fin"`
	PrecioTotal float64 `db:"precio_total"`
	Estado      string  `db:"estado"`
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, "SELECT id, email, password_hash FROM usuarios WHERE email = ?", email)
	var user User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		return User{}, err
	}
	return user, nil
}

func (q *Queries) CreateUser(ctx context.Context, email string, passwordHash string) (User, error) {
	row := q.db.QueryRowContext(ctx, "INSERT INTO usuarios (email, password_hash) VALUES (?, ?) RETURNING id, email, password_hash", email, passwordHash)
	var user User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		return User{}, err
	}
	return user, nil
}

func (q *Queries) ListSpacesByUser(ctx context.Context, usuarioID int64) ([]Space, error) {
	rs, err := q.db.QueryContext(ctx, "SELECT id, usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento FROM espacios WHERE usuario_id = ?", usuarioID)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var spaces []Space
	for rs.Next() {
		var s Space
		if err := rs.Scan(&s.ID, &s.UsuarioID, &s.Nombre, &s.Tipo, &s.HoraApertura, &s.HoraCierre, &s.DuracionMinMinutos, &s.PrecioHora, &s.RecargoFinSemana, &s.DescuentoVolumen, &s.HorasParaDescuento); err != nil {
			return nil, err
		}
		spaces = append(spaces, s)
	}
	return spaces, rs.Err()
}

func (q *Queries) GetSpaceByID(ctx context.Context, id int64) (Space, error) {
	row := q.db.QueryRowContext(ctx, "SELECT id, usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento FROM espacios WHERE id = ?", id)
	var s Space
	if err := row.Scan(&s.ID, &s.UsuarioID, &s.Nombre, &s.Tipo, &s.HoraApertura, &s.HoraCierre, &s.DuracionMinMinutos, &s.PrecioHora, &s.RecargoFinSemana, &s.DescuentoVolumen, &s.HorasParaDescuento); err != nil {
		return Space{}, err
	}
	return s, nil
}

func (q *Queries) CreateSpace(ctx context.Context, usuarioID int64, nombre, tipo, horaApertura, horaCierre string, duracionMinMinutos int32, precioHora, recargoFinSemana, descuentoVolumen float64, horasParaDescuento int32) (Space, error) {
	row := q.db.QueryRowContext(ctx, "INSERT INTO espacios (usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento", usuarioID, nombre, tipo, horaApertura, horaCierre, duracionMinMinutos, precioHora, recargoFinSemana, descuentoVolumen, horasParaDescuento)
	var s Space
	if err := row.Scan(&s.ID, &s.UsuarioID, &s.Nombre, &s.Tipo, &s.HoraApertura, &s.HoraCierre, &s.DuracionMinMinutos, &s.PrecioHora, &s.RecargoFinSemana, &s.DescuentoVolumen, &s.HorasParaDescuento); err != nil {
		return Space{}, err
	}
	return s, nil
}

func (q *Queries) UpdateSpace(ctx context.Context, nombre, tipo, horaApertura, horaCierre string, duracionMinMinutos int32, precioHora, recargoFinSemana, descuentoVolumen float64, horasParaDescuento int32, id int64) error {
	_, err := q.db.ExecContext(ctx, "UPDATE espacios SET nombre=?, tipo=?, hora_apertura=?, hora_cierre=?, duracion_min_minutos=?, precio_hora=?, recargo_fin_semana=?, descuento_volumen=?, horas_para_descuento=? WHERE id = ?", nombre, tipo, horaApertura, horaCierre, duracionMinMinutos, precioHora, recargoFinSemana, descuentoVolumen, horasParaDescuento, id)
	return err
}

func (q *Queries) GetReservasBySpaceAndDate(ctx context.Context, espacioID int64, fecha string) ([]Reserva, error) {
	rs, err := q.db.QueryContext(ctx, "SELECT id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado FROM reservas WHERE espacio_id = ? AND fecha = ? AND estado = 'confirmada'", espacioID, fecha)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var reservas []Reserva
	for rs.Next() {
		var r Reserva
		if err := rs.Scan(&r.ID, &r.EspacioID, &r.Fecha, &r.HoraInicio, &r.HoraFin, &r.PrecioTotal, &r.Estado); err != nil {
			return nil, err
		}
		reservas = append(reservas, r)
	}
	return reservas, rs.Err()
}

func (q *Queries) ListReservasBySpace(ctx context.Context, espacioID int64) ([]Reserva, error) {
	rs, err := q.db.QueryContext(ctx, "SELECT id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado FROM reservas WHERE espacio_id = ? AND estado = 'confirmada' ORDER BY fecha, hora_inicio", espacioID)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var reservas []Reserva
	for rs.Next() {
		var r Reserva
		if err := rs.Scan(&r.ID, &r.EspacioID, &r.Fecha, &r.HoraInicio, &r.HoraFin, &r.PrecioTotal, &r.Estado); err != nil {
			return nil, err
		}
		reservas = append(reservas, r)
	}
	return reservas, rs.Err()
}

func (q *Queries) GetReservaByID(ctx context.Context, id int64) (Reserva, error) {
	row := q.db.QueryRowContext(ctx, "SELECT id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado FROM reservas WHERE id = ?", id)
	var r Reserva
	if err := row.Scan(&r.ID, &r.EspacioID, &r.Fecha, &r.HoraInicio, &r.HoraFin, &r.PrecioTotal, &r.Estado); err != nil {
		return Reserva{}, err
	}
	return r, nil
}

func (q *Queries) CreateReserva(ctx context.Context, espacioID int64, fecha, horaInicio, horaFin string, precioTotal float64) (Reserva, error) {
	row := q.db.QueryRowContext(ctx, "INSERT INTO reservas (espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado) VALUES (?, ?, ?, ?, ?, 'confirmada') RETURNING id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado", espacioID, fecha, horaInicio, horaFin, precioTotal)
	var r Reserva
	if err := row.Scan(&r.ID, &r.EspacioID, &r.Fecha, &r.HoraInicio, &r.HoraFin, &r.PrecioTotal, &r.Estado); err != nil {
		return Reserva{}, err
	}
	return r, nil
}

func (q *Queries) CancelReserva(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, "UPDATE reservas SET estado = 'cancelada' WHERE id = ?", id)
	return err
}
