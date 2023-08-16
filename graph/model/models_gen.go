// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Staff struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Department *Department `json:"department"`
	Role       string      `json:"role"`
}

type StaffInput struct {
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentID"`
	Role         string `json:"role"`
}

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Detail      string    `json:"detail"`
	User        *Staff    `json:"user"`
	Status      string    `json:"status"`
	Project     *Project  `json:"project"`
	Location    *Location `json:"location"`
	Staff       []*Staff  `json:"staff"`
}

type TaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	UserID      int    `json:"userID"`
	Status      string `json:"status"`
	ProjectID   int    `json:"projectID"`
	LocationID  int    `json:"locationID"`
	StaffIDs    []int  `json:"staffIDs"`
}