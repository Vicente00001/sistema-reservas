-- name: GetReservasBySpaceAndDate :many
SELECT id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado FROM reservas WHERE espacio_id = ? AND fecha = ? AND estado = 'confirmada';

-- name: ListReservasBySpace :many
SELECT id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado FROM reservas WHERE espacio_id = ? AND estado = 'confirmada' ORDER BY fecha, hora_inicio;

-- name: GetReservaByID :one
SELECT id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado FROM reservas WHERE id = ?;

-- name: CreateReserva :one
INSERT INTO reservas (espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado)
VALUES (?, ?, ?, ?, ?, 'confirmada')
RETURNING id, espacio_id, fecha, hora_inicio, hora_fin, precio_total, estado;

-- name: CancelReserva :exec
UPDATE reservas SET estado = 'cancelada' WHERE id = ?;
