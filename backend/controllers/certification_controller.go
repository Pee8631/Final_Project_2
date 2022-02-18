package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/certification"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Certification defines the struct for the certification
type Certification struct {
	Code          string
	Diloma        string
	DateOfIssuing string
	DateOfExp     string
	Issuer        string
	User          int
}

// CertificationController defines the struct for the certification controller
type CertificationController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateCertification handles POST requests for adding certification entities
// @Summary Create certification
// @Description Create certification
// @ID create-certification
// @Accept json
// @Produce json
// @Param certification body ent.Certification true "Certification entity"
// @Success 200 {object} ent.Certification
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certifications [post]
func (ctl *CertificationController) CreateCertification(c *gin.Context) {
	obj := Certification{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "certification binding failed",
		})
		return
	}

	dateOfIssuing, err := time.Parse(time.RFC3339, obj.DateOfIssuing)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	dateOfExp, err := time.Parse(time.RFC3339, obj.DateOfExp)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertCertification, err := ctl.client.Certification.
		Create().
		SetCode(obj.Code).
		SetDiloma(obj.Diloma).
		SetDateOfIssuing(dateOfIssuing).
		SetDateOfExp(dateOfExp).
		SetIssuer(obj.Issuer).
		SetDoctorOwnerID(obj.User).
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
		"data":   insertCertification,
	})
}

// GetCertification handles GET requests to retrieve a certification entity
// @Summary Get a certification entity by ID
// @Description get certification by ID
// @ID get-certification
// @Produce json
// @Param id path int true "Certification ID"
// @Success 200 {object} ent.Certification
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certifications/{id} [get]
func (ctl *CertificationController) GetCertification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getCertification, err := ctl.client.Certification.
		Query().
		Where(certification.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getCertification)
}

// ListCertification handles request to get a list of certification entities
// @Summary List certification entities
// @Description list certification entities
// @ID list-certification
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Certification
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certifications [get]
func (ctl *CertificationController) ListCertification(c *gin.Context) {
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

	listCertification, err := ctl.client.Certification.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listCertification)

}

// DeleteCertification handles DELETE requests to delete a certification entity
// @Summary Delete a certification entity by ID
// @Description get certification by ID
// @ID delete-certification
// @Produce json
// @Param id path int true "Certification ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certifications/{id} [delete]
func (ctl *CertificationController) DeleteCertification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Certification.
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

// UpdateCertification handles PUT requests to update a certification entity
// @Summary Update a certification entity by ID
// @Description update certification by ID
// @ID update-certification
// @Accept json
// @Produce json
// @Param id path int true "Certification ID"
// @Param certification body ent.Certification true "Certification entity"
// @Success 200 {object} ent.Certification
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /certifications/{id} [put]
func (ctl *CertificationController) UpdateCertification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Certification{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "certification binding failed",
		})
		return
	}
	dateOfIssuing, err := time.Parse(time.RFC3339, obj.DateOfIssuing)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	dateOfExp, err := time.Parse(time.RFC3339, obj.DateOfExp)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	updateCertification, err := ctl.client.Certification.
		UpdateOneID(int(id)).
		SetCode(obj.Code).
		SetDiloma(obj.Diloma).
		SetDateOfIssuing(dateOfIssuing).
		SetDateOfExp(dateOfExp).
		SetIssuer(obj.Issuer).
		SetDoctorOwnerID(obj.User).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateCertification,
	})
}

// NewCertificationController creates and registers handles for the certification controller
func NewCertificationController(router gin.IRouter, client *ent.Client) *CertificationController {
	uc := &CertificationController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitCertificationController registers routes to the main engine
func (ctl *CertificationController) register() {
	certifications := ctl.router.Group("/certifications")
	certifications.GET("", ctl.ListCertification)
	// CRUD
	certifications.POST("", ctl.CreateCertification)
	certifications.GET(":id", ctl.GetCertification)
	certifications.PUT(":id", ctl.UpdateCertification)
	certifications.DELETE(":id", ctl.DeleteCertification)
}
