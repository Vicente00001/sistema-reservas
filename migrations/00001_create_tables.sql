-- +goose Up
CREATE TABLE usuarios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE espacios (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    usuario_id INTEGER NOT NULL REFERENCES usuarios(id),
    nombre TEXT NOT NULL,
    tipo TEXT NOT NULL,
    hora_apertura TEXT NOT NULL,
    hora_cierre TEXT NOT NULL,
    duracion_min_minutos INTEGER NOT NULL,
    precio_hora REAL NOT NULL,
    recargo_fin_semana REAL NOT NULL DEFAULT 0,
    descuento_volumen REAL NOT NULL DEFAULT 0,
    horas_para_descuento INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE reservas (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    espacio_id INTEGER NOT NULL REFERENCES espacios(id),
    fecha TEXT NOT NULL,
    hora_inicio TEXT NOT NULL,
    hora_fin TEXT NOT NULL,
    precio_total REAL NOT NULL,
    estado TEXT NOT NULL DEFAULT 'confirmada'
);

-- +goose Down
DROP TABLE reservas;
DROP TABLE espacios;
DROP TABLE usuarios;
