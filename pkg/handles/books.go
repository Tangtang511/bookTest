package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}

var db *gorm.DB

func BookRouter(router *gin.Engine) {
	var db *gorm.DB
	var err error
	if db, err = gorm.Open(mysql.Open("tosone:8541539655@tcp(china.tosone.cn:3306)/database?parseTime=true"), &gorm.Config{}); err != nil {
		panic(err)
	}
	var b = &Book{}
	db.Debug().AutoMigrate(b)

	router.POST("/books", func(ctx *gin.Context) {
		var data []byte
		var err error
		if data, err = ioutil.ReadAll(ctx.Request.Body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		var book Book
		if err = json.Unmarshal(data, &book); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}

		if err = db.Debug().Create(&book).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}
	})

	router.GET("/showBooks", func(ctx *gin.Context) {
		var books []Book
		db.Table("books").Find(&books)
		ctx.JSON(http.StatusOK, gin.H{
			"message": books,
		})
	})

	router.POST("/selectBook", func(ctx *gin.Context) {
		var data []byte
		var err error
		if data, err = ioutil.ReadAll(ctx.Request.Body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
			return
		}
		type BookName struct {
			Name string `json:"name"`
		}
		var bookName BookName
		if err = json.Unmarshal(data, &bookName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unmarshal_error",
			})
			return
		}
		var selectBook []Book
		if err = db.Debug().Table("books").Where("name = ?", &bookName.Name).First(&selectBook).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "404",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": selectBook,
		})
	})
}
