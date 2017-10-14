package Repositories

import (
	m "Users/Models"
)

type ICategorer interface {
	GetListFirstCategores() []*m.ViewFirstCategory
	GetListSecondCategores(idParent int) []*m.ViewSecondCategory
	UpdateFirstCategory(m.UpdateFirstCategoryModel)
	UpdateIcoNameInCategory(idCategory int, icoFileName string)
	GetNameIcoById(idCategory int) string
}
