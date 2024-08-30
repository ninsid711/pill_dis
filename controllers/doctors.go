package controllers

import (
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/ninsid711/pill_dis/models"
	"github.com/nrednav/cuid2"
)

func CreateDoctor(ctx *gin.Context) {
	type newDocInput struct {
		name     string `json:"name"`
		age      uint   `json:"age"`
		gender   string `json:"gender"`
		password string `json:"password"`
	}

	var newDocRequest newDocInput
	var newDoc models.Doctor

	if err := ctx.ShouldBindJSON(&newDocRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newDocRequest.name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if newDocRequest.password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	if newDocRequest.gender == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "gender is required"})
		return
	}

	if newDocRequest.age == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "age is required"})
		return
	}

	err := db.DB.Where("email = ? OR username = ?", newDocRequest.name, newDocRequest.age, newDocRequest.password, newDocRequest.gender).First(&newDoc).Error
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
	}

	id := cuid2.Generate()
	if id == "" {
		return
	}

	newDoc.DoctorId = id

	hash, err := argon2id.CreateHash(newDocRequest.password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal("could not hash password", err)
	}

	newDoc = models.Doctor{
		DoctorId: id,
		Name:     newDocRequest.name,
		Age:      newDocRequest.age,
		Gender:   newDocRequest.gender,
		Password: hash,
	}

	res := db.DB.Create(&newDoc)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func FetchDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	var doc models.Doctor
	res := db.DB.Raw("SELECT * FROM doctor WHERE id = ?", id).Scan(&doc)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"doctor": doc})
}
