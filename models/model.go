package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	PatientID         string `json:"patient_id" gorm:"primaryKey"`
	Name              string `json:"name" gorm:"unique;not null"`
	Age               uint   `json:"age" gorm:"not null"`
	Gender            string `json:"gender" gorm:"not null"`
	Illness_or_Injury string `json:"illness/injury" gorm:"not null"`
	Password          string `json:"-" gorm:"not null"`
}

type Doctor struct {
	gorm.Model
	DocID        string    `json:"doc_id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Age          uint      `json:"age" gorm:"not null"`
	Gender       string    `json:"gender" gorm:"not null"`
	Password     string    `json:"-" gorm:"not null"`
	PatientsList []Patient `json:"patients" gorm:"foreignKey:DoctorID"`
}

type Prescription struct {
	gorm.Model
	Pres_ID string `json:"pres_id" gorm:"primaryKey"`
}
