package Services

import (
	m "Users/Models"
	r "Users/Repositories"
	"github.com/gin-gonic/gin"
	"time"
)

type Registrator struct {
}

func (reg *Registrator) Registration(postModel m.UserAdd) int {
	idUser := reg.registration(postModel)
	return idUser
}

func (reg *Registrator) RegistrationAndAfterAuth(responseWriter gin.ResponseWriter,
	postModel m.UserAdd) {
	idUser := reg.registration(postModel)
	if idUser <= 0 {
		return
	}

	auth := Authorithator{}
	auth.AuthorithationByIdUser(responseWriter, idUser)
}

func (reg *Registrator) registration(postModel m.UserAdd) int {
	user := m.User{}
	auth := Authorithator{}
	user.Name = postModel.Name
	user.Email = postModel.Email
	user.DateRegistration = time.Now()
	user.DateActivation = time.Now()
	user.IsLocked = false
	user.IsActivate = true
	user.CountTryAuth = 0
	user.PassworsHash = auth.GetHashPassword(postModel.Password)

	userData := r.IUserer(&r.User{})
	idUser := userData.InsertUserData(user)
	return idUser
}
