package Controllers

import (
	m "Users/Models"
	s "Users/Services"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

type HeadController struct {
	TemplateHtml
}

func (cat *HeadController) ShowHead(c *gin.Context, title string) {
	auth := s.Authorithator{}
	url := s.URL{}
	page := m.ViewHeaderModel{Title: title}
	page.UrlJs = url.GetFullPathJs()
	page.UrlCss = url.GetFullPathCss()
	var isAuth = auth.IsAuth(c.Request)
	if !isAuth {
		cat.showHeader(c.Writer, page, "header-static")
		return
	}

	userService := s.Userator{}
	var user = userService.GetByRequest(c.Request)
	var viewUser = m.ViewUserModel{}
	viewUser.Id = user.Id
	viewUser.Name = user.Name
	viewUser.Rating = user.Rating
	page.User = &viewUser
	cat.showHeader(c.Writer, page, "header-auth")

	/*userRoleService := s.UserRoler{}
	var isAvailableEnter = userRoleService.IsAdmin(user.Id) || userRoleService.IsOperator(user.Id)
	if !isAvailableEnter {
		return
	}*/
}

func (cat *HeadController) IsAuth() {

}
