package Services

import (
	m "Users/Models"
	r "Users/Repositories"
	"strconv"
)

type Categoryzator struct {
}

func (cat *Categoryzator) GetListFirstCategores() []*m.ViewFirstCategory {
	categoresRepository := r.ICategorer(&r.Categores{})
	listCategores := categoresRepository.GetListFirstCategores()
	return listCategores
}

func (cat *Categoryzator) GetListSecondCategores(idParent int) []*m.ViewSecondCategory {
	categoresRepository := r.ICategorer(&r.Categores{})
	listCategores := categoresRepository.GetListSecondCategores(idParent)
	return listCategores
}

func (cat *Categoryzator) GetViewListCategores() m.ViewListCategory {
	viewListCategores := m.ViewListCategory{}
	categoresRepository := r.ICategorer(&r.Categores{})
	firstLavelCategores := categoresRepository.GetListFirstCategores()

	for _, value := range firstLavelCategores {
		if value.IcoFileName != "" {
			value.IcoFullPath = cat.GetIcoFilePathByIdCategory(value.Id)
		}
		value.ListSecondLavelCategory = categoresRepository.GetListSecondCategores(value.Id)
		value.CountSecondLevel = len(value.ListSecondLavelCategory)
	}

	viewListCategores.ListFirstLavelCategory = firstLavelCategores
	return viewListCategores
}

func (cat *Categoryzator) UpdateFirstCategory(categoryModel m.UpdateFirstCategoryModel) {
	rep := r.ICategorer(&r.Categores{})
	rep.UpdateFirstCategory(categoryModel)
}

func (cat *Categoryzator) UpdateSecondCategory(categoryModel m.UpdateSecondCategoryModel) {
	rep := r.ICategorer(&r.Categores{})
	rep.UpdateSecondCategory(categoryModel)
}

func (cat *Categoryzator) GetIcoUrlByIdCategory(id int) string {
	url := URL{}
	urlPath := url.GetPathCatalogIcons()

	rep := r.ICategorer(&r.Categores{})
	nameCategory := rep.GetNameIcoById(id)

	if nameCategory == "" {
		return ""
	}

	icoUrl := urlPath + "/" + nameCategory
	return icoUrl
}

func (cat *Categoryzator) AddCategory(category m.AddCategoryModel) int {
	rep := r.ICategorer(&r.Categores{})
	newId := rep.AddCategory(category)
	return newId
}
func (cat *Categoryzator) DeleteCategory(id int) {
	rep := r.ICategorer(&r.Categores{})
	rep.DeleteCategory(id)
}
func (cat *Categoryzator) GetIcoFilePathByIdCategory(id int) string {
	idStr := strconv.Itoa(id)
	return "./Img/Categores/" + idStr + ".jpg" //TODO Сдулать класс для путей, перенести, пользоваться им
}

func (cat *Categoryzator) UpdateIcoNameInCategory(idCategory int, icoFileName string) {
	rep := r.ICategorer(&r.Categores{})
	rep.UpdateIcoNameInCategory(idCategory, icoFileName)
}
