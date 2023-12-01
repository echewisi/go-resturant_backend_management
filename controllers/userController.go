package controllers

import (
	"github.com/gin-gonic/gin"
)

func getUsers() gin.HandlerFunc{
	return func(ctx *gin.Context){

	}
}

func getUser() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}

func signupUser() gin.HandlerFunc{
	return func(ctx *gin.Context){

	}
}

func loginUser() gin.HandlerFunc{
	return func(ctx *gin.Context){

	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providePassword string)(bool, string){

}