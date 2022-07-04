package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/schedule"
	"FinalProject/ent/user"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Schedule defines the struct for the schedule
type Schedule struct {
	Activity  string
	Detail    string
	Status    string
	StartTime string
	EndTime   string
	UserId    int
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
	obj := Schedule{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "schedule binding failed",
		})
		return
	}
	if !strings.Contains(obj.StartTime, "Z") {
		obj.StartTime += "Z"
	}
	if !strings.Contains(obj.EndTime, "Z") {
		obj.EndTime += "Z"
	}

	StartTime, err := time.Parse(time.RFC3339, obj.StartTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	EndTime, err := time.Parse(time.RFC3339, obj.EndTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertSchedule, err := ctl.client.Schedule.
		Create().
		SetActivity(obj.Activity).
		SetDetail(obj.Detail).
		SetStartTime(StartTime).
		SetEndTime(EndTime).
		SetStatus(obj.Status).
		SetScheduleDoctorID(obj.UserId).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	for index := StartTime; index.Before(EndTime); index = index.Add(time.Minute * 15) {
		values := map[string]interface{}{
			"ReasonForAppointment": nil,
			"Detail":               nil,
			"Status":               "None",
			"StartTime":            index,
			"EndTime":              index.Add(time.Minute * 15),
			"DoctorId":             obj.UserId,
			"UserId":               nil,
			"ScheduleId":           insertSchedule.ID,
		}
		json_data, err := json.Marshal(values)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}

		_, err = http.Post("http://localhost:8080/api/v1/appointments", "application/json", bytes.NewBuffer(json_data))
	}

	c.JSON(200, insertSchedule)
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

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getSchedule, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(id))).
		WithDoctorSchedule().
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getSchedule)
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

	listSchedule, err := ctl.client.Schedule.Query().Limit(limit).Offset(offset).All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listSchedule)

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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getSchedule, err := ctl.client.Schedule.
		Query().
		Where(schedule.IDEQ(int(id))).
		WithScheduleAppointment().
		Only(context.Background())

	for i := range getSchedule.Edges.ScheduleAppointment {
		err = ctl.client.Appointment.DeleteOneID(int(getSchedule.Edges.ScheduleAppointment[i].ID)).Exec(context.Background())
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	err = ctl.client.Schedule.DeleteOneID(int(id)).Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	StartTime, err := time.Parse(time.RFC3339, obj.StartTime+"Z")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	EndTime, err := time.Parse(time.RFC3339, obj.EndTime+"Z")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	updateSchedule, err := ctl.client.Schedule.
		UpdateOneID(int(id)).
		SetActivity(obj.Activity).
		SetDetail(obj.Detail).
		SetStartTime(StartTime).
		SetEndTime(EndTime).
		SetStatus(obj.Status).
		SetScheduleDoctorID(obj.UserId).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, updateSchedule)
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
