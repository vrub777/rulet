package Controllers

import (
	m "Users/Models"
)

type TModelListUsers struct {
	BasePage TModelBasePage
	Title    string
	Users    []*m.User
}
