package storage

import (
	"../models"
	"database/sql"
)

var listAllStudyPlacesQuery string = "select * from StudyPlaces"
var getStudyPlaceByIdQuery string = "select * from StudyPlaces where locationID = ?"
var getSubLocationsByIdQuery string = "select name, capacity, occupied, density, seatingPlan from SubLocations where locationID = ?"
var getLocationHistoriesByIdQuery = "select timestamp, value from LocationHistories where locationID = ? ORDER BY timestamp ASC"


var listAllEventsQuery string = "select * from Events ORDER BY eventID DESC"
var getEventByIdQuery string = "select * from Events where eventID = ?"
var getLastEventIdQuery string = "select max(eventID) from Events"
var addEventQuery string = "INSERT INTO Events (eventID, name, location, eventType, crowd, image) VALUES (:eventID, :name, :location, :eventType, :crowd, :image)"

func (db *dbImpl) GetLastEventId() (int, error) {
	var lastId int
	db.sqliteDB.Get(&lastId, getLastEventIdQuery)
	return lastId, nil
}
func (db *dbImpl) AddEvent(event *models.Event) (error) {
	_, _ = db.sqliteDB.NamedExec(addEventQuery, event)
	return nil
}

func (db *dbImpl) GetStudyPlaceById(id int) (*models.Study_place, error) {
    result := models.Study_place{}
    err := db.sqliteDB.Get(&result, getStudyPlaceByIdQuery, id)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}

func (db *dbImpl) GetLocationHistoriesById(id int) ([]models.Location_history, error) {
    result := []models.Location_history{}
    err := db.sqliteDB.Select(&result, getLocationHistoriesByIdQuery, id)
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return result, err
}

func (db *dbImpl) GetSubLocationsById(id int) ([]models.Sub_location, error) {
    result := []models.Sub_location{}
    err := db.sqliteDB.Select(&result, getSubLocationsByIdQuery, id)
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return result, err
}

func (db *dbImpl) ListAllStudyPlaces() ([]models.Study_place, error) {
	res := []models.Study_place{}
	err := db.sqliteDB.Select(&res, listAllStudyPlacesQuery)
	return res, err
}

func (db *dbImpl) ListAllEvents() ([]models.Event, error) {
	res := []models.Event{}
	err := db.sqliteDB.Select(&res, listAllEventsQuery)
	return res, err
}

func (db *dbImpl) GetEventById(id int) (*models.Event, error) {
    result := models.Event{}
    err := db.sqliteDB.Get(&result, getEventByIdQuery, id)
    
    if err == sql.ErrNoRows {
        return nil, nil
    }

    return &result, err
}