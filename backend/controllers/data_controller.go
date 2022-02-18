package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/data"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Data defines the struct for the data
type Data struct {
	IdCardNumber string
	FirstName    string
	LastName     string
	Gender       int
	BrithDate    string
	BloodGroup   string
	Address      string
	User         int
}

// DataController defines the struct for the data controller
type DataController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateData handles POST requests for adding data entities
// @Summary Create data
// @Description Create data
// @ID create-data
// @Accept json
// @Produce json
// @Param data body ent.Data true "Data entity"
// @Success 200 {object} ent.Data
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datas [post]
func (ctl *DataController) CreateData(c *gin.Context) {
	obj := Data{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "data binding failed",
		})
		return
	}

	brithDate, err := time.Parse(time.RFC3339, obj.BrithDate)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertData, err := ctl.client.Data.
		Create().
		SetIdCardNumber(obj.IdCardNumber).
		SetFirstName(obj.FirstName).
		SetLastName(obj.LastName).
		SetGender(obj.Gender).
		SetBrithDate(brithDate).
		SetBloodGroup(obj.BloodGroup).
		SetWhoIsTheOwnerOfThisDataID(obj.User).
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
		"data":   insertData,
	})
}

// GetData handles GET requests to retrieve a data entity
// @Summary Get a data entity by ID
// @Description get data by ID
// @ID get-data
// @Produce json
// @Param id path int true "Data ID"
// @Success 200 {object} ent.Data
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datas/{id} [get]
func (ctl *DataController) GetData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getData, err := ctl.client.Data.
		Query().
		Where(data.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getData)

}

// ListData handles request to get a list of data entities
// @Summary List data entities
// @Description list data entities
// @ID list-data
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Data
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datas [get]
func (ctl *DataController) ListData(c *gin.Context) {
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


	listData, err := ctl.client.Data.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

		c.JSON(200, listData)

}

// DeleteData handles DELETE requests to delete a data entity
// @Summary Delete a data entity by ID
// @Description get data by ID
// @ID delete-data
// @Produce json
// @Param id path int true "Data ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datas/{id} [delete]
func (ctl *DataController) DeleteData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Data.
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

// UpdateData handles PUT requests to update a data entity
// @Summary Update a data entity by ID
// @Description update data by ID
// @ID update-data
// @Accept json
// @Produce json
// @Param id path int true "Data ID"
// @Param data body ent.Data true "Data entity"
// @Success 200 {object} ent.Data
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /datas/{id} [put]
func (ctl *DataController) UpdateData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Data{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "data binding failed",
		})
		return
	}

	brithDate, err := time.Parse(time.RFC3339, obj.BrithDate)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	UpdateData, err := ctl.client.Data.
		UpdateOneID(int(id)).
		SetIdCardNumber(obj.IdCardNumber).
		SetFirstName(obj.FirstName).
		SetLastName(obj.LastName).
		SetGender(obj.Gender).
		SetBrithDate(brithDate).
		SetBloodGroup(obj.BloodGroup).
		SetWhoIsTheOwnerOfThisDataID(obj.User).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   UpdateData,
	})
}

// NewDataController creates and registers handles for the data controller
func NewDataController(router gin.IRouter, client *ent.Client) *DataController {
	uc := &DataController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDataController registers routes to the main engine
func (ctl *DataController) register() {
	datas := ctl.router.Group("/datas")
	datas.GET("", ctl.ListData)
	// CRUD
	datas.POST("", ctl.CreateData)
	datas.GET(":id", ctl.GetData)
	datas.PUT(":id", ctl.UpdateData)
	datas.DELETE(":id", ctl.DeleteData)
}
