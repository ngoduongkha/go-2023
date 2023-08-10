package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoduongkha/go-2023/assignment-06/constants"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
)

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: "user not authenticated"})
	}
}
