package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ngoduongkha/go-2023/assignment-06/datatransfers"
	"github.com/ngoduongkha/go-2023/assignment-06/handlers"
	"github.com/ngoduongkha/go-2023/assignment-06/models"
)

func GETProducts(c *gin.Context) {
	var err error
	var products []models.Product
	if products, err = handlers.Handler.GetProducts(); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: products})
}

func GETProduct(c *gin.Context) {
	var err error
	var product models.Product
	var id64 uint64
	if id64, err = strconv.ParseUint(c.Param("id"), 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	id := uint(id64)
	if product, err = handlers.Handler.GetProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: product})
}

func PUTProduct(c *gin.Context) {
	var err error
	var product datatransfers.ProductUpdate
	var id64 uint64
	if id64, err = strconv.ParseUint(c.Param("id"), 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	id := uint(id64)
	if err = c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.UpdateProduct(id, product); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: "product updated"})
}

func POSTProduct(c *gin.Context) {
	var err error
	var product datatransfers.ProductCreate
	if err = c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	if err = handlers.Handler.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, datatransfers.Response{Data: "product created"})
}

func DELETEProduct(c *gin.Context) {
	var err error
	var id64 uint64
	if id64, err = strconv.ParseUint(c.Param("id"), 10, 32); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
		return
	}
	id := uint(id64)
	if err = handlers.Handler.DeleteProductByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, datatransfers.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datatransfers.Response{Data: "product deleted"})
}
