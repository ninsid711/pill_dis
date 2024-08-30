package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	PatientId    string     `json:"patientid" gorm:"primaryKey"`
	Name         string     `json:"name" gorm:"not null"`
	Age          uint       `json:"age" gorm:"not null"`
	Gender       string     `json:"gender" gorm:"not null"`
	Illness      string     `json:"illness" gorm:"not null"`
	Password     string     `json:"-" gorm:"not null"`
	Prescription []Medicine `gorm:"foreignKey:PatientID"`
}

type Doctor struct {
	gorm.Model
	DoctorId     string    `json:"doctorid" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null; unique"`
	Age          uint      `json:"age" gorm:"not null"`
	Gender       string    `json:"gender" gorm:"not null"`
	Password     string    `json:"-" gorm:"not null"`
	PatientsList []Patient `gorm:"foreignKey:DoctorID"`
}

type Medicine struct {
	MedId           string `json:"medid" gorm:"primaryKey"`
	Name            string `json:"name" gorm:"not null;unique"`
	DosageToBeTaken uint   `json:"dosage" gorm:"not null"`
	IntakeTime      string `json:"intaketime" gorm:"not null"`
	ExpDate         string `json:"expdate" gorm:"not null"`
}
