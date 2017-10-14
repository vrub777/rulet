package Repositories

import (
	m "Users/Models"
)

type IUserer interface {
	UpdateUserCookie(cookie string, idUser int)
	GetUserModelByCookie(cookieHash string) m.User
	GetIdUserByName(name string) int
	GetNameUserById(idUser int) string
	IsEmailInBase(email string) bool
	InsertUserData(user m.User) (idUser int)
	UpdateUserData(user m.User) bool
	GetUserById(idUser int) m.User
}
