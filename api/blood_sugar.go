package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rocketblack/smartkidney-api/model"
	"gopkg.in/mgo.v2/bson"
)

func (db *MongoDB) ViewBS(c echo.Context) error {
	id := c.Param("id")
	bs := []*model.BloodSugarLog{}

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

	if err := db.BSCol.Pipe(q).All(&bs); err != nil {
		return c.JSON(http.StatusNotFound, "the user not found.")
	}

	if len(bs) == 0 {
		return c.JSON(http.StatusNotFound, "user data is empty or not found.")
	}

	return c.JSON(http.StatusOK, &bs)
}

func (db *MongoDB) AddBS(c echo.Context) error {
	id := c.Param("id")
	bs := &model.BloodSugarLog{
		Date: time.Now(),
	}
	if err := c.Bind(bs); err != nil {
		return err
	}

	q := bson.M{
		"uid": bson.ObjectIdHex(id),
	}

	d := bson.M{
		"$push": bson.M{
			"logs": bs,
		},
	}

	if err := db.BSCol.Update(q, &d); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &bs)
}

func (db *MongoDB) CreateBS(uid bson.ObjectId) {
	bs := &model.BloodSugar{
		ID:  bson.NewObjectId(),
		UID: uid,
	}
	if err := db.BSCol.Insert(&bs); err != nil {
		return
	}
}
