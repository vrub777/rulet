package Services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Cookierator struct {
}

const AUTHCOOKIE = "useram"

func (this *Cookierator) SetCookie(responseWriter gin.ResponseWriter, value string) {
	expiration := time.Now().Add(375 * 24 * time.Hour)
	cookie := &http.Cookie{Name: AUTHCOOKIE, Path: "/", Value: value, Expires: expiration}
	http.SetCookie(responseWriter, cookie)
}

func (this *Cookierator) GetCookie(r *http.Request) string {
	var cookie, err = r.Cookie(AUTHCOOKIE)
	if err != nil {
		return ""
	}

	return cookie.Value
}

func (this *Cookierator) DeleteCookie(responseWriter gin.ResponseWriter) {
	expiration := time.Now().Add(-375 * 24 * time.Hour)
	cookie := &http.Cookie{Name: AUTHCOOKIE, Path: "/", Value: "", Expires: expiration}
	http.SetCookie(responseWriter, cookie)
}
