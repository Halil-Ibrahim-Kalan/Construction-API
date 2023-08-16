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
	ID           int        `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	DepartmentID int        `json:"departmentID"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Role         string     `json:"role"`
}

type Task struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Detail      string   `json:"detail"`
	UserID      int      `json:"userID"`
	User        Staff    `json:"user" gorm:"foreignKey:UserID"`
	Status      string   `json:"status"`
	ProjectID   int      `json:"projectID"`
	Project     Project  `json:"project" gorm:"foreignKey:ProjectID"`
	LocationID  int      `json:"locationID"`
	Location    Location `json:"location" gorm:"foreignKey:LocationID"`
	StaffIDs    []int    `json:"staffIDs" gorm:"-"`
	Staff       []Staff  `json:"staff" gorm:"many2many:task_staffs;"`
}

type StaffInput struct {
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentID"`
	Role         string `json:"role"`
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
