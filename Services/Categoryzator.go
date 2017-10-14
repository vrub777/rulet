package Services

import (
	m "Users/Models"
	r "Users/Repositories"
	"strconv"
)

type Categoryzator struct {
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
	}

	viewListCategores.ListFirstLavelCategory = firstLavelCategores
	return viewListCategores
}

/*func (cat *Categoryzator) GetListSecondLavel(idParent int) {
	rep := r.ICategorer(&r.Categores{})
	rep.GetListSecondCategores(idParent)
}*/

func (cat *Categoryzator) UpdateFirstCategory(categoryModel m.UpdateFirstCategoryModel) {
	rep := r.ICategorer(&r.Categores{})
	rep.UpdateFirstCategory(categoryModel)
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

func (cat *Categoryzator) GetIcoFilePathByIdCategory(id int) string {
	idStr := strconv.Itoa(id)
	return "./Img/Categores/" + idStr + ".jpg" //TODO Сдулать класс для путей, перенести, пользоваться им
}

func (cat *Categoryzator) UpdateIcoNameInCategory(idCategory int, icoFileName string) {
	rep := r.ICategorer(&r.Categores{})
	rep.UpdateIcoNameInCategory(idCategory, icoFileName)
}
