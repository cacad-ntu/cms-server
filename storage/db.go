package storage
import (
    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
    "../models"
    
)

// TODO: Use singleton for db
type DB interface {
    GetStudyPlaceById(id int) (*models.Study_place, error)
    ListAllStudyPlaces() ([]models.Study_place, error)
    GetLocationHistoriesById(id int) ([]models.Location_history, error)
    GetSubLocationsById(id int) ([]models.Sub_location, error)
    ListAllEvents() ([]models.Event, error)
    GetEventById(id int) (*models.Event, error)
    GetLastEventId() (int, error)
    AddEvent(event *models.Event) (error)
}

type dbImpl struct {
    sqliteDB *sqlx.DB
}

func NewDB(fileName string) (DB, error) {
    db, err := sqlx.Open("sqlite3", fileName)
    if err != nil {
        return nil, err
    }
    db.MustExec("PRAGMA foreign_keys = ON;")
    initDBSchema(db)
    return &dbImpl{db}, err
}