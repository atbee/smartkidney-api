package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
)

// ViewBP show logs of blood pressure for each user.
// You can view a list of blood pressure each date by query parameter.
// ?date=2019-10-07 or ?start=2019-10-07&&end=2019-10-10
func (db *MongoDB) ViewBP(c echo.Context) error {
	id := c.Param("id")
	bp := []*model.BloodPressureLog{}

	match := bson.M{
		"$match": bson.M{
			"uid": bson.ObjectIdHex(id),
		},
	}

	unwind := bson.M{
		"$unwind": "$logs",
	}

	replaceRoot := bson.M{
		"$replaceRoot": bson.M{
			"newRoot": "$logs",
		},
	}

	sort := bson.M{
		"$sort": bson.M{
			"date": 1,
		},
	}

	q := []bson.M{match, unwind, replaceRoot, sort}

	d := c.QueryParam("date")
	s := c.QueryParam("start")
	e := c.QueryParam("end")

	if d != "" || (s != "" && e != "") {
		st := StartDate(d)
		et := EndDate(d)

		if d == "" {
			st = StartDate(s)
			et = EndDate(e)
		}

		matchDate := bson.M{
			"$match": bson.M{
				"date": bson.M{
					"$gte": st,
					"$lt":  et,
				},
			},
		}

		q = []bson.M{match, unwind, replaceRoot, matchDate, sort}
	}

	if err := db.BPCol.Pipe(q).All(&bp); err != nil {
		return c.JSON(http.StatusNotFound, "the user not found.")
	}

	if len(bp) == 0 {
		return c.JSON(http.StatusNotFound, "user data not found, please check that the request form is correct.")
	}

	return c.JSON(http.StatusOK, &bp)
}

// AddBP add items systolic and diastolic to that user.
func (db *MongoDB) AddBP(c echo.Context) error {
	id := c.Param("id")
	bp := &model.BloodPressureLog{
		Date: time.Now(),
	}
	if err := c.Bind(bp); err != nil {
		return err
	}

	q := bson.M{
		"uid": bson.ObjectIdHex(id),
	}

	d := bson.M{
		"$push": bson.M{
			"logs": bp,
		},
	}

	if err := db.BPCol.Update(q, &d); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &bp)
}

// CreateBP created logs of blood pressure in database.
func (db *MongoDB) CreateBP(uid bson.ObjectId) {
	bp := &model.BloodPressure{
		ID:  bson.NewObjectId(),
		UID: uid,
	}
	if err := db.BPCol.Insert(&bp); err != nil {
		return
	}
}
