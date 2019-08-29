package api

import (
	"net/http"
	"time"

	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

// Login using email from gmail or facebook.
func (db *MongoDB) Login(c echo.Context) (err error) {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	res := &model.LoginRes{
		FirstLogin: false,
	}

	q := bson.M{
		"email": u.Email,
	}

	if err := db.UCol.Find(&q).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return err
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

	// Find User

	// Validate
	if errs := validator.Validate(u); errs != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid data."}
	}

	if err := db.UCol.Insert(&u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &u)
}
