package postgresql

import (
	"database/sql"
	"ejTestTask/config"
)

func CreateBuyer(db *sql.DB, buyer config.Buyer) (int, error) {
	var id int
	query := `INSERT INTO buyers (name, phone) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, buyer.Name, buyer.Phone).Scan(&id)
	return id, err
}

func GetBuyerByID(db *sql.DB, id int) (config.Buyer, error) {
	var buyer config.Buyer
	query := `SELECT id, name, phone FROM buyers WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&buyer.ID, &buyer.Name, &buyer.Phone)
	return buyer, err
}

func UpdateBuyer(db *sql.DB, buyer config.Buyer) error {
	query := `UPDATE buyers SET name = $1, phone = $2 WHERE id = $3`
	_, err := db.Exec(query, buyer.Name, buyer.Phone, buyer.ID)
	return err
}

func DeleteBuyer(db *sql.DB, id int) error {
	query := `DELETE FROM buyers WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
