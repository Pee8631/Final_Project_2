package controllers

import (
	"FinalProject/ent"
	"fmt"
	"strconv"
	"log"

	"github.com/gin-gonic/gin"
)
//"FinalProject/ent/hospital"
// User defines the struct for the user
type Hospital struct {
	Name	string
}

// HospitalController defines the struct for the hospital controller
type HospitalController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateHospital handles POST requests for adding hospital entities
// @Summary Create hospital
// @Description Create hospital
// @ID create-hospital
// @Accept json
// @Produce json
// @Param hospital body ent.Hospital true "Hospital entity"
// @Success 200 {object} ent.Hospital
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals [post]
func (ctl *HospitalController) CreateHospital(c *gin.Context) {
	db := databaseConnection()
	obj := Hospital{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	insertHospital, err := db.Prepare("INSERT INTO hospital(name) VALUES (?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertHospital.Exec(obj.Name)
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
		"data": insertHospital,
	})
}

// GetHospital handles GET requests to retrieve a hospital entity
// @Summary Get a hospital entity by ID
// @Description get hospital by ID
// @ID get-hospital
// @Produce json
// @Param id path int true "Hospital ID"
// @Success 200 {object} ent.Hospital
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals/{id} [get]
func (ctl *HospitalController) GetHospital(c *gin.Context) {
	db := databaseConnection()
	log.Println("Connected")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Println(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	resultsGetHospital, err := db.Query("SELECT name FROM hospital WHERE id_hospital=1")
	log.Println(resultsGetHospital)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}
	
	for resultsGetHospital.Next() {
		var hospital Hospital
		err = resultsGetHospital.Scan(&hospital.Name)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		// for each row, scan the result into our tag composite object
		log.Println("Username: " + hospital.Name)

		defer db.Close()

		c.JSON(200, hospital)
	}
}

// ListHospital handles request to get a list of hospital entities
// @Summary List hospital entities
// @Description list hospital entities
// @ID list-hospital
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Hospital
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals [get]
func (ctl *HospitalController) ListHospital(c *gin.Context) {
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
	resultsListUser, err := db.Query("SELECT name FROM hospital limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListUser.Next() {
		var hospital Hospital
		// for each row, scan the result into our tag composite object
		err = resultsListUser.Scan(&hospital.Name)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, hospital)
	}
}

// DeleteHospital handles DELETE requests to delete a hospital entity
// @Summary Delete a hospital entity by ID
// @Description get hospital by ID
// @ID delete-hospital
// @Produce json
// @Param id path int true "Hospital ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals/{id} [delete]
func (ctl *HospitalController) DeleteHospital(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteHospital, err := db.Prepare("DELETE FROM hospital WHERE id_hospital=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteHospital.Exec(id)

    defer db.Close()

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateHospital handles PUT requests to update a hospital entity
// @Summary Update a hospital entity by ID
// @Description update hospital by ID
// @ID update-hospital
// @Accept json
// @Produce json
// @Param id path int true "Hospital ID"
// @Param hospital body ent.Hospital true "Hospital entity"
// @Success 200 {object} ent.Hospital
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals/{id} [put]
func (ctl *HospitalController) UpdateHospital(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Hospital{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	UpdateHospital, err := db.Prepare("UPDATE user SET name=? WHERE id_hospital=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateHospital.Exec(obj.Name, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateHospital,
	})
}

// NewHospitalController creates and registers handles for the hospital controller
func NewHospitalController(router gin.IRouter, client *ent.Client) *HospitalController {
	uc := &HospitalController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitHospitalController registers routes to the main engine
func (ctl *HospitalController) register() {
	hospitals := ctl.router.Group("/hospitals")
	hospitals.GET("", ctl.ListHospital)
	// CRUD
	hospitals.POST("", ctl.CreateHospital)
	hospitals.GET(":id", ctl.GetHospital)
	hospitals.PUT(":id", ctl.UpdateHospital)
	hospitals.DELETE(":id", ctl.DeleteHospital)
}
