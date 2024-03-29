package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
)

// FindUser search of user data by user ID.
func (db *MongoDB) FindUser(c echo.Context) error {
	id := c.Param("id")
	u := new(model.User)

	if err := db.UCol.FindId(bson.ObjectIdHex(id)).One(&u); err != nil {
		return c.JSON(http.StatusNotFound, "the user not found.")
	}

	return c.JSON(http.StatusOK, &u)
}

// EditUser edit user personal information.
func (db *MongoDB) EditUser(c echo.Context) error {
	id := c.Param("id")
	ud := new(model.UserData)
	if err := c.Bind(ud); err != nil {
		return err
	}

	u := &model.User{
		Name:     ud.Name,
		Email:    ud.Email,
		Gender:   ud.Gender,
		Hospital: ud.Hospital,
		Weight:   ud.Weight,
		Height:   ud.Height,
	}

	// Parse birthdate
	if ud.BirthDate != "" {
		bd, err := ParseDate(ud.BirthDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid date format.")
		}

		u.BirthDate = bd
	}

	d := bson.M{
		"$set": &u,
	}

	if err := db.UCol.UpdateId(bson.ObjectIdHex(id), &d); err != nil {
		return c.JSON(http.StatusBadRequest, "unable to edit user.")
	}

	return c.JSON(http.StatusOK, "user has been edited.")
}

// DeleteUser delete user information from the database.
func (db *MongoDB) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	// Remove the user in database
	if err := db.UCol.RemoveId(bson.ObjectIdHex(id)); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot delete user.")
	}

	// Remove the user in bloodPressures collection
	if err := db.BPCol.Remove(bson.M{"uid": bson.ObjectIdHex(id)}); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot delete user of bloodPressures collection.")
	}

	// Remove the user in glomerularInfils collection
	if err := db.GIRCol.Remove(bson.M{"uid": bson.ObjectIdHex(id)}); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot delete user of glomerularInfils collection.")
	}

	// Remove the user in bloodSugars collection
	if err := db.BSCol.Remove(bson.M{"uid": bson.ObjectIdHex(id)}); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot delete user of bloodSugars collection.")
	}

	// Remove the user in waterLists collection
	if err := db.WATERCol.Remove(bson.M{"uid": bson.ObjectIdHex(id)}); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot delete user of waterLists collection.")
	}

	// Remove the user in bmiLists collection
	if err := db.BMICol.Remove(bson.M{"uid": bson.ObjectIdHex(id)}); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot delete user of bmiLists collection.")
	}

	return c.JSON(http.StatusOK, "the user has been deleted.")
}

// CheckEmail search for emails that have been registered.
func (db *MongoDB) CheckEmail(u *model.User) string {
	q := bson.M{
		"email": u.Email,
	}
	if err := db.UCol.Find(q).One(&u); err != nil {
		return "not exist"
	}

	return "exist"
}

// CreateUser created a user in database.
func (db *MongoDB) CreateUser(u *model.User) {
	if err := db.UCol.Insert(&u); err != nil {
		return
	}
}
