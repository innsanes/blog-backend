package service

import (
	"blog-backend/global"
	"blog-backend/handler/user/dao"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strconv"
)

type RequestLogin struct {
	Name     string `json:"name" binding:"required,max=10"`
	Password string `json:"password" binding:"required,max=20"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	params := &RequestLogin{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	up, err := dao.QueryUserPasswordByName(params.Name)
	if err != nil {
		global.Log.Error("QueryUserPasswordByName error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if up.UserID == 0 {
		global.Log.Error("NameNotExist error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名不存在"})
		return
	}
	if ComparePassword(params.Password, up.Password) != nil {
		global.Log.Error("PasswordIncorrect error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "密码错误"})
		return
	}

	token, err := MakeSession()
	if err != nil {
		global.Log.Error("MakeSession error: ", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	global.Token.AddToken(token)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
	return
}

func HashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(fromPassword), err
}

func ComparePassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeSession() (string, error) {
	return "t" + strconv.FormatInt(rand.Int63(), 10), nil
}
