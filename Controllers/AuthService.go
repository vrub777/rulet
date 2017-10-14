package Controllers

import (
	s "Users/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthService struct {
}

func (as *AuthService) IsNotAuthRedirectToAuthPage(c *gin.Context) bool {
	auth := s.Authorithator{}
	var isAuth = auth.IsAuth(c.Request)
	if !isAuth {
		c.Redirect(http.StatusFound, "/auth/"+c.Request.URL.RequestURI())
		return false
	}

	return true
}
