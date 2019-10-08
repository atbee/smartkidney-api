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
	d := new(model.UserData)

	if err := c.Bind(d); err != nil {
		return err
	}

	// Parse birthdate
	bd, err := ParseDate(d.BirthDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid date format.")
	}

	u := &model.User{
		ID:        bson.NewObjectId(),
		CreateAt:  time.Now(),
		Name:      d.Name,
		Email:     d.Email,
		BirthDate: bd,
		Gender:    d.Gender,
		Hospital:  d.Hospital,
	}

	// Validate
	if errs := validator.Validate(u); errs != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid request datas."}
	}

	// Save user
	db.CreateUser(u)

	// Init logs
	db.CreateBP(u.ID)
	db.CreateGIR(u.ID)
	db.CreateBS(u.ID)
	db.CreateBMI(u.ID)
	db.CreateWater(u.ID)

	return c.JSON(http.StatusCreated, u) // User has been registered.
}
