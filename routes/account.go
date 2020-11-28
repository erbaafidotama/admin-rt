package routes

import (
	"admin-rt/config"
	"admin-rt/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAccount Route
func GetAccount(c *gin.Context) {
	db := config.GetDB()
	accounts := []models.Account{}

	// select * from User
	if err := db.Find(&accounts).Error; err != nil {
		// return error
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// return complete
	c.JSON(200, gin.H{
		"message": "GET data Account",
		"data":    accounts,
	})
}

// PostAccount this
func PostAccount(c *gin.Context) {
	db := config.GetDB()
	var roleAdmin bool
	// convert string date to date db
	dateStr := c.PostForm("date_birth")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	// check admin role
	if c.PostForm("admin_role") == "true" {
		roleAdmin = true
	}

	fmt.Println("create account")
	// make object from form body
	account := models.Account{
		Username:  c.PostForm("username"),
		FullName:  c.PostForm("full_name"),
		Email:     c.PostForm("email"),
		Password:  c.PostForm("password"),
		DateBirth: date,
		AdminRole: roleAdmin,
	}

	// crete data to db
	// config.DB.Create(&account)
	db.Create(&account)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   account,
	})
}

func UpdateAccount(c *gin.Context) {
	db := config.GetDB()
	var roleAdmin bool

	// get id from url
	accountID := c.Param("id")

	var dataAccount models.Account
	if err := db.Where("id = ?", accountID).First(&dataAccount).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// convert string date to date db
	dateStr := c.PostForm("date_birth")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	if c.PostForm("admin_role") == "true" {
		roleAdmin = true
	}

	db.Model(&dataAccount).Where("id = ?", accountID).Updates(models.Account{
		FullName:  c.PostForm("full_name"),
		DateBirth: date,
		AdminRole: roleAdmin,
	})

	c.JSON(200, gin.H{
		"status": "Success",
		"data":   dataAccount,
	})
}

func DeleteAccount(c *gin.Context) {
	db := config.GetDB()
	// get id from url
	accountID := c.Param("id")

	var dataAccount models.Account
	if err := db.Where("id = ?", accountID).First(&dataAccount).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Where("id = ?", accountID).Delete(&dataAccount)

	c.JSON(200, gin.H{
		"status": "Success Delete",
		"data":   dataAccount,
	})
}
