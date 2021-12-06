package models

import "fundamentos_web/database"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := database.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id asc")
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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos

}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := database.ConectaComBancoDeDados()

	insereDadosnoBanco, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosnoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func DeletaProduto(id string) {
	db := database.ConectaComBancoDeDados()
	deletaProduto, err := db.Prepare("DELETE FROM produtos WHERE id= $1")
	if err != nil {
		panic(err.Error())
	}
	deletaProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := database.ConectaComBancoDeDados()
	produtoDoBanco, err := db.Query("SELECT * FROM PRODUTOS WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}
	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := database.ConectaComBancoDeDados()
	AtualizaProduto, err := db.Prepare("UPDATE produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
