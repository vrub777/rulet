package Services

import (
	r "Users/Repositories"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
)

type Authorithator struct {
}

func (this *Authorithator) IsValidPasswordForUser(name string, password string) bool {
	auth := r.IAuthUserer(&r.AuthUser{})
	var hashPassword = auth.GetUserPasswordHash(name)
	var nowHash = this.GetHashPassword(password)

	if hashPassword == nowHash {
		return true
	}

	return false
}

func (this *Authorithator) GetCountTryAuth(nameUser string) (countTry int) {
	auth := r.IAuthUserer(&r.AuthUser{})
	return auth.GetCountTry(nameUser)
}

func (this *Authorithator) IsNameUserInSystem(name string) bool {
	auth := r.IAuthUserer(&r.AuthUser{})
	return auth.IsNameUserInSystem(name)
}

func (a *Authorithator) IsAuth(r *http.Request) bool {
	cookieService := Cookierator{}
	authCookie := cookieService.GetCookie(r)
	return a.IsAuthCookieInSystem(authCookie)
}

func (a *Authorithator) IsAuthCookieInSystem(value string) bool {
	if value == "" {
		return false
	}

	authRepository := r.IAuthUserer(&r.AuthUser{})
	return authRepository.IsAuthCookieInSystem(value)
}

func (this *Authorithator) Authorithation(responseWriter gin.ResponseWriter, name string, password string) {
	user := r.IUserer(&r.User{})
	var idUser = user.GetIdUserByName(name)
	if idUser <= 0 {
		return
	}

	auth(responseWriter, idUser, name)
}

func (this *Authorithator) AuthorithationByIdUser(responseWriter gin.ResponseWriter, idUser int) {
	user := r.IUserer(&r.User{})
	var nameUser = user.GetNameUserById(idUser)
	if nameUser == "" {
		return
	}
	auth(responseWriter, idUser, nameUser)
}

func (this *Authorithator) Out(responseWriter gin.ResponseWriter) {
	cookie := Cookierator{}
	cookie.DeleteCookie(responseWriter)
}

func (this *Authorithator) GetHashPassword(password string) string {
	md5Hash := md5.New()
	io.WriteString(md5Hash, "777"+password)
	var md5byte = md5Hash.Sum(nil)
	var md5string = base32.HexEncoding.EncodeToString(md5byte[:])

	sha1Hash := sha1.New()
	io.WriteString(sha1Hash, md5string)
	var sha1byte = sha1Hash.Sum(nil)
	var sha1string = base32.HexEncoding.EncodeToString(sha1byte[:])

	return sha1string
}

func auth(responseWriter gin.ResponseWriter, idUser int, nameUser string) {
	auth := r.IAuthUserer(&r.AuthUser{})
	auth.InsertUserInJournalAuthorization(idUser)

	var hashString = getRandHash(nameUser)
	cookie := Cookierator{}
	cookie.SetCookie(responseWriter, hashString)

	rUser := r.IUserer(&r.User{})
	rUser.UpdateUserCookie(hashString, idUser)
}

func getRandHash(nameUser string) string {
	var n = 32
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	var randStr = string(b)
	md5Hash := md5.New()
	io.WriteString(md5Hash, "727"+nameUser+randStr)
	var md5byte = md5Hash.Sum(nil)
	var md5string = base32.HexEncoding.EncodeToString(md5byte[:])

	return md5string
}
