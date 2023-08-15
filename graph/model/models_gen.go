package model

type Department struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Location struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Project struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Staff struct {
	ID         int         `gorm:"primaryKey"`
	Name       string      `gorm:"not null"`
	Department *Department `gorm:"foreignKey:department_id"`
	Role       string      `gorm:"not null"`
}

type StaffInput struct {
	Name         string `gorm:"not null"`
	DepartmentID int    `gorm:"not null"`
	Role         string `gorm:"not null"`
}

type Task struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Detail      string    `gorm:"not null"`
	User        *Staff    `gorm:"foreignKey:userID"`
	Status      string    `gorm:"not null"`
	Project     *Project  `gorm:"foreignKey:projectID"`
	Location    *Location `gorm:"foreignKey:locationID"`
	Staff       []*Staff  `gorm:"manyToMany:staff_tasks"`
}

type TaskInput struct {
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Detail      string `gorm:"not null"`
	UserID      int    `gorm:"not null"`
	Status      string `gorm:"not null"`
	ProjectID   int    `gorm:"not null"`
	LocationID  int    `gorm:"not null"`
	StaffIDs    []int  `gorm:"not null"`
}

func (d *Department) TableName() string {
	return "departments"
}

func (l *Location) TableName() string {
	return "locations"
}

func (p *Project) TableName() string {
	return "projects"
}

func (s *Staff) TableName() string {
	return "staffs"
}

func (t *Task) TableName() string {
	return "tasks"
}
