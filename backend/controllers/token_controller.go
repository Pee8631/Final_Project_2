package controllers

import (
	"FinalProject/ent"
	"FinalProject/ent/token"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Token defines the struct for the token
type Token struct {
	AuthToken   string
	GeneratedAt string
	ExpiresAt   string
	User        int
}

// TokenController defines the struct for the token controller
type TokenController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateToken handles POST requests for adding token entities
// @Summary Create token
// @Description Create token
// @ID create-token
// @Accept json
// @Produce json
// @Param token body ent.Token true "Token entity"
// @Success 200 {object} ent.Token
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tokens [post]
func (ctl *TokenController) CreateToken(c *gin.Context) {
	obj := Token{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "token binding failed",
		})
		return
	}

	generatedAt, err := time.Parse(time.RFC3339, obj.GeneratedAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	expiresAt, err := time.Parse(time.RFC3339, obj.ExpiresAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	insertToken, err := ctl.client.Token.
		Create().
		SetAuthToken(obj.AuthToken).
		SetGeneratedAt(generatedAt).
		SetExpiresAt(expiresAt).
		SetAuthenticationTokenID(obj.User).
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
		"data":   insertToken,
	})
}

func generateToken(id int) (map[string]interface{}, error) {
	randomToken := make([]byte, 32)

	_, err := rand.Read(randomToken)

	if err != nil {
		return nil, err
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)
	//const timeLayout = "2006-01-02T15:04:05Z"

	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 60)

	generatedAt := dt.Format("2006-01-02" + "T" + "15:04:05" + "Z")
	expiresAt := expirtyTime.Format("2006-01-02" + "T" + "15:04:05" + "Z")

	values := map[string]interface{}{
		"AuthToken":   authToken,
		"GeneratedAt": generatedAt,
		"ExpiresAt":   expiresAt,
		"User":        id,
	}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	_, err = http.Post("http://localhost:8080/api/v1/tokens", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return nil, err
	}

	/*tokenDetails := map[string]interface{}{
		"token_type":   "Bearer",
		"auth_token":   authToken,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}*/

	return values, nil
}

// GetToken handles GET requests to retrieve a token entity
// @Summary Get a token entity by ID
// @Description get token by ID
// @ID get-token
// @Produce json
// @Param id path int true "Token ID"
// @Success 200 {object} ent.Token
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tokens/{id} [get]
func (ctl *TokenController) GetToken(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	getToken, err := ctl.client.Token.
		Query().
		Where(token.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, getToken)
}

/*
// testResourceHandler handles GET requests to retrieve a token entity
// @Summary Get a token entity by ID
// @Description get token by ID
// @ID get-token
// @Produce json
// @Param id path int true "Token ID"
// @Success 200 {object} ent.Token
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tokens/{authtoken} [get]
func (ctl *TokenController) TestResourceHandler(c *gin.Context) {
	authToken := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]

	userDetails, err := validateToken(authToken)

	if err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {

		username := fmt.Sprint(userDetails["username"])

		c.JSON(200, username)
	}
}

func validateToken(authToken string) (map[string]interface{}, error) {

	queryString := `SELECT
				token.id_token,
                authtoken,
                generatedat,
                expiresat
            FROM token
            LEFT JOIN project
            ON token.id_user = user.id_user
            WHERE authtoken = ?`
	IDtoken := 0
	authtoken := ""
	generatedAt := ""
	expiresAt := ""

	err := stmt.QueryRow(authToken).Scan(&IDtoken, &authtoken, &generatedAt, &expiresAt)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid access token.\r\n")
		}

		return nil, err
	}

	const timeLayout = "2006-01-02 15:04:05"

	expiryTime, _ := time.Parse(timeLayout, expiresAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return nil, errors.New("The token is expired.\r\n")
	}

	userDetails := map[string]interface{}{
		"user_id":      IDtoken,
		"username":     authtoken,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}

	return userDetails, nil

}*/

// ListToken handles request to get a list of token entities
// @Summary List token entities
// @Description list token entities
// @ID list-token
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Token
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tokens [get]
func (ctl *TokenController) ListToken(c *gin.Context) {
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

	listToken, err := ctl.client.Token.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return // proper error handling instead of panic in your app
	}

	c.JSON(200, listToken)
}

// DeleteToken handles DELETE requests to delete a token entity
// @Summary Delete a token entity by ID
// @Description get token by ID
// @ID delete-token
// @Produce json
// @Param id path int true "Token ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tokens/{id} [delete]
func (ctl *TokenController) DeleteToken(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Token.
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

// UpdateToken handles PUT requests to update a token entity
// @Summary Update a token entity by ID
// @Description update token by ID
// @ID update-token
// @Accept json
// @Produce json
// @Param id path int true "Token ID"
// @Param token body ent.Token true "Token entity"
// @Success 200 {object} ent.Token
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /tokens/{id} [put]
func (ctl *TokenController) UpdateToken(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := Token{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "token binding failed",
		})
		return
	}

	generatedAt, err := time.Parse(time.RFC3339, obj.GeneratedAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	expiresAt, err := time.Parse(time.RFC3339, obj.ExpiresAt)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	updateToken, err := ctl.client.Token.
		UpdateOneID(int(id)).
		SetAuthToken(obj.AuthToken).
		SetGeneratedAt(generatedAt).
		SetExpiresAt(expiresAt).
		SetAuthenticationTokenID(obj.User).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   updateToken,
	})
}

// NewTokenController creates and registers handles for the token controller
func NewTokenController(router gin.IRouter, client *ent.Client) *TokenController {
	uc := &TokenController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTokenController registers routes to the main engine
func (ctl *TokenController) register() {
	tokens := ctl.router.Group("/tokens")
	tokens.GET("", ctl.ListToken)
	// CRUD
	tokens.POST("", ctl.CreateToken)
	tokens.GET(":id", ctl.GetToken)
	//tokens.GET(":authtoken", ctl.TestResourceHandler)
	tokens.PUT(":id", ctl.UpdateToken)
	tokens.DELETE(":id", ctl.DeleteToken)
}
