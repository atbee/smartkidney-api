package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
)

func (db *MongoDB) ViewBMI(c echo.Context) error {
	id := c.Param("id")
	bmi := []*model.BMILog{}

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

	if err := db.BSCol.Pipe(q).All(&bmi); err != nil {
		return c.JSON(http.StatusNotFound, "the user not found.")
	}

	if len(bmi) == 0 {
		return c.JSON(http.StatusNotFound, "user data is empty or not found.")
	}

	return c.JSON(http.StatusOK, &bmi)
}

func (db *MongoDB) AddBMI(c echo.Context) error {
	id := c.Param("id")
	bmi := &model.BMILog{
		Date: time.Now(),
	}
	if err := c.Bind(bmi); err != nil {
		return err
	}

	q := bson.M{
		"uid": bson.ObjectIdHex(id),
	}

	d := bson.M{
		"$push": bson.M{
			"logs": bmi,
		},
	}

	if err := db.BSCol.Update(q, &d); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &bmi)
}

func (db *MongoDB) CreateBMI(uid bson.ObjectId) {
	bmi := &model.BMI{
		ID:  bson.NewObjectId(),
		UID: uid,
	}
	if err := db.BSCol.Insert(&bmi); err != nil {
		return
	}
}
