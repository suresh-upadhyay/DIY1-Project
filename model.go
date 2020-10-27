// model.go

package main

import (
	"database/sql"
	"strconv"

	// tom: errors is removed once functions are implemented
	// "errors"
)


// tom: add backticks to json
type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}


type store struct {
	Store_ID   int     `json:"s_id"`
	Product_ID  string  `json:"p_id"`
	IS_Available bool `json:"is_available"`
}

// tom: these are initial empty definitions
// func (p *product) getProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func (p *product) updateProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func (p *product) deleteProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func (p *product) createProduct(db *sql.DB) error {
//   return errors.New("Not implemented")
// }

// func getProducts(db *sql.DB, start, count int) ([]product, error) {
//   return nil, errors.New("Not implemented")
// }

// tom: these are added after tdd tests
func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}


/////////////////........ DIY........///////////////////

func (s *store) getStoreProduct(db *sql.DB) error {

	return db.QueryRow("SELECT p_id FROM store WHERE s_id=$1 and is_available = $2",
		s.Store_ID , true).Scan(&s.Product_ID)
}


func (s *store) SetStoreProduct(db *sql.DB , productlist []int) error {
	DbQuery := `INSERT INTO store (s_id, p_id, is_available) VALUES `

	for  v := range productlist {
		DbQuery += `(` + strconv.Itoa(s.Store_ID) + `,` + strconv.Itoa(productlist[v])
		DbQuery += "," + "true"

		DbQuery+= `),`
	}
	DbQuery = DbQuery[:len(DbQuery)-1]
	DbQuery += `;`
	_, err := (db.Exec(DbQuery))
	return err
}


