package utils

import (
	"Construction-API/graph/model"
	"errors"
	"strconv"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Permission int

const (
	Administrator Permission = iota
	Staff
)

func ToProject(id int, r *gorm.DB) *model.Project {
	project := model.Project{}
	err := r.First(&project, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &project
}

func ToLocation(id int, r *gorm.DB) *model.Location {
	location := model.Location{}
	err := r.First(&location, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &location
}

func ToDepartment(id int, r *gorm.DB) *model.Department {
	department := model.Department{}
	err := r.First(&department, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return &department
}

func MapStaffFromInput(staffIDs []int, r *gorm.DB) ([]*model.Staff, error) {
	var staff []*model.Staff
	for _, staffID := range staffIDs {
		if staffID != -1 {
			staffMemberData, err := ToStaff(staffID, r)
			if err != nil {
				return nil, err
			}
			staffMember := &model.Staff{
				ID:         staffMemberData.ID,
				Name:       staffMemberData.Name,
				Department: staffMemberData.Department,
				Role:       staffMemberData.Role,
				Password:   "secret",
				Token:      "secret",
			}
			staff = append(staff, staffMember)
		}
	}
	return staff, nil
}

func ToTask(data model.TaskData, r *gorm.DB) (model.Task, error) {
	userData, _ := ToStaff(data.UserID, r)
	user := model.Staff{
		ID:         userData.ID,
		Name:       userData.Name,
		Department: userData.Department,
		Role:       userData.Role,
		Password:   "secret",
		Token:      "secret",
	}
	project := ToProject(data.ProjectID, r)
	location := ToLocation(data.LocationID, r)
	department := ToDepartment(data.DepartmentID, r)
	staff, err := MapStaffFromInput(ArrayToSlice(data.StaffIDs), r)
	if err != nil {
		return model.Task{}, err
	}

	task := model.Task{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Detail:      data.Detail,
		User:        &user,
		Status:      data.Status,
		Project:     project,
		Location:    location,
		Department:  department,
		Staff:       staff,
	}
	return task, nil
}

func SliceToArray(slice []int) pq.Int64Array {
	var array pq.Int64Array
	for _, value := range slice {
		array = append(array, int64(value))
	}
	return array
}

func ToTaskData(input model.TaskInput, r *gorm.DB) model.TaskData {
	var staffIDs []int64

	for _, value := range input.StaffIDs {
		staffIDs = append(staffIDs, int64(value))
	}

	data := model.TaskData{
		Name:         input.Name,
		Description:  input.Description,
		Detail:       input.Detail,
		UserID:       input.UserID,
		Status:       input.Status,
		ProjectID:    input.ProjectID,
		LocationID:   input.LocationID,
		DepartmentID: input.DepartmentID,
		StaffIDs:     pq.Int64Array(staffIDs),
	}
	return data
}

func ToStaff(id int, r *gorm.DB) (*model.Staff, error) {
	var data *model.StaffData
	err := r.Set("gorm:auto_preload", true).Where("id = ?", id).First(&data).Error

	if err != nil {
		return nil, err
	}

	var department *model.Department
	err = r.Set("gorm:auto_preload", true).Where("id = ?", data.DepartmentID).First(&department).Error

	if err != nil {
		return nil, err
	}

	staffMember := model.Staff{
		ID:         data.ID,
		Name:       data.Name,
		Department: department,
		Role:       data.Role,
		Password:   data.Password,
		Token:      data.Token,
	}

	return &staffMember, nil
}

func ArrayToSlice(array pq.Int64Array) []int {
	var intSlice []int
	for _, value := range array {
		intSlice = append(intSlice, int(value))
	}
	return intSlice
}

func CheckTask(input model.TaskInput, r *gorm.DB) (bool, error) {
	user, err := ToStaff(input.UserID, r)
	if user == nil || err != nil {
		return false, errors.New("User not found! UserID: " + strconv.Itoa(input.UserID))
	}

	project := ToProject(input.ProjectID, r)
	if project == nil {
		return false, errors.New("Project not found! ProjectID: " + strconv.Itoa(input.ProjectID))
	}

	location := ToLocation(input.LocationID, r)
	if location == nil {
		return false, errors.New("Location not found! LocationID: " + strconv.Itoa(input.LocationID))
	}

	if input.StaffIDs == nil {
		return false, errors.New("StaffIDs not found")
	}

	for _, staffID := range input.StaffIDs {
		staffMember, err := ToStaff(staffID, r)
		if staffMember == nil || err != nil {
			return false, errors.New("StaffID not found! staffID: " + strconv.Itoa(staffID))
		}
	}

	return true, nil
}

func DataToStaff(data model.StaffData, r *gorm.DB) model.Staff {
	var department *model.Department
	err := r.Set("gorm:auto_preload", true).Where("id = ?", data.DepartmentID).First(&department).Error

	if err != nil {
		return model.Staff{}
	}

	staffMember := model.Staff{
		ID:         data.ID,
		Name:       data.Name,
		Department: department,
		Role:       data.Role,
		Password:   data.Password,
		Token:      data.Token,
	}

	return staffMember
}

func CheckUserPass(username, password string, checkPass bool, r *gorm.DB) (bool, error) {
	var data *model.StaffData
	var err error
	if checkPass {
		err = r.Where("name = ?", username).Where("password = ?", password).First(&data).Error
	} else {
		err = r.Where("name = ?", username).First(&data).Error
	}

	if err != gorm.ErrRecordNotFound && err != nil {
		return false, err
	}

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}

	return false, nil
}

func GenerateToken(name string) (string, error) {
	userPassword := []byte(name)

	hashedPassword, err := bcrypt.GenerateFromPassword(userPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GetPermission(token string, r *gorm.DB) (Permission, error) {
	var role string
	err := r.Model(&model.StaffData{}).Select("role").Where("token = ?", token).Scan(&role).Error

	if err != gorm.ErrRecordNotFound && err != nil {
		return 0, err
	}

	if err == gorm.ErrRecordNotFound {
		return 0, errors.New("invalid token")
	}

	if role == "Administrator" {
		return Administrator, nil
	}

	return Staff, nil
}

func ToMessage(id int, r *gorm.DB) (*model.Message, error) {
	messageData := model.MessageData{}
	err := r.First(&messageData, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	room, err := ToRoom(messageData.RoomID, r)
	if err != nil {
		return nil, err
	}

	sender, err := ToStaff(messageData.SenderID, r)
	if err != nil {
		return nil, err
	}

	recipient, err := ToStaff(messageData.RecipientID, r)
	if err != nil {
		return nil, err
	}

	sender.Password = "secret"
	sender.Token = "secret"

	recipient.Password = "secret"
	recipient.Token = "secret"

	message := model.Message{
		ID:        messageData.ID,
		Room:      room,
		Sender:    sender,
		Recipient: recipient,
		Content:   messageData.Content,
		Timestamp: messageData.Timestamp.Format("2006-01-02 15:04:05"),
	}
	return &message, nil
}

func ToRoom(id int, r *gorm.DB) (*model.Room, error) {
	var roomData model.RoomData
	err := r.Set("gorm:auto_preload", true).Where("id = ?", id).First(&roomData).Error
	if err != nil {
		return nil, err
	}

	participants, err := MapStaffFromInput(ArrayToSlice(roomData.ParticipantIDs), r)
	if err != nil {
		return nil, err
	}

	room := model.Room{
		ID:           roomData.ID,
		Name:         roomData.Name,
		Participants: participants,
	}

	return &room, nil
}

func ToRoomByUser(id, userID int, r *gorm.DB) (*model.Room, error) {
	room, err := ToRoom(id, r)
	if err != nil {
		return nil, err
	}

	found := false
	for _, participant := range room.Participants {
		if participant.ID == userID {
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("permission denied")
	}

	return room, nil
}

func InputToMessageAndMessageData(input model.MessageInput, token string, r *gorm.DB) (*model.Message, *model.MessageData, error) {
	perm, err := GetPermission(token, r)
	if err != nil {
		return nil, nil, err
	}

	var senderID int
	err = r.Model(&model.StaffData{}).Select("id").Where("token = ?", token).First(&senderID).Error
	if err != nil {
		return nil, nil, err
	}

	if senderID == input.RecipientID {
		return nil, nil, errors.New("cannot send message to yourself")
	}

	sender, err := ToStaff(senderID, r)
	if err != nil {
		return nil, nil, err
	}

	recipient, err := ToStaff(input.RecipientID, r)
	if err != nil {
		return nil, nil, err
	}

	var room *model.Room
	if perm == Administrator {
		room, err = ToRoom(input.RoomID, r)
	} else {
		room, err = ToRoomByUser(input.RoomID, senderID, r)
	}
	if err != nil {
		return nil, nil, err
	}

	timestampData := time.Now()
	timestamp := timestampData.Format("2006-01-02 15:04:05")

	messageData := model.MessageData{
		RoomID:      input.RoomID,
		SenderID:    senderID,
		RecipientID: input.RecipientID,
		Content:     input.Content,
		Timestamp:   timestampData,
	}

	message := model.Message{
		Room:      room,
		Sender:    sender,
		Recipient: recipient,
		Content:   input.Content,
		Timestamp: timestamp,
	}

	return &message, &messageData, nil
}

func TokenToUser(token string, r *gorm.DB) (*model.Staff, error) {
	var userID int
	err := r.Model(&model.StaffData{}).Select("id").Where("token = ?", token).First(&userID).Error
	if err != nil {
		return nil, err
	}

	user, err := ToStaff(userID, r)
	if err != nil {
		return nil, err
	}
	return user, nil
}
