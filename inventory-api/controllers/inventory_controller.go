package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"inventory/models"
	"inventory/pkg/ids"
)

func GetParts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Parts)
}

func PostParts(c *gin.Context) {
	var newPart models.Part

	if err := c.BindJSON(&newPart); err != nil {
		log.Fatal("Failed to create a new part: ", err)
	}

	// una vez parseado, checkear si los datos estan completos
	if newPart.Name == "" || newPart.Description == "" || newPart.Quantity == 0 || newPart.Location == "" || newPart.Minimun_Stock == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Insert all fields to add a part")
		return
	}

	// si no estas usando una DB, hardcodear el identity property.
	newPart.ID = models.Parts[len(models.Parts)-1].ID+1

	models.Parts = append(models.Parts, newPart)

	c.IndentedJSON(http.StatusOK, models.Parts)
}

func GetPartByID(c *gin.Context) {
	id := c.Param("id") 

	idInt, err := ids.StringToInt(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	for _, a := range models.Parts {
		if a.ID == idInt {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Part not found"})
}

func UpdatePart(c *gin.Context) {
	id := c.Param("id")

	idInt, err := ids.StringToInt(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	var updatedPart models.Part

	if err := c.BindJSON(&updatedPart); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to parse JSON"})
		return
	}

	// Buscar y actualizar la parte existente
	for i, a := range models.Parts {
		if a.ID == idInt {
			updatedPart.ID = idInt
			models.Parts[i] = updatedPart
			c.IndentedJSON(http.StatusOK, updatedPart)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, models.Parts)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Part not found"})
}


func DeletePart(c *gin.Context) {
	id := c.Param("id")

	idInt, err := ids.StringToInt(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	for i, a := range models.Parts {
		if a.ID == idInt {
			models.Parts = append(models.Parts[:i], models.Parts[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Part deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Part not found"})
}