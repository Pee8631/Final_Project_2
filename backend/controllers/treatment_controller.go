package controllers

import (
	"FinalProject/ent"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Treatment defines the struct for the treatment
type Treatment struct {
	TreatmentRecord		string
	DateTime			string
	TakeTime			float64
	Physician			int
	Patient				int
}

// TreatmentController defines the struct for the treatment controller
type TreatmentController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTreatment handles POST requests for adding treatment entities
// @Summary Create treatment
// @Description Create treatment
// @ID create-treatment
// @Accept json
// @Produce json
// @Param treatment body ent.Treatment true "Treatment entity"
// @Success 200 {object} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments [post]
func (ctl *TreatmentController) CreateTreatment(c *gin.Context) {
	db := databaseConnection()
	obj := Treatment{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "treatment binding failed",
		})
		return
	}

	dateTime, err := time.Parse(time.RFC3339, obj.DateTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	insertTreatment, err := db.Prepare("INSERT INTO treatment(treatmentrecord, datetime, taketime, physician, patient) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertTreatment.Exec(obj.TreatmentRecord, dateTime, obj.TakeTime, obj.Physician, obj.Patient)
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
		"data": insertTreatment,
	})
}

// GetTreatment handles GET requests to retrieve a treatment entity
// @Summary Get a treatment entity by ID
// @Description get treatment by ID
// @ID get-treatment
// @Produce json
// @Param id path int true "Treatment ID"
// @Success 200 {object} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments/{id} [get]
func (ctl *TreatmentController) GetTreatment(c *gin.Context) {
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetTreatment, err := db.Query("SELECT treatmentrecord, datetime, taketime, physician, patient FROM treatment WHERE id_treatment=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetTreatment.Next() {
		var treatment Treatment
		// for each row, scan the result into our tag composite object
		err = resultsGetTreatment.Scan(&treatment.TreatmentRecord, &treatment.DateTime, &treatment.TakeTime, &treatment.Physician, &treatment.Patient)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, treatment)

	}
}

// ListTreatment handles request to get a list of treatment entities
// @Summary List treatment entities
// @Description list treatment entities
// @ID list-treatment
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments [get]
func (ctl *TreatmentController) ListTreatment(c *gin.Context) {
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
	resultsListTreatment, err := db.Query("SELECT treatmentrecord, datetime, taketime, physician, patient FROM treatment limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListTreatment.Next() {
		var treatment Treatment
		// for each row, scan the result into our tag composite object
		err = resultsListTreatment.Scan(&treatment.TreatmentRecord, &treatment.DateTime, &treatment.TakeTime, &treatment.Physician, &treatment.Patient)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, treatment)
	}

	
}

// DeleteTreatment handles DELETE requests to delete a treatment entity
// @Summary Delete a treatment entity by ID
// @Description get treatment by ID
// @ID delete-treatment
// @Produce json
// @Param id path int true "Treatment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments/{id} [delete]
func (ctl *TreatmentController) DeleteTreatment(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteTreatment, err := db.Prepare("DELETE FROM treatment WHERE id_treatment=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteTreatment.Exec(id)

    defer db.Close()

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTreatment handles PUT requests to update a treatment entity
// @Summary Update a treatment entity by ID
// @Description update treatment by ID
// @ID update-treatment
// @Accept json
// @Produce json
// @Param id path int true "Treatment ID"
// @Param treatment body ent.Treatment true "Treatment entity"
// @Success 200 {object} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments/{id} [put]
func (ctl *TreatmentController) UpdateTreatment(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Treatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "treatment binding failed",
		})
		return
	}

	dateTime, err := time.Parse(time.RFC3339, obj.DateTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	UpdateTreatment, err := db.Prepare("UPDATE treatment SET treatmentrecord=?, datetime=?, taketime=?, physician=?, patient=? WHERE id_treatment=?")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	/*
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}*/

	UpdateTreatment.Exec(obj.TreatmentRecord, dateTime, obj.TakeTime, obj.Physician, obj.Patient, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateTreatment,
	})
}

// NewTreatmentController creates and registers handles for the treatment controller
func NewTreatmentController(router gin.IRouter, client *ent.Client) *TreatmentController {
	uc := &TreatmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTreatmentController registers routes to the main engine
func (ctl *TreatmentController) register() {
	treatments := ctl.router.Group("/treatments")
	treatments.GET("", ctl.ListTreatment)
	// CRUD
	treatments.POST("", ctl.CreateTreatment)
	treatments.GET(":id", ctl.GetTreatment)
	treatments.PUT(":id", ctl.UpdateTreatment)
	treatments.DELETE(":id", ctl.DeleteTreatment)
}
