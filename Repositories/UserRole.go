package Repositories

import (
	m "Users/Models"
	connect "Users/Repositories/Connector"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)

type UserRole struct {
}

func (u *UserRole) GetListRoles() *[]m.UserRole {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	listRoles := []m.UserRole{}
	rows, err := db.Query(`select id, name, description from profile."tRole"`)

	for rows.Next() {
		role := m.UserRole{}
		err = rows.Scan(&role.Id, &role.Name, &role.Description)
		listRoles = append(listRoles, role)
	}
	if err != nil {
		panic(err)
	}
	return &listRoles
}
func (u *UserRole) GetListUserRoles(idUser int) *[]m.UserRole {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	listRoles := []m.UserRole{}
	rows, err := db.Query(`select id, name, description 
							from profile."tRole" r
								inner join profile."tUserInRole" ur on(ur.id_role=r.id)
							where ur.id_user = $1`, idUser)
	for rows.Next() {
		role := m.UserRole{}
		err = rows.Scan(&role.Id, &role.Name, &role.Description)
		listRoles = append(listRoles, role)
	}
	if err != nil {
		panic(err)
	}
	return &listRoles
}

func (u *UserRole) IsUserInRole(idUser int, idRole int) bool {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	var isUserInRole sql.NullBool
	var err = db.QueryRow("select TRUE from tUserInRole where id_user = $1 and id_role = $2", idUser, idRole).Scan(&isUserInRole)

	if err != nil {
		return false
	}

	if isUserInRole.Valid {
		return true
	}

	return false
}

func (u *UserRole) AddListRoles(idUser int, listIdRoles []int) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	strIdRoles := ""
	strIdUser := strconv.Itoa(idUser)
	postfix := ", "
	for i, value := range listIdRoles {
		if i == len(listIdRoles)-1 {
			postfix = ""
		}
		strIdRoles = strIdRoles + "(" + strIdUser + "," + strconv.Itoa(value) + ")" + postfix
	}

	_, err := db.Exec(`insert into profile."tUserInRole" (id_user, id_role) values ` + strIdRoles)
	if err != nil {
		panic("Insert query error profile.tUserRole. Prepare. \n " + err.Error())
	}
}

func (u *UserRole) UpdateListRoles(idUser int, listIdRoles []int) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	strIdRoles := ""
	strIdUser := strconv.Itoa(idUser)
	postfix := ", "
	for i, value := range listIdRoles {
		if i == len(listIdRoles)-1 {
			postfix = ""
		}
		strIdRoles = strIdRoles + "(" + strIdUser + "," + strconv.Itoa(value) + ")" + postfix
	}

	_, err := db.Exec(`delete from profile."tUserInRole" where id_user = ` + strIdUser + `; 
				insert into profile."tUserInRole" (id_user, id_role) values ` + strIdRoles)
	if err != nil {
		panic("Insert query error profile.tUserRole. Prepare. \n " + err.Error())
	}
}
