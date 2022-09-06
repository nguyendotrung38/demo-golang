package main

/**
 Author: Trung Nguyen
 Description: Create a simple demo API application for list, read, create API
 */

import (
	"connector"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Create data type called "member"
type member struct {
	ID		int	`json:"id"`
	Name	string	`json:"name"`
	Role	string	`json:"role"`
}

func main()  {
	// Init routes
	router := gin.Default()
	router.GET("/members", getMembers)
	router.GET("/members/:id", getMemberById)
	//router.POST("/albums", postAlbum)

	// Run local server at port 8080
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getMemberById(c *gin.Context) {
	memberId := c.Param("id")
	row := connector.Db.QueryRow("SELECT * FROM members WHERE id = ?", memberId)
	var member member
	err := row.Scan(&member.ID, &member.Name, &member.Role)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Error when getting member who has id by %s", memberId)})
	}

	c.IndentedJSON(http.StatusOK, member)
}

func getMembers(c *gin.Context) {
	rows := connector.Query("SELECT * FROM members")

	var members []member
	for rows.Next() {
		var row member
		if err := rows.Scan(&row.ID, &row.Name, &row.Role); err != nil {
			log.Fatal(err)
		}
		members = append(members, row)
	}

	c.IndentedJSON(http.StatusOK, members)
}

//func postAlbum(c *gin.Context) {
//	var newAlbum album
//
//	if err := c.BindJSON(&newAlbum); err != nil {
//		return
//	}
//
//	albums = append(albums, newAlbum)
//	c.IndentedJSON(http.StatusCreated, newAlbum)
//}