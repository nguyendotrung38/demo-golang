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
	"strconv"
)

// Create data type called "member"
type member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func main() {
	// Init routes
	router := gin.Default()
	router.GET("/members", getMembers)
	router.GET("/members/:id", getMemberById)
	router.POST("/members", addMember)
	router.PATCH("/members/:id", updateMember)
	router.DELETE("/members/:id", deleteMember)

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

func addMember(c *gin.Context) {
	var newMember member
	if err := c.BindJSON(&newMember); err != nil {
		return
	}

	result, err := connector.Db.Exec("INSERT INTO members (name, role) VALUES (?, ?)", newMember.Name, newMember.Role)
	if err != nil {
		log.Fatal("Error when saving member to database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	newMember.ID = int(id)
	c.IndentedJSON(http.StatusCreated, newMember)
}

func updateMember(c *gin.Context) {
	var updateInfo member
	memberId := c.Param("id")
	if err := c.BindJSON(&updateInfo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error when getting bind update value"})
	}
	_, err := connector.Db.Exec("UPDATE members SET name = ?, role = ? WHERE id = ?",
		updateInfo.Name,
		updateInfo.Role,
		memberId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error when updating member info"})
	}
	updateInfo.ID, err = strconv.Atoi(memberId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error converting member id to int"})
	}
	c.IndentedJSON(http.StatusOK, updateInfo)
}

func deleteMember(c *gin.Context) {
	memberId := c.Param("id")
	result, err := connector.Db.Exec("DELETE FROM members WHERE id = ?", memberId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error when deleting member"})
	}
	rowCount, _ := result.RowsAffected()
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Affected %d row(s)", rowCount)})
}
