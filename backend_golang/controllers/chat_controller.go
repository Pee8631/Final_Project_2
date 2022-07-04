package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/chat"
	"FinalProject/ent/predicate"
	"FinalProject/ent/user"
	"context"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Chat defines the struct for the chat
type Chat struct {
	ChatRoomName string
	DoctorId     int
	UserId       int
}

// ChatController defines the struct for the chat controller
type ChatController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateChat handles POST requests for adding chat entities
// @Summary Create chat
// @Description Create chat
// @ID create-chat
// @Accept json
// @Produce json
// @Param chat body ent.Chat true "Chat entity"
// @Success 200 {object} ent.Chat
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chats [post]
func (ctl *ChatController) CreateChat(c *gin.Context) {
	obj := Chat{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "chat binding failed ",
		})
		return
	}
	if obj.ChatRoomName == "" {
		getUser, err := ctl.client.User.
			Query().
			Where(user.IDEQ(int(obj.UserId))).
			WithUserHaveRole().
			Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		getDoctor, err := ctl.client.User.
			Query().
			Where(user.IDEQ(int(obj.DoctorId))).
			WithUserHaveRole().
			Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		insertChat, err := ctl.client.Chat.
			Create().
			SetChatRoomName(getUser.Username + "-" + getDoctor.Username).
			AddChatUserIDs(obj.DoctorId).
			AddChatUserIDs(obj.UserId).
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
			"data":   insertChat,
		})
	} else {
		insertChat, err := ctl.client.Chat.
			Create().
			SetChatRoomName(obj.ChatRoomName).
			AddChatUserIDs(obj.DoctorId).
			AddChatUserIDs(obj.UserId).
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
			"data":   insertChat,
		})
	}

}

// GetChat handles GET requests to retrieve a chat entity
// @Summary Get a chat entity by ID
// @Description get chat by ID
// @ID get-chat
// @Produce json
// @Param id path int true "Chat ID"
// @Success 200 {object} ent.Chat
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chats/chatroom/{id} [get]
func (ctl *ChatController) GetChat(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	getChat, err := ctl.client.Chat.
		Query().
		Where(chat.IDEQ(int(id))).
		WithChatMessage().
		WithChatUser().
		WithChatAppointment().
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	lockedChat := true
	timeNow := time.Now().UTC().Add(time.Hour * 7)
	for _, App := range getChat.Edges.ChatAppointment {

		if timeNow.After(App.StartTime.UTC()) && timeNow.Before(App.EndTime.UTC()) {
			lockedChat = false
			getChat, err = ctl.client.Chat.
				Query().
				Where(chat.IDEQ(int(id))).
				WithChatMessage().
				WithChatUser().
				Only(context.Background())
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return // proper error handling instead of panic in your app
			}
			break

		} else {
			lockedChat = true
		}
	}

	_, err = ctl.client.Chat.
		UpdateOneID(int(getChat.ID)).
		SetIsLockChat(lockedChat).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	getChat, err = ctl.client.Chat.
		Query().
		Where(chat.IDEQ(int(id))).
		WithChatMessage().
		WithChatUser().
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}
	c.JSON(200, getChat)

}

// ListChat handles request to get a list of chat entities
// @Summary List chat entities
// @Description list chat entities
// @ID list-chat
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Chat
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chats/{id} [get]
func (ctl *ChatController) ListChat(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	listChat, err := ctl.client.Chat.
		Query().
		Where(chat.HasChatUserWith(predicate.User(chat.IDEQ(int(id))))).
		WithChatUser().
		WithChatMessage().
		WithChatAppointment().
		All(context.Background())

	lockedChat := true
	timeNow := time.Now().UTC().Add(time.Hour * 7)
	for i, chat := range listChat {
		for _, App := range listChat[i].Edges.ChatAppointment {
			if timeNow.After(App.StartTime.UTC()) && timeNow.Before(App.EndTime.UTC()) {
				lockedChat = false

				_, err := ctl.client.Chat.
					UpdateOneID(int(chat.ID)).
					SetIsLockChat(lockedChat).
					Save(context.Background())
				if err != nil {
					c.JSON(400, gin.H{"error": "update failed"})
					return
				}

				break
			} else {
				lockedChat = true
			}

			_, err := ctl.client.Chat.
				UpdateOneID(int(chat.ID)).
				SetIsLockChat(lockedChat).
				Save(context.Background())
			if err != nil {
				c.JSON(400, gin.H{"error": "update failed"})
				return
			}
		}

	}

	listChat, err = ctl.client.Chat.
		Query().
		Where(chat.HasChatUserWith(predicate.User(chat.IDEQ(int(id))))).
		WithChatUser().
		WithChatMessage().
		WithChatAppointment().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listChat)

}

// DeleteChat handles DELETE requests to delete a chat entity
// @Summary Delete a chat entity by ID
// @Description get chat by ID
// @ID delete-chat
// @Produce json
// @Param id path int true "Chat ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chats/{id} [delete]
func (ctl *ChatController) DeleteChat(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Chat.
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

// UpdateChat handles PUT requests to update a chat entity
// @Summary Update a chat entity by ID
// @Description update chat by ID
// @ID update-chat
// @Accept json
// @Produce json
// @Param id path int true "Chat ID"
// @Param chat body ent.Chat true "Chat entity"
// @Success 200 {object} ent.Chat
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /chats/{id} [put]
func (ctl *ChatController) UpdateChat(c *gin.Context) {
	obj := Chat{}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "chat binding failed",
		})
		return
	}

	UpdateChat, err := ctl.client.Chat.
		UpdateOneID(int(id)).
		SetChatRoomName(obj.ChatRoomName).
		//AddChatUser(getUser).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   UpdateChat,
	})
}

// NewChatController creates and registers handles for the chat controller
func NewChatController(router gin.IRouter, client *ent.Client) *ChatController {
	uc := &ChatController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitChatController registers routes to the main engine
func (ctl *ChatController) register() {
	chats := ctl.router.Group("/chats")
	chats.GET(":id", ctl.ListChat)
	// CRUD
	chats.POST("", ctl.CreateChat)
	chats.GET("/chatroom/:id", ctl.GetChat)
	chats.PUT(":id", ctl.UpdateChat)
	chats.DELETE(":id", ctl.DeleteChat)
}
