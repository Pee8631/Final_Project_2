package controllers

import (
	"FinalProject/ent"

	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Telecom defines the struct for the telecom
type Telecom struct {
	Username   	string
	Platform	string
	User		int
}

// TelecomController defines the struct for the telecom controller
type TelecomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTelecom handles POST requests for adding telecom entities
// @Summary Create telecom
// @Description Create telecom
// @ID create-telecom
// @Accept json
// @Produce json
// @Param telecom body ent.Telecom true "Telecom entity"
// @Success 200 {object} ent.Telecom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /telecoms [post]
func (ctl *TelecomController) CreateTelecom(c *gin.Context) {
	db := databaseConnection()
	obj := Telecom{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "telecom binding failed",
		})
		return
	}

	insertTelecom, err := db.Prepare("INSERT INTO telecom(username, platform, id_user) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertTelecom.Exec(obj.Username, obj.Platform, obj.User)
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
		"data": insertTelecom,
	})
}

// GetTelecom handles GET requests to retrieve a telecom entity
// @Summary Get a telecom entity by ID
// @Description get telecom by ID
// @ID get-telecom
// @Produce json
// @Param id path int true "Telecom ID"
// @Success 200 {object} ent.Telecom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /telecoms/{id} [get]
func (ctl *TelecomController) GetTelecom(c *gin.Context) {
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetTelecom, err := db.Query("SELECT username, platform, id_user FROM telecom WHERE id_telecom=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetTelecom.Next() {
		var telecom Telecom
		// for each row, scan the result into our tag composite object
		err = resultsGetTelecom.Scan(&telecom.Username, &telecom.Platform, &telecom.User)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, telecom)

	}
}

// ListTelecom handles request to get a list of telecom entities
// @Summary List telecom entities
// @Description list telecom entities
// @ID list-telecom
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Telecom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /telecoms [get]
func (ctl *TelecomController) ListTelecom(c *gin.Context) {
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
	resultsListTelecom, err := db.Query("SELECT username, platform, id_user FROM telecom limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListTelecom.Next() {
		var telecom Telecom
		// for each row, scan the result into our tag composite object
		err = resultsListTelecom.Scan(&telecom.Username, &telecom.Platform, &telecom.User)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, telecom)
	}

	
}

// DeleteTelecom handles DELETE requests to delete a telecom entity
// @Summary Delete a telecom entity by ID
// @Description get telecom by ID
// @ID delete-telecom
// @Produce json
// @Param id path int true "Telecom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /telecoms/{id} [delete]
func (ctl *TelecomController) DeleteTelecom(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteTelecom, err := db.Prepare("DELETE FROM telecom WHERE id_telecom=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteTelecom.Exec(id)

    defer db.Close()

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTelecom handles PUT requests to update a telecom entity
// @Summary Update a telecom entity by ID
// @Description update telecom by ID
// @ID update-telecom
// @Accept json
// @Produce json
// @Param id path int true "Telecom ID"
// @Param telecom body ent.Telecom true "Telecom entity"
// @Success 200 {object} ent.Telecom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /telecoms/{id} [put]
func (ctl *TelecomController) UpdateTelecom(c *gin.Context) {
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Telecom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "telecom binding failed",
		})
		return
	}

	UpdateTelecom, err := db.Prepare("UPDATE telecom SET username=?, platform=?, id_user=? WHERE id_telecom=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateTelecom.Exec(obj.Username, obj.Platform, obj.User, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateTelecom,
	})
}

// NewTelecomController creates and registers handles for the telecom controller
func NewTelecomController(router gin.IRouter, client *ent.Client) *TelecomController {
	uc := &TelecomController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTelecomController registers routes to the main engine
func (ctl *TelecomController) register() {
	telecoms := ctl.router.Group("/telecoms")
	telecoms.GET("", ctl.ListTelecom)
	// CRUD
	telecoms.POST("", ctl.CreateTelecom)
	telecoms.GET(":id", ctl.GetTelecom)
	telecoms.PUT(":id", ctl.UpdateTelecom)
	telecoms.DELETE(":id", ctl.DeleteTelecom)
}
