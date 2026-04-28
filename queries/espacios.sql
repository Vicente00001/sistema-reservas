-- name: ListSpacesByUser :many
SELECT id, usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento FROM espacios WHERE usuario_id = ?;

-- name: GetSpaceByID :one
SELECT id, usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento FROM espacios WHERE id = ?;

-- name: CreateSpace :one
INSERT INTO espacios (usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, usuario_id, nombre, tipo, hora_apertura, hora_cierre, duracion_min_minutos, precio_hora, recargo_fin_semana, descuento_volumen, horas_para_descuento;

-- name: UpdateSpace :exec
UPDATE espacios SET nombre=?, tipo=?, hora_apertura=?, hora_cierre=?, duracion_min_minutos=?, precio_hora=?, recargo_fin_semana=?, descuento_volumen=?, horas_para_descuento=?
WHERE id = ?;
