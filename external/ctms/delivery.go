package ctms

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
}
