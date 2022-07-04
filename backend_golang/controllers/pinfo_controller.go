package controllers

import (
	"FinalProject/ent"
	
	"FinalProject/ent/user"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
//"FinalProject/ent/pinfo"
// "FinalProject/ent/predicate"
// PInfo defines the struct for the pinfo
type PInfo struct {
	IdCardNumber string
	FirstName    string
	LastName     string
	Gender       int
	BrithDate    string
	BloodGroup   string
	Address      string
	User         int
}

// PInfoController defines the struct for the pinfo controller
type PInfoController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePInfo handles POST requests for adding pinfo entities
// @Summary Create pinfo
// @Description Create pinfo
// @ID create-pinfo
// @Accept json
// @Produce json
// @Param pinfo body ent.PInfo true "PInfo entity"
// @Success 200 {object} ent.PInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pinfos [post]
func (ctl *PInfoController) CreatePInfo(c *gin.Context) {
	obj := PInfo{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pinfo binding failed",
		})
		return
	}
	obj.BrithDate += "Z";
	brithDate, err := time.Parse(time.RFC3339, obj.BrithDate)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertPInfo, err := ctl.client.PInfo.
		Create().
		SetIdCardNumber(obj.IdCardNumber).
		SetFirstName(obj.FirstName).
		SetLastName(obj.LastName).
		SetGender(obj.Gender).
		SetBrithDate(brithDate).
		SetBloodGroup(obj.BloodGroup).
		SetAddress(obj.Address).
		SetWhoIsTheOwnerOfThisPInfoID(obj.User).
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
		"pinfo":   insertPInfo,
	})
}

// GetPInfo handles GET requests to retrieve a pinfo entity
// @Summary Get a pinfo entity by ID
// @Description get pinfo by ID
// @ID get-pinfo
// @Produce json
// @Param id path int true "PInfo ID"
// @Success 200 {object} ent.PInfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pinfos/{id} [get]
func (ctl *PInfoController) GetPInfo(c *gin.Context) {
	UserId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getUser, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(UserId))).
		WithUserHasPInfo().
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}
	if len(getUser.Edges.UserHasPInfo) > 0 {
		layout := "2006-01-02 15:04:05"
		values := map[string]interface{}{
			"Id":   		getUser.Edges.UserHasPInfo[0].ID,
			"IdCardNumber": getUser.Edges.UserHasPInfo[0].IdCardNumber,
			"FirstName":   	getUser.Edges.UserHasPInfo[0].FirstName,
			"LastName":     getUser.Edges.UserHasPInfo[0].LastName,
			"BrithDate":    getUser.Edges.UserHasPInfo[0].BrithDate.Format(layout),
			"Gender":     	getUser.Edges.UserHasPInfo[0].Gender,
			"BloodGroup":   getUser.Edges.UserHasPInfo[0].BloodGroup,
			"Address":      getUser.Edges.UserHasPInfo[0].Address,
			"User":     	getUser.ID,
		}

		c.JSON(200, values)

	}else {
		c.JSON(404, gin.H{
			"error": "PInfo Not Found",
		})
	}
	// PInfoId := getUser.Edges.UserHasPInfo[0].ID;
	// getPInfo, err := ctl.client.PInfo.
	// 	Query().Where(pinfo.IDEQ(int(PInfoId))).
	// 	Only(context.Background())
	// if err != nil {
	// 	c.JSON(404, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return // proper error handling instead of panic in your app
	// }

}

// ListPInfo handles request to get a list of pinfo entities
// @Summary List pinfo entities
// @Description list pinfo entities
// @ID list-pinfo
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pinfos [get]
func (ctl *PInfoController) ListPInfo(c *gin.Context) {
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


	listPInfo, err := ctl.client.PInfo.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

		c.JSON(200, listPInfo)

}

// DeletePInfo handles DELETE requests to delete a pinfo entity
// @Summary Delete a pinfo entity by ID
// @Description get pinfo by ID
// @ID delete-pinfo
// @Produce json
// @Param id path int true "PInfo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pinfos/{id} [delete]
func (ctl *PInfoController) DeletePInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PInfo.
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

// UpdatePInfo handles PUT requests to update a pinfo entity
// @Summary Update a pinfo entity by ID
// @Description update pinfo by ID
// @ID update-pinfo
// @Accept json
// @Produce json
// @Param id path int true "PInfo ID"
// @Param pinfo body ent.PInfo true "PInfo entity"
// @Success 200 {object} ent.PInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pinfos/{id} [put]
func (ctl *PInfoController) UpdatePInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := PInfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pinfo binding failed",
		})
		return
	}
	
	obj.BrithDate += "Z";
	brithDate, err := time.Parse(time.RFC3339, obj.BrithDate)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	UpdatePInfo, err := ctl.client.PInfo.
		UpdateOneID(int(id)).
		SetIdCardNumber(obj.IdCardNumber).
		SetFirstName(obj.FirstName).
		SetLastName(obj.LastName).
		SetGender(obj.Gender).
		SetBrithDate(brithDate).
		SetBloodGroup(obj.BloodGroup).
		SetAddress(obj.Address).
		SetWhoIsTheOwnerOfThisPInfoID(obj.User).
		Save(context.Background())
	if err != nil {
		c.JSON(404, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"pinfo":   UpdatePInfo,
	})
}

// NewPInfoController creates and registers handles for the pinfo controller
func NewPInfoController(router gin.IRouter, client *ent.Client) *PInfoController {
	uc := &PInfoController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitPInfoController registers routes to the main engine
func (ctl *PInfoController) register() {
	pinfos := ctl.router.Group("/pinfos")
	pinfos.GET("", ctl.ListPInfo)
	// CRUD
	pinfos.POST("", ctl.CreatePInfo)
	pinfos.GET(":id", ctl.GetPInfo)
	pinfos.PUT(":id", ctl.UpdatePInfo)
	pinfos.DELETE(":id", ctl.DeletePInfo)
}
