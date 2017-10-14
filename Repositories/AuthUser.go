package Repositories

import (
	connect "Users/Repositories/Connector"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type AuthUser struct {
}

func (this *AuthUser) InsertUserInJournalAuthorization(idUser int) {
	postgre := connect.Postgre{}
	db := postgre.Open()

	const INSERT_QUERY = `insert into profile."tJournalUserAuthorization"
	(date_authorization, id_user)
	values ($1, $2);`

	insert_query, err := db.Prepare(INSERT_QUERY)
	defer insert_query.Close()

	_, err = insert_query.Exec(time.Now(), idUser)

	if err != nil {
		panic("Error insert into tJournalUserAuthorization")
	}
}

func (this *AuthUser) IsNameUserInSystem(name string) bool {
	postgre := connect.Postgre{}
	db := postgre.Open()

	var isYes sql.NullBool
	var err = db.QueryRow(`select TRUE from profile."tUser" where name = $1`, name).Scan(&isYes)

	if err != nil {
		return false
	}

	if isYes.Valid {
		return true
	}

	return false
}

func (this *AuthUser) GetUserPasswordHash(name string) string {
	postgre := connect.Postgre{}
	db := postgre.Open()

	var passwordHash string
	var err = db.QueryRow(`select password_hash from profile."tUser" where name = $1`, name).Scan(&passwordHash)

	if err != nil {
		return ""
	}

	return passwordHash
}

func (this *AuthUser) GetCountTry(nameUser string) (countTry int) {
	postgre := connect.Postgre{}
	db := postgre.Open()
	defer db.Close()

	rows, err :=
		db.Query(`select count_try_auth from profile."tUser" where name = $1`, nameUser)

	for rows.Next() {
		var countTry int
		err = rows.Scan(&countTry)
		if err != nil {
			return 0
			panic(err.Error())
		}
		return countTry
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return 0
}

func (this *AuthUser) IsAuthCookieInSystem(value string) bool {
	postgre := connect.Postgre{}
	db := postgre.Open()
	var isYes sql.NullBool
	var err = db.QueryRow(`select TRUE from profile."tUser" where cookie_hash = $1`, value).Scan(&isYes)

	if err != nil {
		return false
	}

	if isYes.Valid {
		return true
	}

	return false
}
