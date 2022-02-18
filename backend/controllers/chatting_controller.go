package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/chatting"
	"context"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Chatting defines the struct for the chatting
type Chatting struct {
	Message  string
	DateTime string
	Whose    int
	ToWhom   int
}

// ChattingController defines the struct for the chatting controller
type ChattingController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateChatting handles POST requests for adding chatting entities
// @Summary Create chatting
// @Description Create chatting
// @ID create-chatting
// @Accept json
// @Produce json
// @Param chatting body ent.Chatting true "Chatting entity"
// @Success 200 {object} ent.Chatting
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chattings [post]
func (ctl *ChattingController) CreateChatting(c *gin.Context) {
	obj := Chatting{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "chatting binding failed ",
		})
		return
	}

	dateTime, err := time.Parse(time.RFC3339, obj.DateTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertChatting, err := ctl.client.Chatting.
		Create().
		SetMessage(obj.Message).
		SetDateTime(dateTime).
		SetWhoseIsThisMsgID(obj.Whose).
		SetChattingWithWhomID(obj.ToWhom).
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
		"data":   insertChatting,
	})
}

// GetChatting handles GET requests to retrieve a chatting entity
// @Summary Get a chatting entity by ID
// @Description get chatting by ID
// @ID get-chatting
// @Produce json
// @Param id path int true "Chatting ID"
// @Success 200 {object} ent.Chatting
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chattings/{id} [get]
func (ctl *ChattingController) GetChatting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	getChatting, err := ctl.client.Chatting.
		Query().
		Where(chatting.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getChatting)

}

// ListChatting handles request to get a list of chatting entities
// @Summary List chatting entities
// @Description list chatting entities
// @ID list-chatting
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Chatting
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chattings [get]
func (ctl *ChattingController) ListChatting(c *gin.Context) {
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

	listChatting, err := ctl.client.Chatting.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listChatting)

}

// DeleteChatting handles DELETE requests to delete a chatting entity
// @Summary Delete a chatting entity by ID
// @Description get chatting by ID
// @ID delete-chatting
// @Produce json
// @Param id path int true "Chatting ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chattings/{id} [delete]
func (ctl *ChattingController) DeleteChatting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Chatting.
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

// UpdateChatting handles PUT requests to update a chatting entity
// @Summary Update a chatting entity by ID
// @Description update chatting by ID
// @ID update-chatting
// @Accept json
// @Produce json
// @Param id path int true "Chatting ID"
// @Param chatting body ent.Chatting true "Chatting entity"
// @Success 200 {object} ent.Chatting
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chattings/{id} [put]
func (ctl *ChattingController) UpdateChatting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Chatting{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "chatting binding failed",
		})
		return
	}

	dateTime, err := time.Parse(time.RFC3339, obj.DateTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	UpdateChatting, err := ctl.client.Chatting.
		UpdateOneID(int(id)).
		SetMessage(obj.Message).
		SetDateTime(dateTime).
		SetWhoseIsThisMsgID(obj.Whose).
		SetChattingWithWhomID(obj.ToWhom).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   UpdateChatting,
	})
}

// NewChattingController creates and registers handles for the chatting controller
func NewChattingController(router gin.IRouter, client *ent.Client) *ChattingController {
	uc := &ChattingController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitChattingController registers routes to the main engine
func (ctl *ChattingController) register() {
	chattings := ctl.router.Group("/chattings")
	chattings.GET("", ctl.ListChatting)
	// CRUD
	chattings.POST("", ctl.CreateChatting)
	chattings.GET(":id", ctl.GetChatting)
	chattings.PUT(":id", ctl.UpdateChatting)
	chattings.DELETE(":id", ctl.DeleteChatting)
}
