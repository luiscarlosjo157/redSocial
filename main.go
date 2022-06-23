package main
import (
	"log"
	"github.com/luiscarlosjo157/redSocial/handlers"
	"github.com/luiscarlosjo157/redSocial/bd"
)

func main(){
	if bd.ChequeoConnection()==0 {
		log.Fatal("sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}