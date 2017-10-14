package Repositories

type IAuthUserer interface {
	InsertUserInJournalAuthorization(idUser int)
	IsNameUserInSystem(name string) bool
	GetUserPasswordHash(name string) string
	GetCountTry(nameUser string) (countTry int)
	IsAuthCookieInSystem(value string) bool
}
