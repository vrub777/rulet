package Repositories

import (
	m "Users/Models"
)

type IUserRoleser interface {
	GetListRoles() *[]m.UserRole
	GetListUserRoles(idUser int) *[]m.UserRole
	AddListRoles(idUser int, listIdRoles []int)
	UpdateListRoles(idUser int, listIdRoles []int)
	IsUserInRole(idUser int, idRole int) bool
}
