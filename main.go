package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	router.GET("/getapi", GetApi)

	router.Run("0.0.0.0:8080")
}

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
}

func GetApi(c *gin.Context) {
 //   var posts User
	 var userdetails = User{
		Name:"Sripriya",
		Age:26,
		Address:"ssss",
	 }
 
    c.JSON(http.StatusOK, gin.H{"data": userdetails})
	fmt.Println("testing")
}