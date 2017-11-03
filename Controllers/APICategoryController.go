package Controllers

import (
	//m "Users/Models"
	s "Users/Services"
	//"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type APICategoryController struct {
}

func (ac *APICategoryController) GetListCategores(c *gin.Context) {
	categoryService := s.Categoryzator{}
	listCategores := categoryService.GetListFirstCategores()

	var listJSONCategores []gin.H

	for _, value := range listCategores {
		listJSONCategores = append(listJSONCategores,
			gin.H{"Id": value.Id, "Name": value.Name, "Order": value.Order,
				"CountRequest": value.CountRequest, "CountSecondLevel": value.CountSecondLevel,
				"IcoFullPath": value.IcoFullPath})
	}
	c.JSON(200, listJSONCategores)
}

func (ac *APICategoryController) GetListSecondCategoresById(c *gin.Context) {
	idParent, err := strconv.Atoi(c.Params.ByName("idparent"))

	if err != nil {
		return
	}

	categoryService := s.Categoryzator{}
	listCategores := categoryService.GetListSecondCategores(idParent)

	var listJSONCategores []gin.H

	for _, value := range listCategores {
		listJSONCategores = append(listJSONCategores,
			gin.H{"Id": value.Id, "Name": value.Name, "Order": value.Order,
				"CountRequest": value.CountRequest, "IdParent": value.IdParent})
	}
	c.JSON(200, listJSONCategores)
}
