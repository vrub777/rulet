package Services

import (
	m "Users/Models"
	r "Users/Repositories"
	"github.com/patrickmn/go-cache"
	"time"
)

type UserRoler struct {
}

func (ur *UserRoler) GetListUserRoles() *[]m.UserRole {
	if CacheRepository == nil {
		CacheRepository = cache.New(30*time.Minute, 60*time.Second)
	}

	cacheKeyValue := "userRolesModel"
	iRoles, _ := CacheRepository.Get(cacheKeyValue)
	if iRoles == nil {
		userRoleRepository := r.IUserRoleser(&r.UserRole{})
		listUserRoles := userRoleRepository.GetListRoles()
		CacheRepository.Set(cacheKeyValue, listUserRoles, cache.DefaultExpiration)
		return listUserRoles
	}
	listUserRoles := iRoles.(*[]m.UserRole)
	return listUserRoles
}

func (ur *UserRoler) GetUserRolesById(idUser int) *[]m.UserRole {
	userRoleRepository := r.IUserRoleser(&r.UserRole{})
	listUserRoles := userRoleRepository.GetListUserRoles(idUser)
	return listUserRoles
}

func (ur *UserRoler) IsAdmin(idUser int) bool {
	isAdmin := isItRole(idUser, "Admin")
	return isAdmin
}

func (ur *UserRoler) IsOperator(idUser int) bool {
	isOperator := isItRole(idUser, "Operator")
	return isOperator
}

func (ur *UserRoler) AddListRolesForUser(idUser int, ids []int) {
	userRoleRepository := r.IUserRoleser(&r.UserRole{})
	userRoleRepository.AddListRoles(idUser, ids)
}

func (ur *UserRoler) UpdateListRolesForUser(idUser int, ids []int) {
	userRoleRepository := r.IUserRoleser(&r.UserRole{})
	userRoleRepository.UpdateListRoles(idUser, ids)
}

func isItRole(idUser int, nameRole string) bool {
	repositoryRole := r.UserRole{}
	var listUserRoles = repositoryRole.GetListUserRoles(idUser)
	for _, role := range *listUserRoles {
		if role.Name == nameRole {
			return true
		}
	}
	return false
}
