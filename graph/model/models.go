package model

import pq "github.com/lib/pq"

type StaffData struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	DepartmentID int
	Role         string
}

type TaskData struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Detail      string
	UserID      int
	Status      string
	ProjectID   int
	LocationID  int
	StaffIDs    pq.Int64Array `gorm:"type:integer[]"`
}
