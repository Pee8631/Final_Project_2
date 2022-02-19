package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/user"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//Function determine database connection
/*func databaseConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "project"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=True")
	if err != nil {
		panic(err.Error())
	}
	return db
}*/

// User defines the struct for the user
type User struct {
	Username   string
	Password   string
	Department int
	Hospital   int
}

// UserController defines the struct for the user controller
type UserController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateUser handles POST requests for adding user entities
// @Summary Create user
// @Description Create user
// @ID create-user
// @Accept json
// @Produce json
// @Param user body ent.User true "User entity"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (ctl *UserController) CreateUser(c *gin.Context) {
	obj := User{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}
	if obj.Department != 0 && obj.Hospital != 0 {

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(obj.Password), 14)

		insertUser, err := ctl.client.User.
			Create().
			SetUsername(obj.Username).
			SetPassword(string(hashedPassword)).
			SetHasDepartmentID(obj.Department).
			SetFromHospitalID(obj.Hospital).
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
			"data":   insertUser,
		})
	} else {

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(obj.Password), 14)

		insertUser, err := ctl.client.User.
			Create().
			SetUsername(obj.Username).
			SetPassword(string(hashedPassword)).
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
			"data":   insertUser,
		})
	}
}

// GetUser handles GET requests to retrieve a user entity
// @Summary Get a user entity by ID
// @Description get user by ID
// @ID get-user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [get]
func (ctl *UserController) GetUser(c *gin.Context) {
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
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getUser)

}

// FindUser handles POST requests to retrieve a user entity
// @Summary Get a user entity by Username
// @Description get user by Username
// @Username get-user
// @Produce json
// @Param username path int true "User Username"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{username} [post]
func (ctl *UserController) AuthUser(c *gin.Context) {
	obj := ent.User{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	getUser, err := ctl.client.User.
		Query().
		Where(user.UsernameEQ(obj.Username)).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Username",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(obj.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	token, err := generateToken(getUser.ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, token)
}

/*func CreateToken(username string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}*/

// ListUser handles request to get a list of user entities
// @Summary List user entities
// @Description list user entities
// @ID list-user
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (ctl *UserController) ListUser(c *gin.Context) {
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

	listUser, err := ctl.client.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listUser)

}

// DeleteUser handles DELETE requests to delete a user entity
// @Summary Delete a user entity by ID
// @Description get user by ID
// @ID delete-user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (ctl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.User.
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

// UpdateUser handles PUT requests to update a user entity
// @Summary Update a user entity by ID
// @Description update user by ID
// @ID update-user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body ent.User true "User entity"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [put]
func (ctl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := User{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(obj.Password), 14)

	updateUser, err := ctl.client.User.
		UpdateOneID(int(id)).
		SetUsername(obj.Username).
		SetPassword(string(hashedPassword)).
		SetHasDepartmentID(obj.Department).
		SetFromHospitalID(obj.Hospital).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateUser,
	})
}

// NewUserController creates and registers handles for the user controller
func NewUserController(router gin.IRouter, client *ent.Client) *UserController {
	uc := &UserController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitUserController registers routes to the main engine
func (ctl *UserController) register() {
	users := ctl.router.Group("/users")
	users.GET("", ctl.ListUser)
	// CRUD
	users.POST("", ctl.CreateUser)
	users.GET(":id", ctl.GetUser)
	users.POST(":username", ctl.AuthUser)
	users.PUT(":id", ctl.UpdateUser)
	users.DELETE(":id", ctl.DeleteUser)
}
