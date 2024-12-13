package jwt

import (
	"github.com/DSXRIIIII/go-utils/go-jwt/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = 200
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			logrus.Info("token is nil")
			ecode = 404
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				logrus.Info("token parse error")
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = 401
				default:
					ecode = 402
				}
			}
		}
		if ecode != 200 {
			c.JSON(ecode, "token false")
			c.Abort()
		}
		c.Next()
	}
}
