package control

import (
	"fmt"

	_ "github.com/et-zone/embi/dao"
	"github.com/et-zone/embi/model"
	"github.com/gin-gonic/gin"
)

func TPost(c *gin.Context) {
	t := &model.ETarget{}

	err := c.Bind(t)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"message": "ok",
		})
		return
	}
	// err = dao.InsertEtarget(t)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	c.JSON(200, gin.H{
		"message": "ok",
	})
	return

}

func TGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
