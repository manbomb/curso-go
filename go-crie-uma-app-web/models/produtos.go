package produtos

import "curso-go/go-crie-uma-app-web/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeProdutos.Scan(
			&id,
			&nome,
			&descricao,
			&preco,
			&quantidade,
		)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Quantidade = quantidade
		p.Descricao = descricao
		p.Nome = nome
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
}

func GetProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	produtoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	return produto
}

func EditaProduto(id, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	editaDadosNoBanco, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	editaDadosNoBanco.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
