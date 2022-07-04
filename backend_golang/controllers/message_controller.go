package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/message"
	"FinalProject/ent/predicate"
	"context"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Message defines the struct for the message
type Message struct {
	 MessageText string
	 SentDateTime string
	 ChatMessage int
	 UserMessage int
}

// MessageController defines the struct for the message controller
type MessageController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMessage handles POST requests for adding message entities
// @Summary Create message
// @Description Create message
// @ID create-message
// @Accept json
// @Produce json
// @Param message body ent.Message true "Message entity"
// @Success 200 {object} ent.Message
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /messages [post]
func (ctl *MessageController) CreateMessage(c *gin.Context) {
	obj := Message{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "message binding failed ",
		})
		return
	}

	SentDateTime, err := time.Parse(time.RFC3339, obj.SentDateTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertMessage, err := ctl.client.Message.
		Create().
		SetMessageText(obj.MessageText).
		SetSentDateTime(SentDateTime).
		SetWhatMessagesAreInThisChatID(obj.ChatMessage).
		SetWhoSendMessagesID(obj.UserMessage).
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
		"data":   insertMessage,
	})
}

// GetMessage handles GET requests to retrieve a message entity
// @Summary Get a message entity by ID
// @Description get message by ID
// @ID get-message
// @Produce json
// @Param id path int true "Message ID"
// @Success 200 {object} ent.Message
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /messages/{id} [get]
func (ctl *MessageController) GetMessage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Where(chat.HasChatUserWith( predicate.User(chat.IDEQ(int(id))))).
	getMessage, err := ctl.client.Message.
		Query().
		Where(message.HasWhatMessagesAreInThisChatWith(predicate.Chat(message.IDEQ(int(id))))).
		WithWhatMessagesAreInThisChat().
		WithWhoSendMessages().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getMessage)

}

// ListMessage handles request to get a list of message entities
// @Summary List message entities
// @Description list message entities
// @ID list-message
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Message
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /messages [get]
func (ctl *MessageController) ListMessage(c *gin.Context) {
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

	listMessage, err := ctl.client.Message.
		Query().
		Limit(limit).
		Offset(offset).
		WithWhatMessagesAreInThisChat().
		WithWhoSendMessages().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listMessage)

}

// DeleteMessage handles DELETE requests to delete a message entity
// @Summary Delete a message entity by ID
// @Description get message by ID
// @ID delete-message
// @Produce json
// @Param id path int true "Message ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /messages/{id} [delete]
func (ctl *MessageController) DeleteMessage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Message.
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

// UpdateMessage handles PUT requests to update a message entity
// @Summary Update a message entity by ID
// @Description update message by ID
// @ID update-message
// @Accept json
// @Produce json
// @Param id path int true "Message ID"
// @Param message body ent.Message true "Message entity"
// @Success 200 {object} ent.Message
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /messages/{id} [put]
func (ctl *MessageController) UpdateMessage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Message{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "message binding failed",
		})
		return
	}

	SentDateTime, err := time.Parse(time.RFC3339, obj.SentDateTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	UpdateMessage, err := ctl.client.Message.
		UpdateOneID(int(id)).
		SetMessageText(obj.MessageText).
		SetSentDateTime(SentDateTime).
		SetWhatMessagesAreInThisChatID(obj.ChatMessage).
		SetWhoSendMessagesID(obj.UserMessage).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   UpdateMessage,
	})
}

// NewMessageController creates and registers handles for the message controller
func NewMessageController(router gin.IRouter, client *ent.Client) *MessageController {
	uc := &MessageController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitMessageController registers routes to the main engine
func (ctl *MessageController) register() {
	messages := ctl.router.Group("/messages")
	messages.GET("", ctl.ListMessage)
	// CRUD
	messages.POST("", ctl.CreateMessage)
	messages.GET(":id", ctl.GetMessage)
	messages.PUT(":id", ctl.UpdateMessage)
	messages.DELETE(":id", ctl.DeleteMessage)
}
