package Controllers

import (
	s "Users/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListUsersController struct {
	AuthService
	TemplateHtml
}

func (lu *ListUsersController) ShowPage(c *gin.Context) {
	ListUsersController := ListUsersController{}
	var isAuth = ListUsersController.IsNotAuthRedirectToAuthPage(c)
	if !isAuth {
		return
	}
	userService := s.Userator{}
	var user = userService.GetByRequest(c.Request)
	userRoleService := s.UserRoler{}
	var isAdmin = userRoleService.IsAdmin(user.Id)
	if !isAdmin {
		url := s.URL{}
		c.Redirect(http.StatusFound, url.Get404())
		return
	}
	users := userService.GetListBySearchName("")
	var errors []string
	pageBase := TModelBasePage{Title: "Личный кабинет", Errors: errors}
	page := TModelListUsers{BasePage: pageBase, Title: "Список пользователей", Users: users}
	ListUsersController.showHtml(c.Writer, page, "listUsers")
}

func (lu *ListUsersController) PostAction(c *gin.Context) {
	action := c.Params.ByName("action")
	switch action {
	case "lock":
		PostLockUserById(c)
	case "unlock":
		PostUnLockUserById(c)
	}
}

func PostLockUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	idNum, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(200, gin.H{"status": "", "error": "Некорректные данные клиента"})
	}

	var status = "lock"
	userService := s.Userator{}
	err := userService.LockUserById(idNum)
	if err != "" {
		status = "unlock"
	}

	c.JSON(200, gin.H{"status": status, "error": err})
}

func PostUnLockUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	idNum, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(200, gin.H{"status": "", "error": "Некорректные данные клиента"})
	}

	var status = "unlock"
	userService := s.Userator{}
	err := userService.UnLockUserById(idNum)

	if err != "" {
		status = "lock"
	}

	c.JSON(200, gin.H{"status": status, "error": err})
}
