package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/appointment"
	"FinalProject/ent/chat"
	"FinalProject/ent/user"
	"context"
	"time"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Appointment defines the struct for the appointment
type Appointment struct {
	ReasonForAppointment string
	Detail               string
	Status               string
	StartTime            string
	EndTime              string
	DoctorId             int
	UserId               int
	ScheduleId           int
}

// AppointmentController defines the struct for the appointment controller
type AppointmentController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateAppointment handles POST requests for adding appointment entities
// @Summary Create appointment
// @Description Create appointment
// @ID create-appointment
// @Accept json
// @Produce json
// @Param appointment body ent.Appointment true "Appointment entity"
// @Success 200 {object} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments [post]
func (ctl *AppointmentController) CreateAppointment(c *gin.Context) {
	obj := Appointment{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "appointment binding failed",
		})
		return
	}

	StartTime, err := time.Parse(time.RFC3339, obj.StartTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	EndTime, err := time.Parse(time.RFC3339, obj.EndTime)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertAppointment, err := ctl.client.Appointment.
		Create().
		SetReasonForAppointment(obj.ReasonForAppointment).
		SetDetail(obj.Detail).
		SetStartTime(StartTime).
		SetEndTime(EndTime).
		SetStatus(obj.Status).
		SetDoctorId(obj.DoctorId).
		SetUserId(obj.UserId).
		SetAppointmentScheduleID(obj.ScheduleId).
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
		"data":   insertAppointment,
	})
}

// GetAppointment handles GET requests to retrieve a appointment entity
// @Summary Get a appointment entity by ID
// @Description get appointment by ID
// @ID get-appointment
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments/{id} [get]
func (ctl *AppointmentController) GetAppointment(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getAppointment, err := ctl.client.Appointment.
		Query().
		Where(appointment.IDEQ(int(id))).
		WithAppointmentChat().
		Only(context.Background())
		//user.HasUserAppointmentWith(predicate.Appointment(user.IDEQ(int(id))))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getAppointment)
}

// ListAppointment handles request to get a list of appointment entities
// @Summary List appointment entities
// @Description list appointment entities
// @ID list-appointment
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments/list/{id} [get]
func (ctl *AppointmentController) ListAppointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	getUser, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(id))).
		WithUserHaveRole().
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}
	if getUser.Edges.UserHaveRole[0].ID == 2 {
		listAppointment, err := ctl.client.Appointment.
			Query().
			Where(appointment.DoctorIdEQ(int(id))).
			WithAppointmentChat().
			// Limit(limit).
			// Offset(offset).
			All(context.Background())

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}

		c.JSON(200, listAppointment)
	} else if getUser.Edges.UserHaveRole[0].ID == 3 {
		listAppointment, err := ctl.client.Appointment.
			Query().
			Where(appointment.UserIdEQ(int(id))).
			WithAppointmentChat().
			// Limit(limit).
			// Offset(offset).
			All(context.Background())

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}

		c.JSON(200, listAppointment)
	} else {
		c.JSON(400, gin.H{"error": "Not Found User in Role Doctor and User"})
		return // proper error handling instead of panic in your app
	}

}

