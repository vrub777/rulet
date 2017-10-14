package Services

import (
	m "Users/Models"
	r "Users/Repositories"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

type Userator struct {
}

var CacheRepository *cache.Cache

func (u *Userator) GetByRequest(req *http.Request) *m.User {
	if CacheRepository == nil {
		CacheRepository = cache.New(5*time.Minute, 60*time.Second)
	}

	cookierator := Cookierator{}
	cookieValue := cookierator.GetCookie(req)
	cacheKeyValue := "oneUserModel_" + cookieValue
	iUser, _ := CacheRepository.Get(cacheKeyValue)
	if iUser == nil {
		userRepository := r.IUserer(&r.User{})
		userModel := userRepository.GetUserModelByCookie(cookieValue)
		CacheRepository.Set(cacheKeyValue, &userModel, cache.DefaultExpiration)
		return &userModel
	}
	userModel := iUser.(*m.User)
	return userModel
}

func (u *Userator) GetById(idUser int) m.User {
	userRepository := r.IUserer(&r.User{})
	userModel := userRepository.GetUserById(idUser)
	return userModel
}

func (u *Userator) GetListBySearchName(searchString string) []*m.User {
	userRepository := r.User{}
	listUsers := userRepository.SelectSliceUsers(searchString)
	return listUsers
}

func (u *Userator) GetUserIdByName(name string) int {
	userRepository := r.IUserer(&r.User{})
	idUser := userRepository.GetIdUserByName(name)
	return idUser
}

func (u *Userator) IsEmailInSystem(email string) bool {
	userRepository := r.IUserer(&r.User{})
	isYes := userRepository.IsEmailInBase(email)
	return isYes
}

func (u *Userator) UpdateUser(postModel m.UserAdd) bool {
	user := m.User{}
	auth := Authorithator{}
	user.Id = postModel.Id
	user.Name = postModel.Name
	user.Email = postModel.Email
	user.PassworsHash = auth.GetHashPassword(postModel.Password)

	userRepository := r.IUserer(&r.User{})
	isYes := userRepository.UpdateUserData(user)
	return isYes
}

func (u *Userator) LockUserById(idUser int) (err string) {
	userRepository := r.User{}
	isLock := userRepository.IsLock(idUser)
	if isLock {
		return "Пользователь уже заблокирован"
	}

	isOk := userRepository.SetLockStatus(idUser)
	if !isOk {
		return "Пользователь не заблокирован, ошибка на сервере"
	}

	return ""
}

func (u *Userator) UnLockUserById(idUser int) (err string) {
	userRepository := r.User{}
	isLock := userRepository.IsLock(idUser)
	if !isLock {
		return "Пользователь уже разблокирован"
	}

	isOk := userRepository.SetUnLockStatus(idUser)
	if !isOk {
		return "Пользователь не разблокирован, ошибка на сервере"
	}

	return ""
}
