package Repositories

import (
	m "Users/Models"
	connect "Users/Repositories/Connector"
	"database/sql"
	_ "github.com/lib/pq"
)

type Categores struct {
}

func (cat *Categores) GetListFirstCategores() []*m.ViewFirstCategory {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	categores := []*m.ViewFirstCategory{}
	rows, err := db.Query(`select id, name, count_request, name_file_ico from catalog."tFirstLavelCategory" 
						   order by id`)

	if err != nil {
		panic(err)
	}

	var countRequest sql.NullInt64
	var icoFileName sql.NullString
	for rows.Next() {
		category := m.ViewFirstCategory{}
		// TODO Решить будет Null у IcoPath или нет
		err = rows.Scan(&category.Id, &category.Name, &countRequest, &icoFileName)
		if countRequest.Valid {
			category.CountRequest = int(countRequest.Int64)
		}
		if icoFileName.Valid {
			category.IcoFileName = icoFileName.String
		}
		categores = append(categores, &category)
	}
	if err != nil {
		panic(err)
	}

	return categores
}
func (cat *Categores) GetListSecondCategores(idParent int) []*m.ViewSecondCategory {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	categores := []*m.ViewSecondCategory{}
	rows, err := db.Query(`select id, name, count_request, id_parent 
						   from catalog."tSecondLavelCategory"
						   where id_parent = $1	 
						   order by id`, idParent)

	if err != nil {
		panic(err)
	}

	var countRequest sql.NullInt64
	for rows.Next() {
		category := m.ViewSecondCategory{}
		err = rows.Scan(&category.Id, &category.Name, &countRequest, &category.IdParent)
		if countRequest.Valid {
			category.CountRequest = int(countRequest.Int64)
		}
		categores = append(categores, &category)
	}
	if err != nil {
		panic(err)
	}

	return categores
}

func (cat *Categores) UpdateFirstCategory(updateModel m.UpdateFirstCategoryModel) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	_, err := db.Exec(`update catalog."tFirstLavelCategory" SET name=$1 where id=$2`,
		updateModel.Name, updateModel.Id)
	if err != nil {
		panic("Update query error category.tCategory. Exec. \n " + err.Error())
	}
}

func (cat *Categores) UpdateIcoNameInCategory(idCategory int, icoFileName string) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	_, err := db.Exec(`update catalog."tFirstLavelCategory" SET name_file_ico=$1 where id=$2`,
		icoFileName, idCategory)
	if err != nil {
		panic("Update query error category.tCategory. Exec. \n " + err.Error())
	}
}

func (cat *Categores) GetNameIcoById(idCategory int) string {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	var nameIco sql.NullString
	var err = db.QueryRow(`select name_file_ico from catalog."tFirstLavelCategory" where id = $1`, idCategory).Scan(&nameIco)

	if err != nil {
		return ""
	}

	if nameIco.Valid {
		return nameIco.String
	}

	return ""
}
