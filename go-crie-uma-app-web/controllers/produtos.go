package controllers

import (
	produtos "curso-go/go-crie-uma-app-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := produtos.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoParsed, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco", err.Error())
		}

		quantidadeParsed, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade", err.Error())
		}

		produtos.CriaNovoProduto(nome, descricao, precoParsed, quantidadeParsed)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("Id")
	produtos.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoParsed, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco", err.Error())
		}

		quantidadeParsed, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade", err.Error())
		}

		produtos.EditaProduto(id, nome, descricao, precoParsed, quantidadeParsed)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	idDoProduto := r.URL.Query().Get("Id")
	produto := produtos.GetProduto(idDoProduto)
	templates.ExecuteTemplate(w, "Edit", produto)
}
