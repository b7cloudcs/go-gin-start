package middleware

import (
	"github.com/gin-gonic/gin"
)

// AuthenticationRequired ...
func AuthenticationRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		// // header params
		// token := c.Request.Header.Get("authorization")

		// // load user
		// entUser, err := util.DBC.User.
		// 	Query().
		// 	Where(user.TokenEQ(token)).
		// 	First(util.Ctx)

		// // err
		// if err != nil {
		// 	c.JSON(
		// 		http.StatusUnauthorized,
		// 		res.ErrorMsg{
		// 			Code: 401,
		// 			Msg:  "Unauthorized",
		// 		},
		// 	)
		// 	c.Abort()
		// 	return
		// }

		// // auth user && next
		// c.Set("user", *entUser)
		c.Next()

	}
}
