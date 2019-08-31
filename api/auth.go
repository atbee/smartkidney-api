package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

// Login using email from gmail or facebook.
func (db *MongoDB) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Check email address in database
	s := db.CheckEmail(u)
	if s == "not exist" {
		return c.JSON(http.StatusUnauthorized, map[string]bool{"firstLogin": true})
	}

	// Response user data
	res := &model.LoginRes{
		FirstLogin: false,
		User:       *u,
	}

	return c.JSON(http.StatusOK, &res)
}

// Register the user to database.
func (db *MongoDB) Register(c echo.Context) (err error) {
	u := &model.User{
		ID:       bson.NewObjectId(),
		CreateAt: time.Now(),
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	// Validate
	if errs := validator.Validate(u); errs != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid data."}
	}

	// Save user
	db.CreateUser(u)

	return c.JSON(http.StatusCreated, u) // User has been registered.
}
