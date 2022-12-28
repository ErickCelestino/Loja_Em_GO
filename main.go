package main

import (
	"database/sql"
	"net/http"
	"text/template" //Biblioteca para se criar templates

	_ "github.com/lib/pq" //biblioteca de postgres o _ se refe quando formos utlizar em tempo de execução a biblioteca
)

func ConectaComBancoDeDados() *sql.DB {
	//Conecxão com o banco deve passar as informações que pegamos depois que criamos o banco
	conexao := "user=postgres dbname=AluraLoja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// Variavel que pega todos os arquivos html e coloca no nosso template
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	//Estamos falando que vai abrir na porta 8000 do nosso localhost
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	db := ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
