package storage

import (
	"github.com/jmoiron/sqlx"
)

var createSubLocationQuery string = `create table if not exists SubLocations
(
	locationID INTEGER,
	name TEXT,
	capacity INTEGER,
	occupied INTEGER,
	density INTEGER,
	seatingPlan TEXT
)`

var createLocationHistoryQuery string = `create table if not exists LocationHistories
(
	locationID INTEGER,
	timestamp TEXT,
	value INTEGER
)`

var createStudyPlaceQuery string = `create table if not exists StudyPlaces
(
	locationID INTEGER,
	name TEXT,
	capacity INTEGER,
	occupied INTEGER,
	density INTEGER,
	image TEXT,
	type TEXT,
	hasSeatingPlan INTEGER
)`

var createEventQuery string = `create table if not exists Events
(
	eventID INTEGER,
	name TEXT,
	location TEXT,
	eventType TEXT,
	crowd INTEGER,
	image TEXT
)`

func initDBSchema(db *sqlx.DB) error {
	_, err := db.Exec(createSubLocationQuery)
	_, err = db.Exec(createLocationHistoryQuery)
	_, err = db.Exec(createStudyPlaceQuery)
	_, err = db.Exec(createEventQuery)
	return err
}