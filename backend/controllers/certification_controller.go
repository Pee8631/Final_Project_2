package controllers

import (
	"FinalProject/ent"
	"fmt"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Certification defines the struct for the certification
type Certification struct {
	Code			string
	Diloma			string
	DateOfIssuing	string
	DateOfExp		string
	Issuer			string
	User			int
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
	db := databaseConnection()
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
			"status" : false,
			"error": err,
		})
		return
	}

	dateOfExp, err := time.Parse(time.RFC3339, obj.DateOfExp)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	insertCertification, err := db.Prepare("INSERT INTO certification( code, diloma, dateofissuing, dateofexp, issuer, id_user) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertCertification.Exec(obj.Code, obj.Diloma, dateOfIssuing, dateOfExp, obj.Issuer, obj.User)
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
		"data": insertCertification,
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
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetCertification, err := db.Query("SELECT code, diloma, dateofissuing, dateofexp, issuer, id_user FROM certification WHERE id_certification=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetCertification.Next() {
		var certification Certification
		// for each row, scan the result into our tag composite object
		err = resultsGetCertification.Scan(&certification.Code, &certification.Diloma, &certification.DateOfIssuing, &certification.DateOfExp, &certification.Issuer, &certification.User)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, certification)

	}
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
	resultsListCertification, err := db.Query("SELECT code, diloma, dateofissuing, dateofexp, issuer, id_user FROM certification limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListCertification.Next() {
		var certification Certification
		// for each row, scan the result into our tag composite object
		err = resultsListCertification.Scan(&certification.Code, &certification.Diloma, &certification.DateOfIssuing, &certification.DateOfExp, &certification.Issuer, &certification.User)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, certification)
	}

	
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
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteCertification, err := db.Prepare("DELETE FROM certification WHERE id_certification=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteCertification.Exec(id)

    defer db.Close()

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
	db := databaseConnection()
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
			"status" : false,
			"error": err,
		})
		return
	}
	
	dateOfExp, err := time.Parse(time.RFC3339, obj.DateOfExp)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	UpdateCertification, err := db.Prepare("UPDATE certification SET code=?, diloma=?, dateofissuing=?, dateofexp=?, issuer=?, id_user=? WHERE id_certification=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateCertification.Exec(obj.Code, obj.Diloma, dateOfIssuing, dateOfExp, obj.Issuer, obj.User, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateCertification,
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
