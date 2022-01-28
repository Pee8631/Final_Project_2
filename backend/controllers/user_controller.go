package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/user"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Function determine database connection
func databaseConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "project"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=True")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// User defines the struct for the user
type User struct {
	Username   	string
	Password   	string
	Email      	string
	Telephone  	string
	Department 	int
	Hospital	int
}

// UserController defines the struct for the user controller
type UserController struct {
	client *ent.Client
	router gin.IRouter
}

func (p User) GetIDDepartmentAsString() string {
	return strconv.Itoa(p.Department)
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
	db := databaseConnection()
	obj := User{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	insertUser, err := db.Prepare("INSERT INTO user(username, password, email, telephone, id_department, id_hospital) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}
	insertUser.Exec(obj.Username, obj.Password, obj.Email, obj.Telephone, obj.Department, obj.Hospital)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error": err,
		})
		return
	}

	defer db.Close()

	println("Username: " + obj.Username + ", Password: " + obj.Password + ", Email: " + obj.Telephone)

	
	c.JSON(200, gin.H{
		"status" : true,
		"data": insertUser,
	})
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
	db := databaseConnection()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	resultsGetUser, err := db.Query("SELECT username, password, email, telephone, id_department, id_hospital FROM user WHERE id_user=?", id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	for resultsGetUser.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = resultsGetUser.Scan(&user.Username, &user.Password, &user.Email, &user.Telephone, &user.Department, &user.Hospital)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer db.Close()

		c.JSON(200, user)

	}
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
	object := ent.User{}
	if err := c.ShouldBind(&object); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}
	user, err := ctl.client.User.
		Query().
		Where(user.UsernameEQ(object.Username)).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if  user.Password != object.Password {
		c.JSON(404, gin.H{"error": "invalid Password"})
		return
	}

	token, err := CreateToken(user.Username)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, token)

}

func CreateToken(username string) (string, error) {
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
}

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
	resultsListUser, err := db.Query("SELECT username, password, email, telephone, id_department, id_hospital FROM user limit ? offset ?", limit, offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	for resultsListUser.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = resultsListUser.Scan(&user.Username, &user.Password, &user.Email, &user.Telephone, &user.Department, &user.Hospital)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		defer db.Close()

		c.JSON(200, user)
	}

	
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
	db := databaseConnection()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteUser, err := db.Prepare("DELETE FROM user WHERE id_user=?")
    if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
    }
    deleteUser.Exec(id)

    defer db.Close()

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
	db := databaseConnection()
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

	UpdateUser, err := db.Prepare("UPDATE user SET username=?,password=?,email=?,telephone=?,id_department=?,id_hospital=? WHERE id_user=?")
	if err != nil {
		c.JSON(400, gin.H{"error": "update1 failed"})
		return
	}

	UpdateUser.Exec(obj.Username, obj.Password, obj.Email, obj.Telephone, obj.Department, obj.Hospital, int(id))

	defer db.Close()

	c.JSON(200, gin.H{
		"status" : true,
		"data": UpdateUser,
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
