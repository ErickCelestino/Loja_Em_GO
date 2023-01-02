package controllers

import (
	"Loja_Em_GO/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Variavel que pega todos os arquivos html e coloca no nosso template
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Função que pega os dados dos forms e converte eles
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço:", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da Quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDOProuto := r.URL.Query().Get("id")
	models.DeletaProduto(idDOProuto)
	http.Redirect(w, r, "/", 301)
}
