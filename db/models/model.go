package db

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	PatientID         string     `json:"patient_id" gorm:"primaryKey"`
	Name              string     `json:"name" gorm:"unique;not null"`
	Age               uint       `json:"age" gorm:"not null"`
	Gender            string     `json:"gender" gorm:"not null"`
	Illness_or_Injury string     `json:"illness/injury" gorm:"not null"`
	Password          string     `json:"-" gorm:"not null"`
	Prescription      []Medicine `gorm:"foreignKey:PatientID"`
}

type Doctor struct {
	gorm.Model
	DocID        string    `json:"doc_id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Age          uint      `json:"age" gorm:"not null"`
	Gender       string    `json:"gender" gorm:"not null"`
	Password     string    `json:"-" gorm:"not null"`
	PatientsList []Patient `gorm:"foreignKey:DoctorID"`
}

type Medicine struct {
	gorm.Model
	Med_ID       string `json:"pres_id" gorm:"primaryKey"`
	MedicineName string `json:"medname" gorm:"unique;not null"`
	UsageTime    string `json:"usagetime" gorm:"not null"`
	TabletCount  uint   `json:"no.oftablets" gorm:"not null"`
	ExpDate      string `json:"expdate" gorm:"not null"`
}
