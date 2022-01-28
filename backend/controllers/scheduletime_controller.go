package controllers

import (
	"FinalProject/ent"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ScheduleTime defines the struct for the scheduletime
type ScheduleTime struct {
	StartTime 	string
	StopTime  	string
	Schedule	int
}

// ScheduleTimeController defines the struct for the scheduletime controller
type ScheduleTimeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateScheduleTime handles POST requests for adding scheduletime entities
// @Summary Create scheduletime
// @Description Create scheduletime
// @ID create-scheduletime
// @Accept json
// @Produce json
// @Param scheduletime body ent.ScheduleTime true "ScheduleTime entity"
// @Success 200 {object} ent.ScheduleTime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scheduletimes [post]
func (ctl *ScheduleTimeController) CreateScheduleTime(c *gin.Context) {
	db := databaseConnection()
	obj := ScheduleTime{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "scheduletime binding failed",
		})
		return
	}

	startTime, err := time.Parse(time.RFC3339, obj.StartTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	stopTime, err := time.Parse(time.RFC3339, obj.StopTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	insertScheduleTime, err := db.Prepare("INSERT INTO scheduletime(starttime, stoptime, id_schedule) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}
	insertScheduleTime.Exec(startTime, stopTime, obj.Schedule)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	defer db.Close()

	c.JSON(200, gin.H{
		"status": true,
		"data":   insertScheduleTime,
	})
}

// GetScheduleTime handles GET requests to retrieve a scheduletime entity
// @Summary Get a scheduletime entity by ID
// @Description get scheduletime by ID
// @ID get-scheduletime
// @Produce json
// @Param id path int true "ScheduleTime ID"
// @Success 200 {object} ent.ScheduleTime
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scheduletimes/{id} [get]
func (ctl *ScheduleTimeController) GetScheduleTime(c *gin.Context) {
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetScheduleTime, err := db.Query("SELECT starttime, stoptime, id_schedule FROM scheduletime WHERE id_scheduletime=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetScheduleTime.Next() {
		var scheduletime ScheduleTime
		// for each row, scan the result into our tag composite object
		err = resultsGetScheduleTime.Scan(&scheduletime.StartTime, &scheduletime.StopTime, &scheduletime.Schedule)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, scheduletime)

	}
}

// ListScheduleTime handles request to get a list of scheduletime entities
// @Summary List scheduletime entities
// @Description list scheduletime entities
// @ID list-scheduletime
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ScheduleTime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scheduletimes [get]
func (ctl *ScheduleTimeController) ListScheduleTime(c *gin.Context) {
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
	resultsListScheduleTime, err := db.Query("SELECT starttime, stoptime, id_schedule FROM scheduletime limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListScheduleTime.Next() {
		var scheduletime ScheduleTime
		// for each row, scan the result into our tag composite object
		err = resultsListScheduleTime.Scan(&scheduletime.StartTime, &scheduletime.StopTime, &scheduletime.Schedule)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, scheduletime)
	}

}

// DeleteScheduleTime handles DELETE requests to delete a scheduletime entity
// @Summary Delete a scheduletime entity by ID
// @Description get scheduletime by ID
// @ID delete-scheduletime
// @Produce json
// @Param id path int true "ScheduleTime ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scheduletimes/{id} [delete]
func (ctl *ScheduleTimeController) DeleteScheduleTime(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteScheduleTime, err := db.Prepare("DELETE FROM scheduletime WHERE id_scheduletime=?")
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	deleteScheduleTime.Exec(id)

	defer db.Close()

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateScheduleTime handles PUT requests to update a scheduletime entity
// @Summary Update a scheduletime entity by ID
// @Description update scheduletime by ID
// @ID update-scheduletime
// @Accept json
// @Produce json
// @Param id path int true "ScheduleTime ID"
// @Param scheduletime body ent.ScheduleTime true "ScheduleTime entity"
// @Success 200 {object} ent.ScheduleTime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scheduletimes/{id} [put]
func (ctl *ScheduleTimeController) UpdateScheduleTime(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ScheduleTime{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "scheduletime binding failed",
		})
		return
	}

	startTime, err := time.Parse(time.RFC3339, obj.StartTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	stopTime, err := time.Parse(time.RFC3339, obj.StopTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	UpdateScheduleTime, err := db.Prepare("UPDATE scheduletime SET starttime=?, stoptime=?, id_schedule=? WHERE id_scheduletime=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateScheduleTime.Exec(startTime, stopTime, obj.Schedule, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status": true,
		"data":   UpdateScheduleTime,
	})
}

// NewScheduleTimeController creates and registers handles for the scheduletime controller
func NewScheduleTimeController(router gin.IRouter, client *ent.Client) *ScheduleTimeController {
	uc := &ScheduleTimeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitScheduleTimeController registers routes to the main engine
func (ctl *ScheduleTimeController) register() {
	scheduletimes := ctl.router.Group("/scheduletimes")
	scheduletimes.GET("", ctl.ListScheduleTime)
	// CRUD
	scheduletimes.POST("", ctl.CreateScheduleTime)
	scheduletimes.GET(":id", ctl.GetScheduleTime)
	scheduletimes.PUT(":id", ctl.UpdateScheduleTime)
	scheduletimes.DELETE(":id", ctl.DeleteScheduleTime)
}
