package Controllers

import (
	m "Users/Models"
	s "Users/Services"
	"github.com/gin-gonic/gin"
	//"github.com/patrickmn/go-cache"
	"regexp"
	"unicode/utf8"
)

type RegistrationController struct {
	TModelBasePage
	TemplateHtml
}

func (ac *RegistrationController) ShowPageRegistration(c *gin.Context) {
	RegistrationController := RegistrationController{}
	var errors []string
	pageBase := TModelBasePage{Title: "Регистрация пользователя", Errors: errors}
	page := TModelAuthYes{BasePage: pageBase}
	RegistrationController.showHtml(c.Writer, page, "reg")
}

func (ac *RegistrationController) PostPageRegistration(c *gin.Context) {
	name := c.Request.PostFormValue("name")
	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")
	repeatPassword := c.Request.PostFormValue("repeatpassword")

	//region := c.Request.PostFormValue("region")

	var errors []string
	if utf8.RuneCountInString(name) < 3 {
		errors = append(errors, "Имя пользователя слишком короткое")
	}
	if utf8.RuneCountInString(name) > 50 {
		errors = append(errors, "Имя пользователя слишком длинное")
	}
	if utf8.RuneCountInString(password) < 3 {
		errors = append(errors, "Пароль слишком короткий")
	}
	if password != repeatPassword {
		errors = append(errors, "Пароли не совпадают. Пожалуйста, повторите ввод пароля.")
	}
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
		errors = append(errors, "Вы ввели неправильный адрес электронной почты. Пожалуйста, введите адрес корректно")
	}

	userService := s.Userator{}
	idUser := userService.GetUserIdByName(name)
	if idUser > 0 {
		errors = append(errors, "Такой пользователь уже существует, пожалуйста, выберите другой логин")
	}

	isEmailConflict := userService.IsEmailInSystem(email)
	if isEmailConflict {
		errors = append(errors, "Такой почтовый адрес уже существует, пожалуйста, выберите другую электронную почту")
	}

	if len(errors) > 0 {
		showErrorAfterSave(c.Writer, errors)
		return
	}

	userPostModel := m.UserAdd{Name: name, Email: email, Password: password}
	reg := s.Registrator{}
	reg.RegistrationAndAfterAuth(c.Writer, userPostModel)
	showSucessAfterSave(c.Writer)
}

func showSucessAfterSave(responseWriter gin.ResponseWriter) {
	title := "Регистрация успешно началась. Пожалуйста, закончите её - перейдите по ссылке в письме"
	page := TModelHeader{Title: title}
	RegistrationController := RegistrationController{}
	RegistrationController.showHtmlWithHeader(responseWriter, page, "afterreg")
}

func showErrorAfterSave(responseWriter gin.ResponseWriter, errors []string) {
	title := "Регистрация не завершена. Пожалуйста, заполните поля заново."
	pageBase := TModelBasePage{Title: title, Errors: errors}
	page := TModelAuthYes{BasePage: pageBase}
	RegistrationController := RegistrationController{}
	RegistrationController.showHtml(responseWriter, page, "reg")
}
