package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/hospital"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//"FinalProject/ent/hospital"
// User defines the struct for the user
type Hospital struct {
	Name string
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
	obj := Hospital{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	insertHospital, err := ctl.client.Hospital.
		Create().
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   insertHospital,
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Println(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getHospital, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getHospital)
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

	listHospitals, err := ctl.client.Hospital.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listHospitals)
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Hospital.
		DeleteOneID(int(id)).
		Exec(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	updateHospital, err := ctl.client.Hospital.UpdateOneID(int(id)).SetName(obj.Name).Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateHospital,
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
