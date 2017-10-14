package Controllers

import (
	m "Users/Models"
)

type TModelAuthYes struct {
	BasePage  TModelBasePage
	Name      string
	UserModel *m.User
}
