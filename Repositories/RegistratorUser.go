package Repositories

import (
	m "Users/Models"
	connect "Users/Repositories/Connector"
	_ "github.com/lib/pq"
)

type RegistratorUser struct {
}

func (this *RegistratorUser) InsertUserData(user m.User) (idUser int) {
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
