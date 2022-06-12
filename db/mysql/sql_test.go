package mysql

import (
	"log"
	"testing"
)

type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission int    `json:"permission"`
}

func (m *Manager) SelectUserInformationByUsername(username string) (User, error) {
	var user User
	row := m.handler.QueryRow("select id, username, password, email, phone, permission from user where username = ?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
		return User{}, nil
	}
	return user, nil
}

func TestSql(t *testing.T) {
	t.Log(M.SelectUserInformationByUsername("FANG"))
}
