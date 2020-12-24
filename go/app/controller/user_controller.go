package controller

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"../model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	Db *gorm.DB
}

var err error

func SignUpError(user model.User, password string) []string {

	var errormessage []string

	if user.Username == "" {
		errormessage = append(errormessage, "ユーザー名を入力してください\n")
	}
	if password == "" {
		errormessage = append(errormessage, "パスワードを入力してください\n")
	} else if len(password) < 4 || len(password) > 8 {
		errormessage = append(errormessage, "パスワードは4〜8文字で入力してください\n")
	}

	return errormessage
}

func (handler *UserHandler) SignUp(c *gin.Context) {
	user := model.User{}
	username, _ := c.GetPostForm("username")
	username = strings.TrimSpace(username)

	password, _ := c.GetPostForm("password")
	password = strings.TrimSpace(password)

	user.Username = username

	// エラー判定
	errormessage := SignUpError(user, password)
	if errormessage != nil {
		c.JSON(200, gin.H{
			"errormessage": errormessage,
			"user":         user,
		})
		return
	}

	// パスワードをハッシュ値に変更
	setpassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	user.Password = setpassword

	// ユーザートークンを作成する
	seedtoken := username + password
	usertoken := fmt.Sprintf("%x", seedtoken)
	user.Usertoken = usertoken

	if errormessage == nil {
		// 入力情報が登録されているか確認する
		if err := handler.Db.Create(&user).Error; err != nil {
			errormessage = append(errormessage, "入力された情報はすでに登録されています")
			c.JSON(200, gin.H{
				"errormessage": errormessage,
				"user":         user,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"errormessage": errormessage,
		"user":         user,
	})
}

func (handler *UserHandler) SignIn(c *gin.Context) {
	user := model.User{}
	var message string

	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	searchtoken := fmt.Sprintf("%x", (username + password))

	handler.Db.Where("Usertoken = ?", searchtoken).Find(&user)

	if user.Usertoken == searchtoken {
		message = "ログインしました"
	}

	c.JSON(200, gin.H{
		"message": message,
		"user":    user,
	})
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	var users []model.User

	userid, _ := c.GetPostForm("id")
	fmt.Println(userid)

	handler.Db.Where("ID <> ?", userid).Find(&users)

	c.JSON(200, users)
}

func (handler *UserHandler) UserChatRoom(c *gin.Context) {

	userchatroom := model.Userchatroom{}
	curentuserid, _ := c.GetPostForm("curentuserid")
	seconduserid, _ := c.GetPostForm("seconduserid")

	intcurentuserid, _ := strconv.Atoi(curentuserid)
	setcurentuserid := uint(intcurentuserid)

	intseconduserid, _ := strconv.Atoi(seconduserid)
	setseconduserid := uint(intseconduserid)

	if err := handler.Db.Where("Firstuserid in (?) AND Seconduserid in (?)", []uint{setcurentuserid, setseconduserid}, []uint{setcurentuserid, setseconduserid}).Find(&userchatroom).Error; err != nil {
		fmt.Println("err")
		userchatroom.Firstuserid = setcurentuserid
		userchatroom.Seconduserid = setseconduserid
		handler.Db.Create(&userchatroom)
	}

	c.JSON(200, userchatroom)
}
