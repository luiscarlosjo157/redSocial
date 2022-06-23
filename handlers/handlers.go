package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/luiscarlosjo157/redSocial/middlew"
	"github.com/luiscarlosjo157/redSocial/routers"
	"github.com/gorilla/mux"	
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pogo a escuchar al Servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Reguistro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))

}