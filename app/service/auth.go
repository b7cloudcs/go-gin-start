package service

import (
	"go-gin-start/ent"

	"github.com/gin-gonic/gin"
)

/**
 * 获取登陆的用户
**/
func GetAuthUser(c *gin.Context) *ent.User {

	user, ok := c.Get("user")

	if !ok {
		return nil
	}

	u := user.(ent.User)

	return &u
}

/**
 * 获取登陆的用户ID
**/
func GetAuthUserID(c *gin.Context) int {

	user, ok := c.Get("user")

	if !ok {
		return 0
	}

	u := user.(ent.User)

	return u.ID
}
