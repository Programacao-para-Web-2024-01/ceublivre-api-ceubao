package Produto

import (
	"database/sql"
)

type ProdutoRepository struct {
	db *sql.DB
}

func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{
		db: db,
	}
}

func (sr *ProdutoRepository) List() ([]Produto, error) {
	rows, err := sr.db.Query(`SELECT id, name, preco, descricao, categoria FROM Produtos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []Produto

	for rows.Next() {
		var produto Produto
		err = rows.Scan(&Produto.Id, &Produto.Name, &Produto.Preco, &Produto.Descricao, &Produto.Categoria)
		if err != nil {
			return nil, err
		}

		Produtos = append(Produtos, Produto)
	}

	return Produtos, nil
}

func (sr *ProdutoRepository) Get(id int) (*Produto, error) {
	row := sr.db.QueryRow(`
		SELECT id, name, preco, descricao, categoria
		FROM Produtos
		WHERE id = ?`, id)

	var Produto Produto
	err := row.Scan(&Produto.Id, &Produto.Name, &Produto.Preco, &Produto.Descricao, &Produto.Categoria)
	if err != nil {
		return nil, err
	}

	return &Produto, nil
}

func (sr *ProdutoRepository) Create(Produto Produto) (int64, error) {
	result, err := sr.db.Exec(`INSERT INTO Produtos(name, preco, descricao, categoria)
					  VALUES (?, ?, ?, ?)`,
		Produto.Name, Produto.Preco, Produto.Descricao, Produto.Categoria)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (sr *ProdutoRepository) Update(id int, Produto Produto) error {
	_, err := sr.db.Exec(`UPDATE Produtos
						SET name=?,
						    preco=?,
						    descricao=?,
						    categoria=?
						WHERE id=?`,
		Produto.Name, Produto.Preco, Produto.Descricao, Produto.Categoria, id)

	if err != nil {
		return err
	}

	return nil
}

func (sr *ProdutoRepository) Delete(id int) error {
	_, err := sr.db.Exec(`DELETE
							FROM Produtos
							WHERE id = ?`, id)

	if err != nil {
		return err
	}

	return nil
}
