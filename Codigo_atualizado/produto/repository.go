package produto

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

func (pr *ProdutoRepository) List() ([]Produto, error) {
	rows, err := pr.db.Query(`SELECT id, name, preco, descricao, categoria FROM Produtos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []Produto

	for rows.Next() {
		var produto Produto
		err = rows.Scan(&produto.Id, &produto.Name, &produto.Preco, &produto.Descricao, &produto.Categoria)
		if err != nil {
			return nil, err
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}

func (pr *ProdutoRepository) Get(id int) (*Produto, error) {
	row := pr.db.QueryRow(`
		SELECT id, name, preco, descricao, categoria
		FROM Produtos
		WHERE id = ?`, id)

	var produto Produto
	err := row.Scan(&produto.Id, &produto.Name, &produto.Preco, &produto.Descricao, &produto.Categoria)
	if err != nil {
		return nil, err
	}

	return &produto, nil
}

func (pr *ProdutoRepository) Create(produto Produto) (int64, error) {
	result, err := pr.db.Exec(`INSERT INTO Produtos(name, preco, descricao, categoria)
					  VALUES (?, ?, ?, ?)`,
		produto.Name, produto.Preco, produto.Descricao, produto.Categoria)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pr *ProdutoRepository) Update(id int, produto Produto) error {
	_, err := pr.db.Exec(`UPDATE Produtos
						SET name=?,
						    preco=?,
						    descricao=?,
						    categoria=?
						WHERE id=?`,
		produto.Name, produto.Preco, produto.Descricao, produto.Categoria, id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProdutoRepository) Delete(id int) error {
	_, err := pr.db.Exec(`DELETE
							FROM Produtos
							WHERE id = ?`, id)

	if err != nil {
		return err
	}

	return nil
}
