package main

import (
	"githug.com/ninsid711/pill_dis/initializers"
	"githug.com/ninsid711/pill_dis/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Doctor{}, &models.Medicine{}, &models.Patient{})
}
