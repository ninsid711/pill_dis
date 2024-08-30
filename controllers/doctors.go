package controllers

import (
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/ninsid711/pill_dis/initializers"
	"github.com/ninsid711/pill_dis/models"
	"github.com/nrednav/cuid2"
)

func CreateDoctor(c *gin.Context) {
	type newReq struct {
		name     string `json:"name"`
		age      uint   `json:"age"`
		gender   string `json:"gender"`
		illness  string `json:"illness"`
		password string `json:"password"`
	}
	var newDoc models.Doctor
	var newDocReq newReq

	if err := c.ShouldBindJSON(&newDocReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newDocReq.name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if newDocReq.age == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "age is required"})
		return
	}

	if newDocReq.gender == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender is required"})
		return
	}

	if newDocReq.illness == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "illness is required"})
		return
	}

	if newDocReq.password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	// Generate ID using cuid
	id := cuid2.Generate()
	newDocReq.password = id

	// Hash the password using argon2id
	hashedPassword, err := argon2id.CreateHash(newDocReq.password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newDocReq.password = hashedPassword

	newDoc = models.Doctor{
		ID:       id,
		Name:     newDocReq.name,
		Age:      newDocReq.age,
		Gender:   newDocReq.gender,
		Illness:  newDocReq.illness,
		Password: newDocReq.password,
	}
	if err := initializers.DB.Create(&newDoc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor"})
		return
	}

	c.JSON(http.StatusOK, newDoc)
}

func CreatePatient(c *gin.Context) {
	doctorID := c.Param("doctorID")
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate ID using cuid
	patient.ID = cuid2.Generate()
	patient.DoctorID = doctorID

	if err := initializers.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func GetPatientsByDoctor(c *gin.Context) {
	doctorID := c.Param("doctorID")
	var patients []models.Patient

	if err := initializers.DB.Where("doctor_id = ?", doctorID).Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func AddMedicineToPatient(c *gin.Context) {
	patientID := c.Param("patientID")
	var patient models.Patient

	if err := initializers.DB.Where("id = ?", patientID).First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var medicine models.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate ID using cuid
	medicine.ID = cuid2.New().String()

	// Add medicine to the patient's list
	patient.Medicines = append(patient.Medicines, medicine)

	if err := initializers.DB.Save(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add medicine"})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func GetMedicinesByPatient(c *gin.Context) {
	patientID := c.Param("patientID")
	var patient Patient

	if err := db.Where("id = ?", patientID).Preload("Medicines").First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient.Medicines)
}

func DispensePills(c *gin.Context) {
	patientID := c.Param("patientID")
	var patient models.Patient

	if err := initializers.DB.Where("id = ?", patientID).Preload("Medicines").First(&patient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Integrate with Arduino to send the dispense signal
	err := sendDispenseSignal(patient.Medicines) // Implement this function
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dispense pills"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Pills dispensed successfully"})
}
