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

type Message struct {
	ID        int    `json:"id"`
	Sender    *Staff `json:"sender"`
	Recipient *Staff `json:"recipient"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type MessageInput struct {
	RoomID  int    `json:"roomID"`
	Content string `json:"content"`
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Messages     []*Message `json:"messages"`
	Participants []*Staff   `json:"participants"`
}

type RoomInput struct {
	Name           string `json:"name"`
	ParticipantIDs []int  `json:"participantIDs"`
}

type Staff struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Department *Department `json:"department"`
	Role       string      `json:"role"`
	Password   string      `json:"password"`
	Token      string      `json:"token"`
}

type StaffInput struct {
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentID"`
	Role         string `json:"role"`
	Password     string `json:"password"`
}

type Task struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Detail      string      `json:"detail"`
	User        *Staff      `json:"user"`
	Status      string      `json:"status"`
	Project     *Project    `json:"project"`
	Location    *Location   `json:"location"`
	Department  *Department `json:"department"`
	Staff       []*Staff    `json:"staff"`
}

type TaskInput struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Detail       string `json:"detail"`
	UserID       int    `json:"userID"`
	Status       string `json:"status"`
	ProjectID    int    `json:"projectID"`
	LocationID   int    `json:"locationID"`
	DepartmentID int    `json:"departmentID"`
	StaffIDs     []int  `json:"staffIDs"`
}

type Login struct {
	UserID int    `json:"userID"`
	Token  string `json:"token"`
}
