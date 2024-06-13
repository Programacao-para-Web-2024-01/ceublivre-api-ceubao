package products

import (
	"database/sql"
)

type Products struct {
	Id                    int
	Name                  string
	Description           string
	Price                 float64
	Categoria_idcategoria int
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) List() ([]Product, error) {
	rows, err := pr.db.Query(`
		SELECT id,
		       name,
		       description,
		       price,
		       categoria_idcategoria
		FROM products`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Categoria_idcategoria,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) Get(id int) (*Product, error) {
	row := pr.db.QueryRow(`
		SELECT id, name, description, price, categoria_idcategoria
		FROM products
		WHERE id = ?`, id)

	var product Product
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Categoria_idcategoria)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) Create(product Product) (int64, error) {
	result, err := pr.db.Exec(`INSERT INTO products(name, description, price, categoria_idcategoria)
					  VALUES (?, ?, ?, ?)`,
		product.Name, product.Description, product.Price, product.Categoria_idcategoria)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) Update(id int, product Product) error {
	_, err := pr.db.Exec(`UPDATE products
						SET name=?,
						    description=?,
						    price=?,
						    categoria_idcategoria=?
						WHERE id=?`,
		product.Name, product.Description, product.Price, product.Categoria_idcategoria, id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) Delete(id int) error {
	_, err := pr.db.Exec(`DELETE
							FROM products
							WHERE id = ?`, id)

	if err != nil {
		return err
	}

	return nil
}