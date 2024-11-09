package main

import (
	"github.com/arfandyam/Whisper-Me/initializers"
	"github.com/arfandyam/Whisper-Me/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnDB()
}

func main(){
	 /*Pisah AutoMigrate untuk tabel yg punya constraint karena dalam prosesnya tabel 
	 terkadang belum terbuat tapi sudah ada constraint terhadap tabel lain*/
	initializers.DB.AutoMigrate(&models.User{}, &models.Authentication{})
	initializers.DB.AutoMigrate(&models.Question{})
	initializers.DB.AutoMigrate(&models.Response{})
}