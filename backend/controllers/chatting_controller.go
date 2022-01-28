package controllers

import (
	"FinalProject/ent"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Chatting defines the struct for the chatting
type Chatting struct {
	Message 	string 
	DateTime 	string 
	Whose 		int 
	ToWhom 		int
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
	db := databaseConnection()
	obj := Chatting{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "chatting binding failed " ,
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

	insertChatting, err := db.Prepare("INSERT INTO chatting(message, datetime, whose, towhom) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertChatting.Exec(obj.Message, dateTime, obj.Whose, obj.ToWhom)
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
		"data": insertChatting,
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
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetChatting, err := db.Query("SELECT message, datetime, whose, towhom FROM chatting WHERE id_chatting=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetChatting.Next() {
		var chatting Chatting
		// for each row, scan the result into our tag composite object
		err = resultsGetChatting.Scan(&chatting.Message, &chatting.DateTime, &chatting.ToWhom, &chatting.Whose)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, chatting)

	}
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
	resultsListChatting, err := db.Query("SELECT message, datetime, whose, towhom FROM chatting limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListChatting.Next() {
		var chatting Chatting
		// for each row, scan the result into our tag composite object
		err = resultsListChatting.Scan(&chatting.Message, &chatting.DateTime, &chatting.ToWhom, &chatting.Whose)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, chatting)
	}

	
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
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteChatting, err := db.Prepare("DELETE FROM chatting WHERE id_chatting=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteChatting.Exec(id)

    defer db.Close()

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
	db := databaseConnection()
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
			"status" : false,
			"error": err,
		})
		return
	}

	UpdateChatting, err := db.Prepare("UPDATE chatting SET message=?, datetime=?, whose=?, towhom=? WHERE id_chatting=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateChatting.Exec(obj.Message, dateTime, obj.Whose, obj.ToWhom, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateChatting,
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
