package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func CreateSeller(seller config.Seller) (int, error) {
	var id int
	query := `INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id`
	err := database.PsqlDB.QueryRow(query, seller.Name, seller.Phone).Scan(&id)
	return id, err

}

func GetSellerByID(id int) (config.Seller, error) {
	var seller config.Seller
	query := `SELECT id, name, phone FROM sellers WHERE id = $1`
	err := database.PsqlDB.QueryRow(query, id).Scan(&seller.ID, &seller.Name, &seller.Phone)
	return seller, err
}

func GetSellers() ([]config.Seller, error) {
	var sellers []config.Seller

	query := `SELECT * FROM sellers`
	rows, err := database.PsqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var seller config.Seller

		err = rows.Scan(&seller.ID, &seller.Name, &seller.Phone)
		if err != nil {
			return nil, err
		}
		sellers = append(sellers, seller)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sellers, nil
}

func UpdateSeller(seller config.Seller) error {
	query := `UPDATE sellers SET name = $1, phone = $2 WHERE id = $3`
	_, err := database.PsqlDB.Exec(query, seller.Name, seller.Phone, seller.ID)
	return err
}

func DeleteSeller(id int) error {
	query := `DELETE FROM sellers WHERE id = $1`
	_, err := database.PsqlDB.Exec(query, id)
	return err
}
