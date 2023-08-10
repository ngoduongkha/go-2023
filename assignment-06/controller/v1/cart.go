package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngoduongkha/go-2023/assignment-06/constants"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"github.com/ngoduongkha/go-2023/assignment-06/handlers"
	"github.com/ngoduongkha/go-2023/assignment-06/models"
	"net/http"
	"strconv"
)

func GETCart(c *gin.Context) {
	var err error
	var cart models.Cart
	idStr, isExists := c.Get(constants.UserIDKey)
	if isExists {
		fmt.Println(idStr)
	}
	var id64, _ = strconv.Atoi(idStr.(string))
	if cart, err = handlers.Handler.GetCart(uint(id64)); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: cart})
}

func POSTCartItem(c *gin.Context) {
	var err error
	var dto datatransfers.AddCartItem

	if err = c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	idStr, isExists := c.Get(constants.UserIDKey)
	if isExists {
		fmt.Println(idStr)
	}
	var id64, _ = strconv.Atoi(idStr.(string))
	if err = handlers.Handler.AddToCart(uint(id64), dto); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: "added to cart"})
}

func DELETECartItem(c *gin.Context) {
	var err error

	idStr, isExists := c.Get(constants.UserIDKey)
	if isExists {
		fmt.Println(idStr)
	}
	var id64, _ = strconv.Atoi(idStr.(string))
	cartItemIdStr := c.Param("id")
	var cartItemId64, _ = strconv.Atoi(cartItemIdStr)
	if err = handlers.Handler.RemoveFromCart(uint(id64), uint(cartItemId64)); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: "deleted from cart"})
}

func POSTCheckout(c *gin.Context) {
	var err error
	var checkout datatransfers.Checkout

	idStr, isExists := c.Get(constants.UserIDKey)
	if isExists {
		fmt.Println(idStr)
	}
	var id64, _ = strconv.Atoi(idStr.(string))
	if checkout, err = handlers.Handler.Checkout(uint(id64)); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: checkout})
}
