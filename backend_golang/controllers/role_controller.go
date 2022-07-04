package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/role"

	"context"

	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Role defines the struct for the role
type Role struct {
	Name string
}

// RoleController defines the struct for the role controller
type RoleController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRole handles POST requests for adding role entities
// @Summary Create role
// @Description Create role
// @ID create-role
// @Accept json
// @Produce json
// @Param role body ent.Role true "Role entity"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles [post]
func (ctl *RoleController) CreateRole(c *gin.Context) {
	obj := Role{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "role binding failed",
		})
		return
	}

	insertRole, err := ctl.client.Role.
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
		"data":   insertRole,
	})
}

// GetRole handles GET requests to retrieve a role entity
// @Summary Get a role entity by ID
// @Description get role by ID
// @ID get-role
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/{id} [get]
func (ctl *RoleController) GetRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	getRole, err := ctl.client.Role.
		Query().
		Where(role.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getRole)

}

// ListRole handles request to get a list of role entities
// @Summary List role entities
// @Description list role entities
// @ID list-role
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles [get]
func (ctl *RoleController) ListRole(c *gin.Context) {
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

	listRole, err := ctl.client.Role.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listRole)

}

// DeleteRole handles DELETE requests to delete a role entity
// @Summary Delete a role entity by ID
// @Description get role by ID
// @ID delete-role
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/{id} [delete]
func (ctl *RoleController) DeleteRole(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Role.
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

// UpdateRole handles PUT requests to update a role entity
// @Summary Update a role entity by ID
// @Description update role by ID
// @ID update-role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body ent.Role true "Role entity"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/{id} [put]
func (ctl *RoleController) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Role{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "role binding failed",
		})
		return
	}

	updateRole, err := ctl.client.Role.
		UpdateOneID(int(id)).
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateRole,
	})
}

// NewRoleController creates and registers handles for the role controller
func NewRoleController(router gin.IRouter, client *ent.Client) *RoleController {
	uc := &RoleController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRoleController registers routes to the main engine
func (ctl *RoleController) register() {
	roles := ctl.router.Group("/roles")
	roles.GET("", ctl.ListRole)
	// CRUD
	roles.POST("", ctl.CreateRole)
	roles.GET(":id", ctl.GetRole)
	roles.PUT(":id", ctl.UpdateRole)
	roles.DELETE(":id", ctl.DeleteRole)
}
