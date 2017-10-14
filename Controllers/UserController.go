package Controllers

import (
	m "Users/Models"
	s "Users/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
	AuthService
	TemplateHtml
}

func (u *UserController) ShowAddUser(c *gin.Context) {
	u.isAvailableEnter(c)
	userAdd := m.UserAdd{}
	userAdd.IsAddUser = true
	userAdd.NameOkButton = "Добавить"
	userAdd.NameAction = "addUser"

	userRoleService := s.UserRoler{}
	title := "Добавить пользователя"
	pageHeader := TModelHeader{Title: title}
	var listRoles = userRoleService.GetListUserRoles()
	page := TModelEditUser{HeaderPage: pageHeader, UserAdd: userAdd, UserRoles: listRoles}
	u.showHtmlWithHeader(c.Writer, page, "addUser")
}

func (u *UserController) PostAddUser(c *gin.Context) {
	name := c.Request.PostFormValue("name")
	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")

	userRoleService := s.UserRoler{}
	roles := userRoleService.GetListUserRoles()
	var listRolesPost []m.UserRole
	var idsCheckRole []int

	for _, role := range *roles {
		valuePost := c.Request.PostFormValue("role-" + strconv.Itoa(role.Id))
		id, _ := strconv.Atoi(valuePost)
		if id == role.Id {
			role.IsCheck = true
			idsCheckRole = append(idsCheckRole, id)
		}
		listRolesPost = append(listRolesPost, role)
	}

	var errors = u.validateForm(c)

	userAdd := m.UserAdd{}
	userAdd.Name = name
	userAdd.Email = email
	userAdd.Password = password
	userAdd.IdsRole = idsCheckRole
	userAdd.IsAddUser = false
	userAdd.NameOkButton = "Сохранить"
	userAdd.NameAction = "addUser"

	if len(errors) > 0 {
		u.showErrorAfterSave(c.Writer, userAdd, listRolesPost, errors)
		return
	}

	reg := s.Registrator{}
	idUser := reg.Registration(userAdd)
	if idUser <= 0 {
		errors = append(errors, "Зарегистрировать пользователя не удалось. Сервер не доступен.")
		u.showErrorAfterSave(c.Writer, userAdd, listRolesPost, errors)
		return
	}

	userRoleService.AddListRolesForUser(idUser, idsCheckRole)

	url := s.URL{}
	c.Redirect(http.StatusFound, url.GetHostNameWithProtocol()+"/listusers")
}

func (u *UserController) showErrorAfterSave(responseWriter gin.ResponseWriter,
	userAdd m.UserAdd, listRoles []m.UserRole, errors []string) {
	title := "Добавить пользователя"
	pageHeader := TModelHeader{Title: title}
	page := TModelEditUser{HeaderPage: pageHeader, UserRoles: &listRoles,
		UserAdd: userAdd, Errors: errors}
	u.showHtmlWithHeader(responseWriter, page, "addUser")
}

func (u *UserController) ShowEditUser(c *gin.Context) {
	u.isAvailableEnter(c)
	idUser, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		return
	}
	userRoleService := s.UserRoler{}
	roles := userRoleService.GetListUserRoles()
	userRoles := userRoleService.GetUserRolesById(idUser)
	var listRoles []m.UserRole

	for _, role := range *roles {
		for _, userRole := range *userRoles {
			if role.Id == userRole.Id {
				role.IsCheck = true
			}
		}
		listRoles = append(listRoles, role)
	}

	// TODO Id нужен и на указатель перевести
	userService := s.Userator{}
	var user = userService.GetById(idUser)
	userAdd := m.UserAdd{}
	userAdd.Id = idUser
	userAdd.Name = user.Name
	userAdd.Email = user.Email
	userAdd.Roles = listRoles
	userAdd.NameOkButton = "Сохранить"
	userAdd.NameAction = "/editUser/" + strconv.Itoa(idUser)

	var errors []string
	title := "Редактировать пользователя"
	pageHeader := TModelHeader{Title: title}
	page := TModelEditUser{HeaderPage: pageHeader, UserRoles: &listRoles,
		UserAdd: userAdd, Errors: errors}
	u.showHtmlWithHeader(c.Writer, page, "addUser")
}

func (u *UserController) PostEditUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.PostFormValue("id"))
	if err != nil {
		return
	}
	name := c.Request.PostFormValue("name")
	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")

	userRoleService := s.UserRoler{}
	roles := userRoleService.GetListUserRoles()
	var listRolesPost []m.UserRole
	var idsCheckRole []int

	for _, role := range *roles {
		valuePost := c.Request.PostFormValue("role-" + strconv.Itoa(role.Id))
		id, _ := strconv.Atoi(valuePost)
		if id == role.Id {
			role.IsCheck = true
			idsCheckRole = append(idsCheckRole, id)
		}
		listRolesPost = append(listRolesPost, role)
	}

	var errors = u.validateForm(c)

	userAdd := m.UserAdd{}
	userAdd.Id = id
	userAdd.Name = name
	userAdd.Email = email
	userAdd.Password = password
	userAdd.IdsRole = idsCheckRole
	userAdd.IsAddUser = false
	userAdd.NameOkButton = "Сохранить"
	userAdd.NameAction = "editUser"

	if len(errors) > 0 {
		u.showErrorAfterSave(c.Writer, userAdd, listRolesPost, errors)
		return
	}

	userService := s.Userator{}
	isYes := userService.UpdateUser(userAdd)
	if !isYes {
		errors = append(errors, "Зарегистрировать пользователя не удалось. Сервер не доступен.")
		u.showErrorAfterSave(c.Writer, userAdd, listRolesPost, errors)
		return
	}

	userRoleService.UpdateListRolesForUser(userAdd.Id, idsCheckRole)

	url := s.URL{}
	c.Redirect(http.StatusFound, url.GetHostNameWithProtocol()+"/listusers")
}

func (u *UserController) isAvailableEnter(c *gin.Context) {
	var isAuth = u.IsNotAuthRedirectToAuthPage(c)
	if !isAuth {
		return
	}

	userService := s.Userator{}
	var user = userService.GetByRequest(c.Request)
	userRoleService := s.UserRoler{}
	var isAvailableEnter = userRoleService.IsAdmin(user.Id) || userRoleService.IsOperator(user.Id)
	if !isAvailableEnter {
		return
	}
}

func (u *UserController) validateForm(c *gin.Context) []string {
	name := c.Request.PostFormValue("name")
	email := c.Request.PostFormValue("email")
	password := c.Request.PostFormValue("password")

	var errors []string
	if len(name) < 3 {
		errors = append(errors, "Имя пользователя слишком короткое")
	}
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
		errors = append(errors, "Адрес электронной почты введён некорректно")
	}
	if len(password) < 3 {
		errors = append(errors, "Пароль слишком короткий")
	}
	return errors
}
