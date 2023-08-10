package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngoduongkha/go-2023/assignment-06/config"
	"github.com/ngoduongkha/go-2023/assignment-06/constants"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
)

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}
	claims, err := parseToken(token, config.AppConfig.JWTSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: err.Error()})
		return
	}
	c.Set(constants.IsAuthenticatedKey, true)
	c.Set(constants.UserIDKey, claims.ID)
	c.Next()
}

func parseToken(tokenString, secret string) (claims jwt.RegisteredClaims, err error) {
	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil || !token.Valid {
		return jwt.RegisteredClaims{}, fmt.Errorf("invalid token. %s", err)
	}
	return
}
