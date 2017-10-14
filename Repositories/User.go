package Repositories

import (
	m "Users/Models"
	connect "Users/Repositories/Connector"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
}

func (u *User) UpdateUserCookie(cookie string, idUser int) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	const updateQuery = `update profile."tUser" SET cookie_hash=$1 WHERE id = $2`
	update, err := db.Prepare(updateQuery)
	defer update.Close()

	if err != nil {
		panic("Update prepare query error")
	}
	_, err = update.Exec(cookie, idUser)

	if err != nil {
		panic("Update error query error")
	}
}

func (u *User) GetUserModelByCookie(cookieHash string) m.User {
	var user = m.User{}
	if cookieHash == "" {
		return user
	}

	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	rows, errQuery := db.Query(`select id, name, email, date_registration, date_activation, is_locked, 
	is_activate, count_try_auth, phone, description from profile."tUser" where cookie_hash = $1`, cookieHash)

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.DateRegistration, &user.DateActivation,
			&user.IsLocked, &user.IsActivate, &user.CountTryAuth, &user.Phone, &user.Description)
	}
	errQuery = rows.Err()

	if errQuery != nil {
		user.ErrValue = errQuery.Error()
	}

	return user
}
func (u *User) GetUserById(idUser int) m.User {
	var user = m.User{}
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	rows, errQuery := db.Query(`select id, name, email, date_registration, date_activation, is_locked, 
	is_activate, count_try_auth, phone, description from profile."tUser" where id = $1`, idUser)

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.DateRegistration, &user.DateActivation,
			&user.IsLocked, &user.IsActivate, &user.CountTryAuth, &user.Phone, &user.Description)
	}
	errQuery = rows.Err()

	if errQuery != nil {
		user.ErrValue = errQuery.Error()
	}

	return user
}

func (u *User) GetIdUserByName(name string) int {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()
	var id sql.NullInt64
	var err = db.QueryRow(`select id from profile."tUser" where name = $1`, name).Scan(&id)
	if err != nil {
		return 0
	}

	if id.Valid {
		return int(id.Int64)
	}

	return 0
}

func (u *User) SelectSliceUsers(searchNameUser string) []*m.User {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	usersList := []*m.User{}
	rows, err := db.Query(`select id, name, email, description,
	date_activation, date_registration, is_locked, is_activate, phone 
	from profile."tUser" where name like $1`, "%"+searchNameUser+"%")

	for rows.Next() {
		user := m.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Description,
			&user.DateActivation, &user.DateRegistration,
			&user.IsLocked, &user.IsActivate, &user.Phone)
		usersList = append(usersList, &user)
	}
	if err != nil {
		panic(err)
	}

	return usersList
}

func (u *User) IsEmailInBase(email string) bool {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	var isYes sql.NullBool
	var err = db.QueryRow(`select TRUE from profile."tUser" where email = $1`, email).Scan(&isYes)

	fmt.Printf("%s -- \n", err)
	if err != nil {
		return false
	}

	if isYes.Valid {
		fmt.Printf("%t -- \n", bool(isYes.Bool))
		return bool(isYes.Bool)
	}

	return true
}

func (u *User) IsLock(idUser int) bool {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	var isLock sql.NullBool
	var err = db.QueryRow(`select is_locked from profile."tUser" where id = $1`, idUser).Scan(&isLock)

	if err != nil {
		return false
	}

	if isLock.Valid {
		return bool(isLock.Bool)
	}

	return true
}

func (u *User) SetLockStatus(idUser int) (isUp bool) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	_, err := db.Exec(`update profile."tUser" set is_locked = $1 where id = $2`, true, idUser)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func (u *User) SetUnLockStatus(idUser int) (isUp bool) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	_, err := db.Exec(`update profile."tUser" set is_locked = $1 where id = $2`, false, idUser)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func (u *User) GetNameUserById(idUser int) string {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()
	var name sql.NullString
	var err = db.QueryRow(`select name from profile."tUser" where id = $1`, idUser).Scan(&name)
	if err != nil {
		return ""
	}

	if name.Valid {
		return string(name.String)
	}

	return ""
}

func (u *User) InsertUserData(user m.User) (idUser int) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	const INSERT_QUERY = `insert into profile."tUser"
	(name, email, date_registration, date_activation, is_locked, is_activate, 
	 count_try_auth, phone, description, password_hash)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;`

	var id = 0
	insertQuery, err := db.Prepare(INSERT_QUERY)
	if err != nil {
		panic("Insert query error profile.tUser")
		return 0
	}

	err = insertQuery.QueryRow(user.Name,
		user.Email,
		user.DateRegistration,
		user.DateActivation,
		user.IsLocked,
		user.IsActivate,
		user.CountTryAuth,
		user.Phone,
		user.Description,
		user.PassworsHash).Scan(&id)

	if err != nil {
		panic("Insert query error profile.tUser")
		return 0
	}

	return id
}

func (u *User) UpdateUserData(user m.User) bool {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	const UPDATE_QUERY = `update profile."tUser" set name = $2, email = $3,
	password_hash = $4 where id = $1`

	_, err := db.Exec(UPDATE_QUERY, user.Id, user.Name, user.Email, user.PassworsHash)
	if err != nil {
		panic("Update query error profile.tUser")
		return false
	}

	return true
}
