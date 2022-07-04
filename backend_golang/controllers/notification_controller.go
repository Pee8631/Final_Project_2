package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/notification"
	"FinalProject/ent/user"
	"time"

	"context"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Notification defines the struct for the notification
type Notification struct {
	Message       string
	SenderId      int
	RecipientId   int
	AppointmentId int
}

// NotificationController defines the struct for the notification controller
type NotificationController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateNotification handles POST requests for adding notification entities
// @Summary Create notification
// @Description Create notification
// @ID create-notification
// @Accept json
// @Produce json
// @Param notification body ent.Notification true "Notification entity"
// @Success 200 {object} ent.Notification
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /notifications [post]
func (ctl *NotificationController) CreateNotification(c *gin.Context) {
	obj := Notification{}
	if err := c.ShouldBind(&obj); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"error": "notification binding failed",
		})
		return
	}

	getSender, err := ctl.client.User.
		Query().
		Where(user.IDEQ(obj.SenderId)).
		Only(context.Background())

	if err != nil {
		fmt.Println(err.Error() + "Sender: %s", obj.SenderId)
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Sender Not found",
		})
		//"Invalid Username"
		return
	}

	getRecipient, err := ctl.client.User.
		Query().
		Where(user.IDEQ(obj.RecipientId)).
		Only(context.Background())

	if err != nil {
		fmt.Println(err.Error() + "Recipient: %s", obj.RecipientId)
		c.JSON(400, gin.H{
			"status": false,
			"error":  "Recipient Not found",
		})
		//"Invalid Username"
		return
	}

	CreatedDate := time.Now()
	insertNotification, err := ctl.client.Notification.
		Create().
		SetSenderId(obj.SenderId).
		SetRecipientId(obj.RecipientId).
		SetMessage(obj.Message).
		SetCreatedDate(CreatedDate).
		SetAppointmentId(obj.AppointmentId).
		AddUserNotification(getSender).
		AddUserNotification(getRecipient).
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
		"data":   insertNotification,
	})
}

// GetNotification handles GET requests to retrieve a notification entity
// @Summary Get a notification entity by ID
// @Description get notification by ID
// @ID get-notification
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} ent.Notification
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /notifications/{id} [get]
func (ctl *NotificationController) GetNotification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	getNotification, err := ctl.client.Notification.
		Query().
		Where(notification.RecipientIdEQ(int(id))).
		WithUserNotification().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getNotification)

}

// ListNotification handles request to get a list of notification entities
// @Summary List notification entities
// @Description list notification entities
// @ID list-notification
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Notification
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /notifications [get]
func (ctl *NotificationController) ListNotification(c *gin.Context) {


	listNotification, err := ctl.client.Notification.
		Query().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listNotification)

}

// DeleteNotification handles DELETE requests to delete a notification entity
// @Summary Delete a notification entity by ID
// @Description get notification by ID
// @ID delete-notification
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /notifications/{id} [delete]
func (ctl *NotificationController) DeleteNotification(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Notification.
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

// UpdateNotification handles PUT requests to update a notification entity
// @Summary Update a notification entity by ID
// @Description update notification by ID
// @ID update-notification
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Param notification body ent.Notification true "Notification entity"
// @Success 200 {object} ent.Notification
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /notifications/{id} [put]
func (ctl *NotificationController) UpdateNotification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Notification{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "notification binding failed",
		})
		return
	}

	updateNotification, err := ctl.client.Notification.
		UpdateOneID(int(id)).
		SetRecipientId(obj.RecipientId).
		SetMessage(obj.Message).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateNotification,
	})
}

// NewNotificationController creates and registers handles for the notification controller
func NewNotificationController(router gin.IRouter, client *ent.Client) *NotificationController {
	uc := &NotificationController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitNotificationController registers routes to the main engine
func (ctl *NotificationController) register() {
	notifications := ctl.router.Group("/notifications")
	notifications.GET("", ctl.ListNotification)
	// CRUD
	notifications.POST("", ctl.CreateNotification)
	notifications.GET(":id", ctl.GetNotification)
	notifications.PUT(":id", ctl.UpdateNotification)
	notifications.DELETE(":id", ctl.DeleteNotification)
}