// DeleteAppointment handles DELETE requests to delete a appointment entity
// @Summary Delete a appointment entity by ID
// @Description get appointment by ID
// @ID delete-appointment
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments/{id} [delete]
func (ctl *AppointmentController) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Appointment.DeleteOneID(int(id)).Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateAppointment handles PUT requests to update a appointment entity
// @Summary Update a appointment entity by ID
// @Description update appointment by ID
// @ID update-appointment
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param appointment body ent.Appointment true "Appointment entity"
// @Success 200 {object} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments/{id} [put]
func (ctl *AppointmentController) UpdateAppointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	Month := [12]string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}

	obj := Appointment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "appointment binding failed",
		})
		return
	}

	getAppointment, err := ctl.client.Appointment.
		Query().
		Where(appointment.IDEQ(int(id))).
		WithAppointmentSchedule().
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "Not Found Appointment"})
		return
	}

	if obj.Status == "Waiting" {

		StartTime, err := time.Parse(time.RFC3339, obj.StartTime)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}

		EndTime, err := time.Parse(time.RFC3339, obj.EndTime)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}
		updateAppointment, err := ctl.client.Appointment.
			UpdateOneID(int(id)).
			SetReasonForAppointment(obj.ReasonForAppointment).
			SetDetail(obj.Detail).
			SetStartTime(StartTime).
			SetEndTime(EndTime).
			SetStatus(obj.Status).
			SetDoctorId(getAppointment.DoctorId).
			SetUserId(obj.UserId).
			SetAppointmentScheduleID(getAppointment.Edges.AppointmentSchedule.ID).
			Save(context.Background())

		if err != nil {
			c.JSON(400, gin.H{"error": "update failed"})
			return
		}

		getUser, err := ctl.client.User.
			Query().
			Where(user.IDEQ(int(obj.UserId))).
			WithUserHasPInfo().
			Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{"error": "Not Found User"})
			return
		}
		dmy := updateAppointment.StartTime
		shr, smin, _ := updateAppointment.StartTime.Clock()
		ehr, emin, _ := updateAppointment.EndTime.Clock()
		monththai := Month[dmy.Month()]
		yearthai := strconv.Itoa(dmy.Year() + 543)
		paddings := ""
		paddinge := ""
		if smin <= 9 {
			paddings = "0"
		}
		if emin <= 9 {
			paddinge = "0"
		}
		var Message string = ""
		if len(getUser.Edges.UserHasPInfo) > 0 {
			Message = getUser.Edges.UserHasPInfo[0].FirstName +
				" " + getUser.Edges.UserHasPInfo[0].LastName +
				" ได้ส่งคำขอเข้ารับการปรึกษากับคุณ ในวันที่ " + strconv.Itoa(dmy.Day()) +
				" เดือน" + monththai +
				" ปี " + yearthai +
				" ตั้งแต่ " + strconv.Itoa(int(shr)) + ":" + paddings + strconv.Itoa(int(smin)) +
				" น. จนถึง " + strconv.Itoa(int(ehr)) + ":" + paddinge + strconv.Itoa(int(emin)) +
				" น. คุณยืนยันที่จะให้คำปรึกษาหรือไม่"
		} else {
			Message = getUser.Username +
				" ได้ส่งคำขอเข้ารับการปรึกษากับคุณ ในวันที่ " + strconv.Itoa(dmy.Day()) +
				" เดือน" + monththai +
				" ปี " + yearthai +
				" ตั้งแต่ " + strconv.Itoa(int(shr)) + ":" + paddings + strconv.Itoa(int(smin)) +
				" น. จนถึง " + strconv.Itoa(int(ehr)) + ":" + paddinge + strconv.Itoa(int(emin)) +
				" น. คุณยืนยันที่จะให้คำปรึกษาหรือไม่"
		}
		CreatedDate := time.Now()

		insertNotification, err := ctl.client.Notification.
			Create().
			SetSenderId(updateAppointment.UserId).
			SetRecipientId(updateAppointment.DoctorId).
			SetMessage(Message).
			SetCreatedDate(CreatedDate).
			SetAppointmentId(updateAppointment.ID).
			AddUserNotificationIDs(updateAppointment.DoctorId).
			AddUserNotificationIDs(updateAppointment.UserId).
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
			"status":       true,
			"Notification": insertNotification,
			"Appointment":  updateAppointment,
		})
	} else if obj.Status == "Confirm" {

		getUser, err := ctl.client.User.
			Query().
			Where(user.IDEQ(int(getAppointment.UserId))).
			WithUserHaveRole().
			WithWhoIsInThisChat().
			Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		getDoctor, err := ctl.client.User.
			Query().
			Where(user.IDEQ(int(getAppointment.DoctorId))).
			WithUserHaveRole().
			Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		var ChatId int = 0
		for i := range getUser.Edges.WhoIsInThisChat {
			getChat, _ := ctl.client.Chat.
				Query().
				Where(chat.IDEQ(int(getUser.Edges.WhoIsInThisChat[i].ID))).
				WithChatUser().
				Only(context.Background())
			if getChat != nil {
				for i := range getChat.Edges.ChatUser {
					if getChat.Edges.ChatUser[i].ID == getAppointment.DoctorId {
						ChatId = getChat.ID
					}
				}
			}
		}

		if ChatId == 0 {
			_, err := ctl.client.Chat.
				Create().
				SetChatRoomName(getUser.Username + "-" + getDoctor.Username).
				AddChatUserIDs(getAppointment.DoctorId).
				AddChatUserIDs(getAppointment.UserId).
				Save(context.Background())

			if err != nil {
				fmt.Println(err)
				c.JSON(400, gin.H{
					"status": false,
					"error":  err,
				})
				return
			}

		}

		updateAppointment, err := ctl.client.Appointment.
			UpdateOneID(int(id)).
			SetStatus(obj.Status).
			SetAppointmentChatID(ChatId).
			Save(context.Background())

		if err != nil {
			c.JSON(400, gin.H{"error": "update failed"})
			return
		}

		dmy := updateAppointment.StartTime
		shr, smin, _ := updateAppointment.StartTime.Clock()
		ehr, emin, _ := updateAppointment.EndTime.Clock()
		monththai := Month[dmy.Month()]
		yearthai := strconv.Itoa(dmy.Year() + 543)
		paddings := ""
		paddinge := ""
		if smin <= 9 {
			paddings = "0"
		}
		if emin <= 9 {
			paddinge = "0"
		}
		var Message string = ""
		if len(getDoctor.Edges.UserHasPInfo) > 0 {
			Message = getDoctor.Edges.UserHasPInfo[0].FirstName +
				" " + getDoctor.Edges.UserHasPInfo[0].LastName +
				" ได้ตอบรับคำขอของคุณแล้ว คุณสามารถเข้าไปรับคำปรึกษากับคุณหมอได้ในวันที่ " +
				strconv.Itoa(dmy.Day()) + " เดือน" + monththai + " ปี " + yearthai +
				" ตั้งแต่ " + strconv.Itoa(int(shr)) + ":" + paddings + strconv.Itoa(int(smin)) +
				" น. จนถึง " + strconv.Itoa(int(ehr)) + ":" + paddinge + strconv.Itoa(int(emin)) +
				" น."
		} else {
			Message = getDoctor.Username +
				" ได้ตอบรับคำขอของคุณแล้ว คุณสามารถเข้าไปรับคำปรึกษากับคุณหมอได้ในวันที่ " +
				strconv.Itoa(dmy.Day()) + " เดือน" + monththai + " ปี " + yearthai +
				" ตั้งแต่ " + strconv.Itoa(int(shr)) + ":" + paddings + strconv.Itoa(int(smin)) +
				" น. จนถึง " + strconv.Itoa(int(ehr)) + ":" + paddinge + strconv.Itoa(int(emin)) +
				" น."
		}
		CreatedDate := time.Now()

		insertNotification, err := ctl.client.Notification.
			Create().
			SetSenderId(updateAppointment.DoctorId).
			SetRecipientId(updateAppointment.UserId).
			SetMessage(Message).
			SetCreatedDate(CreatedDate).
			SetAppointmentId(updateAppointment.ID).
			AddUserNotificationIDs(updateAppointment.DoctorId).
			AddUserNotificationIDs(updateAppointment.UserId).
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
			"status":       true,
			"Appointment":  updateAppointment,
			"Notification": insertNotification,
		})
	} else if obj.Status == "Reject" {
		getDoctor, err := ctl.client.User.
			Query().
			Where(user.IDEQ(int(getAppointment.DoctorId))).
			WithUserHaveRole().
			Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		updateAppointment, err := ctl.client.Appointment.
			UpdateOneID(int(id)).
			SetStatus(obj.Status).
			Save(context.Background())

		if err != nil {
			c.JSON(400, gin.H{"error": "update failed"})
			return
		}

		var Message string = ""
		if len(getDoctor.Edges.UserHasPInfo) > 0 {
			Message = getDoctor.Edges.UserHasPInfo[0].FirstName +
				" " + getDoctor.Edges.UserHasPInfo[0].LastName +
				" ได้ปฏิเสธคำขอของคุณ กรุณานัดหมายใหม่อีกครั้ง"
		} else {
			Message = getDoctor.Username +
				" ได้ปฏิเสธคำขอของคุณ คุณสามารถเข้าไปรับคำปรึกษากับคุณหมอได้ในวันที่"
		}
		CreatedDate := time.Now()

		insertNotification, err := ctl.client.Notification.
			Create().
			SetSenderId(updateAppointment.DoctorId).
			SetRecipientId(updateAppointment.UserId).
			SetMessage(Message).
			SetCreatedDate(CreatedDate).
			SetAppointmentId(updateAppointment.ID).
			AddUserNotificationIDs(updateAppointment.DoctorId).
			AddUserNotificationIDs(updateAppointment.UserId).
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
			"status":      true,
			"Appointment": updateAppointment,
			"Notification": insertNotification,
		})
	} else {

		StartTime, err := time.Parse(time.RFC3339, obj.StartTime)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}

		EndTime, err := time.Parse(time.RFC3339, obj.EndTime)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err,
			})
			return
		}

		updateAppointment, err := ctl.client.Appointment.
			UpdateOneID(int(id)).
			SetReasonForAppointment(obj.ReasonForAppointment).
			SetDetail(obj.Detail).
			SetStartTime(StartTime).
			SetEndTime(EndTime).
			SetStatus(obj.Status).
			SetDoctorId(getAppointment.DoctorId).
			SetUserId(obj.UserId).
			SetAppointmentScheduleID(getAppointment.Edges.AppointmentSchedule.ID).
			Save(context.Background())

		if err != nil {
			c.JSON(400, gin.H{"error": "update failed"})
			return
		}

		c.JSON(200, gin.H{
			"status": true,
			"data":   updateAppointment,
		})
	}

}

// NewAppointmentController creates and registers handles for the appointment controller
func NewAppointmentController(router gin.IRouter, client *ent.Client) *AppointmentController {
	uc := &AppointmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitAppointmentController registers routes to the main engine
func (ctl *AppointmentController) register() {
	appointments := ctl.router.Group("/appointments")
	appointments.GET("/list/:id", ctl.ListAppointment)
	// CRUD
	appointments.POST("", ctl.CreateAppointment)
	appointments.GET(":id", ctl.GetAppointment)
	appointments.PUT(":id", ctl.UpdateAppointment)
	appointments.DELETE(":id", ctl.DeleteAppointment)
}
