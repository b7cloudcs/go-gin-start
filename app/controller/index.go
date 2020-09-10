package controller

import (
	"jinrongdati/app/util"
	"time"

	"github.com/gin-gonic/gin"
)

type Index struct{}

/**
 * Index
**/
func (Index) Index(c *gin.Context) {

	// return
	c.JSON(200, gin.H{
		"create_at":   util.FirstIni.Section("").Key("created_at").String(),
		"server_time": time.Now(),
	})
}
