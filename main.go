package main

import (
	"net/http"
	//"text/template" //Biblioteca para se criar template
	//_ "github.com/lib/pq" //biblioteca de postgres o _ se refe quando formos utlizar em tempo de execução a biblioteca
	"Loja_Em_GO/routes"
)

func main() {
	routes.CarregaRotas()
	//Estamos falando que vai abrir na porta 8000 do nosso localhost
	http.ListenAndServe(":8000", nil)

}
