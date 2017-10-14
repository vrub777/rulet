package Repositories

import (
	m "Users/Models"
)

type IRegistratorUserer interface {
	InsertUserData(user m.User) (idUser int)
}
