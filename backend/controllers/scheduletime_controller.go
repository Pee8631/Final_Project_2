package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/scheduletime"
	"context"
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

	insertScheduleTime, err := ctl.client.ScheduleTime.Create().
	SetStartTime(startTime).
	SetStopTime(stopTime).
	SetWhatTimeIsTheScheduleID(obj.Schedule).
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	GetScheduleTime, err := ctl.client.ScheduleTime.
		Query().
		Where(scheduletime.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, GetScheduleTime)
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

	listScheduleTime, err := ctl.client.ScheduleTime.Query().Limit(limit).Offset(offset).All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listScheduleTime)

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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ScheduleTime.DeleteOneID(int(id)).Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	updateScheduleTime, err := ctl.client.ScheduleTime.
	UpdateOneID(int(id)).
	SetStartTime(startTime).
	SetStopTime(stopTime).
	SetNillableWhatTimeIsTheScheduleID(&obj.Schedule).
	Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateScheduleTime,
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
