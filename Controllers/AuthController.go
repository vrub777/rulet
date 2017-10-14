package Controllers

import (
	s "Users/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	TemplateHtml
}

func (ac *AuthController) ShowSimpleAuthUser(c *gin.Context) {
	auth := s.Authorithator{}
	var isAuth = auth.IsAuth(c.Request)
	if isAuth {
		ShowYesForm(c)
		return
	}
	var errors []string
	ShowFormAuth(c, errors)
}

func (ac *AuthController) PostSimpleAuthUser(c *gin.Context) {
	name := c.Request.PostFormValue("name")
	password := c.Request.PostFormValue("password")

	var errors []string
	if GetCountTry(name) >= 10 {
		errors = append(errors, "Число попыток ввода ограничено, попробуйте через 1 час")
	}

	if len(errors) == 0 && !IsNameUserInSystem(name) {
		errors = append(errors, "Такого пользователя нет в системе")
	} else {
		auth := s.Authorithator{}
		var isNotValidPassord bool = auth.IsValidPasswordForUser(name, password)
		if !isNotValidPassord {
			errors = append(errors, "Пароль неверный, пожалуйста, повторите ввод")
		}
	}

	if len(errors) > 0 {
		ShowFormAuth(c, errors)
		return
	}
	fmt.Printf("START \n")
	auth := s.Authorithator{}
	auth.Authorithation(c.Writer, name, password)

	backURL := c.Request.PostFormValue("backURL")
	if backURL == "" {
		backURL = "auth"
	}
	fmt.Printf(" ISAUTH ---- %s ------- \n", "http://"+c.Request.Host+"/"+backURL)
	url := s.URL{}
	c.Redirect(http.StatusFound, url.GetHostNameWithProtocol()+"/"+backURL)
}

func (ac *AuthController) OutAuthUser(c *gin.Context) {
	auth := s.Authorithator{}
	auth.Out(c.Writer)
	url := s.URL{}
	c.Redirect(http.StatusFound, url.GetHostNameWithProtocol()+"/auth")
}

func ShowYesForm(c *gin.Context) {
	AuthController := AuthController{}
	var errors []string
	userService := s.Userator{}
	user := userService.GetByRequest(c.Request)
	pageBase := TModelBasePage{Title: "Личный кабинет", Errors: errors}
	page := TModelAuthYes{BasePage: pageBase, Name: user.Name, UserModel: user}
	AuthController.showHtml(c.Writer, page, "authOk")
}

func ShowFormAuth(c *gin.Context, errors []string) {
	AuthController := AuthController{}
	pageBase := TModelBasePage{Title: "Авторизация", Errors: errors}
	page := TModelAuth{BasePage: pageBase, BackURL: c.Params.ByName("backURL")}
	AuthController.showHtml(c.Writer, page, "auth")
}

func IsNameUserInSystem(nameUser string) bool {
	auth := s.Authorithator{}
	return auth.IsNameUserInSystem(nameUser)
}

func GetCountTry(nameUser string) int {
	auth := s.Authorithator{}
	return auth.GetCountTryAuth(nameUser)
}
