package controllers

import (
	"FinalProject/ent"

	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Disease defines the struct for the disease
type Disease struct {
	Name   		string
	Symptoms	string
}

// DiseaseController defines the struct for the disease controller
type DiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDisease handles POST requests for adding disease entities
// @Summary Create disease
// @Description Create disease
// @ID create-disease
// @Accept json
// @Produce json
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [post]
func (ctl *DiseaseController) CreateDisease(c *gin.Context) {
	db := databaseConnection()
	obj := Disease{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "disease binding failed",
		})
		return
	}

	insertDisease, err := db.Prepare("INSERT INTO disease(name, symptoms) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertDisease.Exec(obj.Name, obj.Symptoms)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	defer db.Close()

	
	c.JSON(200, gin.H{
		"status" : true,
		"data": insertDisease,
	})
}

// GetDisease handles GET requests to retrieve a disease entity
// @Summary Get a disease entity by ID
// @Description get disease by ID
// @ID get-disease
// @Produce json
// @Param id path int true "Disease ID"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [get]
func (ctl *DiseaseController) GetDisease(c *gin.Context) {
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetDisease, err := db.Query("SELECT name, symptoms FROM disease WHERE id_disease=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetDisease.Next() {
		var disease Disease
		// for each row, scan the result into our tag composite object
		err = resultsGetDisease.Scan(&disease.Name, &disease.Symptoms)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, disease)

	}
}

// ListDisease handles request to get a list of disease entities
// @Summary List disease entities
// @Description list disease entities
// @ID list-disease
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [get]
func (ctl *DiseaseController) ListDisease(c *gin.Context) {
	db := databaseConnection()
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}
	resultsListDisease, err := db.Query("SELECT name, symptoms FROM disease limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListDisease.Next() {
		var disease Disease
		// for each row, scan the result into our tag composite object
		err = resultsListDisease.Scan(&disease.Name, &disease.Symptoms)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, disease)
	}

	
}

// DeleteDisease handles DELETE requests to delete a disease entity
// @Summary Delete a disease entity by ID
// @Description get disease by ID
// @ID delete-disease
// @Produce json
// @Param id path int true "Disease ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [delete]
func (ctl *DiseaseController) DeleteDisease(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteDisease, err := db.Prepare("DELETE FROM disease WHERE id_disease=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteDisease.Exec(id)

    defer db.Close()

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateDisease handles PUT requests to update a disease entity
// @Summary Update a disease entity by ID
// @Description update disease by ID
// @ID update-disease
// @Accept json
// @Produce json
// @Param id path int true "Disease ID"
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [put]
func (ctl *DiseaseController) UpdateDisease(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "disease binding failed",
		})
		return
	}

	UpdateDisease, err := db.Prepare("UPDATE disease SET name=?,symptoms=? WHERE id_disease=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateDisease.Exec(obj.Name, obj.Symptoms, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateDisease,
	})
}

// NewDiseaseController creates and registers handles for the disease controller
func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	uc := &DiseaseController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDiseaseController registers routes to the main engine
func (ctl *DiseaseController) register() {
	diseases := ctl.router.Group("/diseases")
	diseases.GET("", ctl.ListDisease)
	// CRUD
	diseases.POST("", ctl.CreateDisease)
	diseases.GET(":id", ctl.GetDisease)
	diseases.PUT(":id", ctl.UpdateDisease)
	diseases.DELETE(":id", ctl.DeleteDisease)
}
