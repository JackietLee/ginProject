package admin

import (
	"fmt"
	"ginProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) UserIndex(c *gin.Context) {
	user := []model.User{}
	model.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	//con.success(c)
}

func (con UserController) UserAdd(c *gin.Context) {
	/*	value, _ := c.Get("username")
		fmt.Println(value)
		s, ok := value.(string)
		if ok {
			c.String(200, "用户列表-add"+s)

		} else {
			c.String(200, "用户列表-add获取用户失败")
		}*/
	user := model.User{
		Name:  "lijian",
		Age:   30,
		State: 2,
	}
	model.DB.Create(&user)
	fmt.Println(user)
	c.String(http.StatusOK, "成功")
}

func (con UserController) UserUpdate(c *gin.Context) {
	/*	user := model.User{Id: 1}
		model.DB.Find(&user)
		user.Age = 100
		model.DB.Save(&user)
		c.String(http.StatusOK, "成功")*/

	user := model.User{}
	model.DB.Model(&user).Where("id = ?", 2).Update("name", "帅哥")
	c.String(http.StatusOK, "成功")
}

func (con UserController) UserDel(c *gin.Context) {
	/*	user := model.User{Id: 1}
		model.DB.Delete(&user)*/
	user := model.User{}
	model.DB.Where("name=?", "帅哥").Delete(&user)
	c.String(http.StatusOK, "成功")
}
