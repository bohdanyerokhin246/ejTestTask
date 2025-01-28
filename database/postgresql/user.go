package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func CreateUser(user config.User) error {

	query := `INSERT INTO users (login, password, role) VALUES ($1, $2, $3)`
	err := database.PsqlDB.QueryRow(query, user.Login, user.Password, user.Role).Err()
	return err

}

func GetUserRole(login, password string) (string, error) {
	var user config.User
	query := `SELECT role FROM users WHERE login = $1 AND password = $2`
	err := database.PsqlDB.QueryRow(query, login, password).Scan(&user.Role)
	return user.Role, err
}
