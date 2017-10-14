package Services

type IAuthorizater interface {
	IsAuthCookieInSystem(value string) bool // Есть ли такой куки-хэш у пользователя
}
