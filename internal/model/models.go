package model

type ViewData struct {
	Title   string
	UserID  int64
	Flash   string
	Error   string
	Payload any
}

type SpaceForm struct {
	Nombre             string
	Tipo               string
	HoraApertura       string
	HoraCierre         string
	DuracionMinMinutos int32
	PrecioHora         float64
	RecargoFinSemana   float64
	DescuentoVolumen   float64
	HorasParaDescuento int32
}

type ReservaForm struct {
	Fecha      string
	HoraInicio string
	HoraFin    string
}
