package routes

import (
	"admin-rt/config"
	"admin-rt/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PostCitizen Route
func PostCitizen(c *gin.Context) {
	db := config.GetDB()
	dateStr := c.PostForm("date_birth")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	n, _ := strconv.ParseUint(c.PostForm("account_id"), 10, 64)

	citizen := models.Citizen{
		AccountID:  uint(n),
		FullName:   c.PostForm("full_name"),
		DateBirth:  date,
		NoKk:       c.PostForm("no_kk"),
		NoKtp:      c.PostForm("no_ktp"),
		PlaceBirth: c.PostForm("place_birth"),
		Address:    c.PostForm("address"),
	}

	db.Create(&citizen)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   citizen,
	})
}
