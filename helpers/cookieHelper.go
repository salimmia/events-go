package helpers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(context *gin.Context, name string, value string, expiration time.Time) {
   cookie := buildCookie(name, value, expiration.Second())
   http.SetCookie(context.Writer, cookie)
}

func ClearCookie(context *gin.Context, name string) {
   cookie := buildCookie(name, "", -1)

   http.SetCookie(context.Writer, cookie)
}

func buildCookie(name string, value string, expires int) *http.Cookie {
   cookie := &http.Cookie{
       Name:     name,
       Value:    value,
       Path:     "/",
       HttpOnly: true,
       MaxAge:   expires,
   }

   return cookie
}