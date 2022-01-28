package controllers

import (
	"FinalProject/ent"

	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Schedule defines the struct for the schedule
type Schedule struct {
	Activity   	string
	Detail		string
	Status		string
	User		int
}

// ScheduleController defines the struct for the schedule controller
type ScheduleController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSchedule handles POST requests for adding schedule entities
// @Summary Create schedule
// @Description Create schedule
// @ID create-schedule
// @Accept json
// @Produce json
// @Param schedule body ent.Schedule true "Schedule entity"
// @Success 200 {object} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules [post]
func (ctl *ScheduleController) CreateSchedule(c *gin.Context) {
	db := databaseConnection()
	obj := Schedule{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "schedule binding failed",
		})
		return
	}

	insertSchedule, err := db.Prepare("INSERT INTO schedule(activity, detail, status, id_user) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertSchedule.Exec(obj.Activity, obj.Detail, obj.Status, obj.User)
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
		"data": insertSchedule,
	})
}

// GetSchedule handles GET requests to retrieve a schedule entity
// @Summary Get a schedule entity by ID
// @Description get schedule by ID
// @ID get-schedule
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [get]
func (ctl *ScheduleController) GetSchedule(c *gin.Context) {
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetSchedule, err := db.Query("SELECT activity, detail, status, id_user FROM schedule WHERE id_schedule=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetSchedule.Next() {
		var schedule Schedule
		// for each row, scan the result into our tag composite object
		err = resultsGetSchedule.Scan(&schedule.Activity, &schedule.Detail, &schedule.Status, &schedule.User)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, schedule)

	}
}

// ListSchedule handles request to get a list of schedule entities
// @Summary List schedule entities
// @Description list schedule entities
// @ID list-schedule
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules [get]
func (ctl *ScheduleController) ListSchedule(c *gin.Context) {
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
	resultsListSchedule, err := db.Query("SELECT activity, detail, status, id_user FROM schedule limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListSchedule.Next() {
		var schedule Schedule
		// for each row, scan the result into our tag composite object
		err = resultsListSchedule.Scan(&schedule.Activity, &schedule.Detail, &schedule.Status, &schedule.User)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, schedule)
	}

	
}

// DeleteSchedule handles DELETE requests to delete a schedule entity
// @Summary Delete a schedule entity by ID
// @Description get schedule by ID
// @ID delete-schedule
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [delete]
func (ctl *ScheduleController) DeleteSchedule(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteSchedule, err := db.Prepare("DELETE FROM schedule WHERE id_schedule=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteSchedule.Exec(id)

    defer db.Close()

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateSchedule handles PUT requests to update a schedule entity
// @Summary Update a schedule entity by ID
// @Description update schedule by ID
// @ID update-schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Param schedule body ent.Schedule true "Schedule entity"
// @Success 200 {object} ent.Schedule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [put]
func (ctl *ScheduleController) UpdateSchedule(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Schedule{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "schedule binding failed",
		})
		return
	}

	UpdateSchedule, err := db.Prepare("UPDATE schedule SET activity=?, detail=?, status=?, id_user=? WHERE id_schedule=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateSchedule.Exec(obj.Activity, obj.Detail, obj.Status, obj.User, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateSchedule,
	})
}

// NewScheduleController creates and registers handles for the schedule controller
func NewScheduleController(router gin.IRouter, client *ent.Client) *ScheduleController {
	uc := &ScheduleController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitScheduleController registers routes to the main engine
func (ctl *ScheduleController) register() {
	schedules := ctl.router.Group("/schedules")
	schedules.GET("", ctl.ListSchedule)
	// CRUD
	schedules.POST("", ctl.CreateSchedule)
	schedules.GET(":id", ctl.GetSchedule)
	schedules.PUT(":id", ctl.UpdateSchedule)
	schedules.DELETE(":id", ctl.DeleteSchedule)
}
