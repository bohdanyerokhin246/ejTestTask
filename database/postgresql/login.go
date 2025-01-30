package postgresql

import (
	"ejTestTask/config"
	"ejTestTask/database"
)

func GetUserRole(login, password string) (config.User, error) {
	var user config.User
	query := `SELECT login, role FROM users WHERE login = $1 AND password = $2`
	err := database.PsqlDB.QueryRow(query, login, password).Scan(&user.Login, &user.Role)
	return user, err
}
