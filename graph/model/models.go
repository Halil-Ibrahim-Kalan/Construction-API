package model

import "github.com/lib/pq"

type StaffData struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	DepartmentID int
	Role         string
	Password     string
	Token        string
}

type TaskData struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Description  string
	Detail       string
	UserID       int
	Status       string
	ProjectID    int
	LocationID   int
	DepartmentID int
	StaffIDs     pq.Int64Array `gorm:"type:integer[]"`
}
