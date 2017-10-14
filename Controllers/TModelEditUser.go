package Controllers

import (
	m "Users/Models"
)

type TModelEditUser struct {
	HeaderPage TModelHeader
	UserRoles  *[]m.UserRole
	UserAdd    m.UserAdd
	Errors     []string
}
