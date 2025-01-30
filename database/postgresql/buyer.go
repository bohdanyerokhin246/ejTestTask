package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func CreateBuyer(buyer config.Buyer) (int, error) {
	var id int
	query := `INSERT INTO buyers (name, phone) VALUES ($1, $2) RETURNING id`
	err := database.PsqlDB.QueryRow(query, buyer.Name, buyer.Phone).Scan(&id)
	return id, err
}

func GetBuyerByID(id int) (config.Buyer, error) {
	var buyer config.Buyer
	query := `SELECT id, name, phone FROM buyers WHERE id = $1`
	err := database.PsqlDB.QueryRow(query, id).Scan(&buyer.ID, &buyer.Name, &buyer.Phone)
	return buyer, err
}

func GetBuyers() ([]config.Buyer, error) {
	var buyers []config.Buyer

	query := `SELECT * FROM buyers`
	rows, err := database.PsqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buyer config.Buyer

		err = rows.Scan(&buyer.ID, &buyer.Name, &buyer.Phone)
		if err != nil {
			return nil, err
		}
		buyers = append(buyers, buyer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return buyers, nil
}

func UpdateBuyer(buyer config.Buyer) error {
	query := `UPDATE buyers SET name = $1, phone = $2 WHERE id = $3`
	_, err := database.PsqlDB.Exec(query, buyer.Name, buyer.Phone, buyer.ID)
	return err
}

func DeleteBuyer(id int) error {
	query := `DELETE FROM buyers WHERE id = $1`
	_, err := database.PsqlDB.Exec(query, id)
	return err
}
