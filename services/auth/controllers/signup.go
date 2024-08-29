package controllers

import (
	"net/http"
	"pill_dis/models"

	"github.com/gin-gonic/gin"
)

func createPatient(ctx *gin.Context) {
	type newPatientInput struct {
		Name           string `json:"username"`
		Age            uint   `json:"age"`
		Gender         string `json:"gender"`
		Illness_Injury string `json:"illness/injury"`
		Password       string `json:"password"`
	}

	var newPatientReq newPatientInput
	var newPatient models.Patient

	if err := ctx.ShouldBindJSON(&newPatientReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newPatientReq.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Patient name required"})
		return
	}
	if newPatientReq.Age == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Patient age required"})
		return
	}
	if newPatientReq.Gender == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Patient gender required"})
		return
	}
	if newPatientReq.Illness_Injury == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Patient Illness or Injury required"})
		return
	}
	if newPatientReq.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password required"})
		return
	}

	err := db.DB
}
