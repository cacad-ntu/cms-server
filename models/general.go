package models

type Sub_location struct {
	Name string `json:"name,omitempty" db:"name"`
	Capacity int `json:"capacity,omitempty" db:"capacity"`
	Occupied int `json:"occupied,omitempty" db:"occupied"`
	Density int `json:"density,omitempty" db:"density"`
	SeatingPlan string `json:"seatingPlan,omitempty" db:"seatingPlan"`
}

type Location_history struct {
	Timestamp string `json:"timestamp,omitempty" db:"timestamp"`
	Value int `json:"value,omitempty" db:"value"`
}

type Study_place struct {
	Id int `json:"id" db:"locationID"`
	Name string `json:"name,omitempty" db:"name"`
	Capacity int `json:"capacity,omitempty" db:"capacity"`
	Occupied int `json:"occupied,omitempty" db:"occupied"`
	Density int `json:"density,omitempty" db:"density"`
	Image string `json:"image" db:"image"` 
	Images []string `json:"images"`
	Type string `json:"type,omitempty" db:"type"`
	HasSeatingPlan bool `json:"hasSeatingPlan," db:"hasSeatingPlan"`
	Levels []Sub_location `json:"levels"`
	Timestamp []string `json:"timestamp"`
	HistoricalDensity [][]int `json:"historicalDensity"`
}

type Event struct {
	Id int `json:"id" db:"eventID"`
	Name string `json:"name" db:"name"`
	Location string `json:"location" db:"location"`
	EventType string `json:"eventType" db:"eventType"`
	Crowd int `json:"crowd" db:"crowd"`
	Image string `json:"image" db:"image"`
	Images []string `json:"images"`
}