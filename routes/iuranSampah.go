package routes

import (
	"admin-rt/config"
	"admin-rt/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var db = config.GetDB()

// PostIuranSampah Route
func PostIuranSampah(c *gin.Context) {
	dateStr := c.PostForm("pay_date")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	n, _ := strconv.ParseUint(c.PostForm("account_id"), 10, 64)

	iuranSampah := models.IuranSampah{
		AccountID:   uint(n),
		PayDate:     date,
		Description: c.PostForm("description"),
	}

	db.Create(&iuranSampah)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   iuranSampah,
	})
}
