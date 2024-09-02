package service

import (
	"blog-backend/global"
	"blog-backend/handler/user/dao"
	"blog-backend/model/mymodel"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestRegister struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	params := &RequestRegister{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := global.MySQL.Transaction(func(tx *gorm.DB) error {
		user := &mymodel.User{Name: params.Name}
		err := dao.CreateUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
		// 验证密码是否合法 不过默认密码肯定合法就是了
		password, err := HashPassword(params.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
		err = dao.CreateUserPassword(user.ID, params.Name, password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
	return
}
