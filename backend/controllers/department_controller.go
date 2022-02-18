package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/department"
	"context"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Department defines the struct for the department
type Department struct {
	Name string
}

// DepartmentController defines the struct for the department controller
type DepartmentController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDepartment handles POST requests for adding department entities
// @Summary Create department
// @Description Create department
// @ID create-department
// @Accept json
// @Produce json
// @Param department body ent.Department true "Department entity"
// @Success 200 {object} ent.Department
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /departments [post]
func (ctl *DepartmentController) CreateDepartment(c *gin.Context) {
	obj := Department{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "department binding failed",
		})
		return
	}

	insertDepartment, err := ctl.client.Department.
		Create().
		SetName(obj.Name).
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
		"data":   insertDepartment,
	})
}

// GetDepartment handles GET requests to retrieve a department entity
// @Summary Get a department entity by ID
// @Description get department by ID
// @ID get-department
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} ent.Department
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /departments/{id} [get]
func (ctl *DepartmentController) GetDepartment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getDepartment, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getDepartment)
}

// ListDepartment handles request to get a list of department entities
// @Summary List department entities
// @Description list department entities
// @ID list-department
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Department
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /departments [get]
func (ctl *DepartmentController) ListDepartment(c *gin.Context) {
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
	listDepartment, err := ctl.client.Department.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listDepartment)

}

// DeleteDepartment handles DELETE requests to delete a department entity
// @Summary Delete a department entity by ID
// @Description get department by ID
// @ID delete-department
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /departments/{id} [delete]
func (ctl *DepartmentController) DeleteDepartment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Department.
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

// UpdateDepartment handles PUT requests to update a department entity
// @Summary Update a department entity by ID
// @Description update department by ID
// @ID update-department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param department body ent.Department true "Department entity"
// @Success 200 {object} ent.Department
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /departments/{id} [put]
func (ctl *DepartmentController) UpdateDepartment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Department{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "department binding failed",
		})
		return
	}

	updateDepartment, err := ctl.client.Department.
		UpdateOneID(int(id)).
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateDepartment,
	})
}

// NewDepartmentController creates and registers handles for the department controller
func NewDepartmentController(router gin.IRouter, client *ent.Client) *DepartmentController {
	uc := &DepartmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDepartmentController registers routes to the main engine
func (ctl *DepartmentController) register() {
	departments := ctl.router.Group("/departments")
	departments.GET("", ctl.ListDepartment)
	// CRUD
	departments.POST("", ctl.CreateDepartment)
	departments.GET(":id", ctl.GetDepartment)
	departments.PUT(":id", ctl.UpdateDepartment)
	departments.DELETE(":id", ctl.DeleteDepartment)
}
