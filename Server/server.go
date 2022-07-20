package server

import "net/http"

//	Contiene las rutas del servidor
func New(addr string) *http.Server {
	return &http.Server{
		Addr: addr,
	}
}
