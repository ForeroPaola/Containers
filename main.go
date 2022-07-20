package main

import (
	"fmt"
	"net/http"
)

func main() {

	//Configuración de rutas
	//Especifica la ruta y la función a ejecutar
	//http.ResponseWriter donde vamos a enviar las respuestas
	//http.Request Recibe todo lo que viene desde fuera
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Validamos tipo de metodo que utiliza el usuario
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
		fmt.Fprintf(w, "Hola mundo %s", "PAOLA")
	})
	//Ayuda a configurar el servidor dirección y pueto donde s
	//Ejecuta API
	srv := http.Server{
		Addr: ":8080", //Donde corre la APP
	}
	//Servidor va a iniciar, retorna un error en caso que ocurra
	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
