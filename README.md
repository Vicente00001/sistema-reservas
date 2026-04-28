# Proyecto Monolito de Reservas

Sistema de reservas de espacios monolítico en Go con frontend renderizado en el servidor.

## Ejecutar con Docker

1. Clona el repositorio.
2. Ejecuta:

```bash
docker-compose up --build
```

3. Abre `http://localhost:8080`.

## Características

- Go 1.22
- Router `chi`
- SQLite con `modernc.org/sqlite`
- Migraciones `goose`
- Acceso a datos tipado con sqlc
- Autenticación por email/contraseña
- Sesiones con `gorilla/sessions`
- Frontend con plantillas Go, HTMX y CSS estático

## Pruebas

```bash
go test ./... -coverprofile=coverage.out
```

## Variables de entorno

- `PORT` (default `8080`)
- `DB_PATH` (default `data.db`)
- `APP_SECRET` (default `clave-super-secreta`)
